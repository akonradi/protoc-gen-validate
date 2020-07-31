package golang

const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if .MessageRules.GetSkip }}
		// skipping validation for {{ $f.Name }}
	{{ else }}
		if v, ok := interface{}({{ accessor . }}).(interface{ ValidateWithMask(*field_mask.FieldMask) error }); m.maskHas(mask, "{{ $f.Name }}") && ok {
			// update the mask to remove the outer level
			if mask != nil {
				paths := []string{}
				for _, path := range mask.GetPaths() {
					parts := strings.SplitN(path, ".", 2)[1:]
					if len(parts) > 0 {
						paths = append(paths, strings.Join(parts, "."))
					}
				}
				if len(paths) > 0 {
					// if fields were explicitly given within the sub-message, we only
					// validate those specific fields. We remove the prefix and pass the
					// remaining fields down as a new FieldMask for sub-message validation.
					mask = &field_mask.FieldMask{Paths: paths}
				} else {
					// if a sub-message is specified in the last position of the field mask,
					// then we validate the entire sub-message. This matches the expectation
					// of FieldMask on Update operations to overwrite the entire sub-message.
					// https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/field-mask
					mask = nil
				}
			}

			if err := v.ValidateWithMask(mask); err != nil {
				return {{ errCause . "err" "embedded message failed validation" }}
			}
		}
	{{ end }}
`
