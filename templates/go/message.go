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
				mask = &field_mask.FieldMask{Paths: paths}
			}

			if err := v.ValidateWithMask(mask); err != nil {
				return {{ errCause . "err" "embedded message failed validation" }}
			}
		}
	{{ end }}
`
