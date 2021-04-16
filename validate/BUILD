load("@com_google_protobuf//:protobuf.bzl", "py_proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@io_bazel_rules_go//proto:compiler.bzl", "go_proto_compiler")
load("@rules_cc//cc:defs.bzl", "cc_library", "cc_proto_library")
load("@rules_java//java:defs.bzl", "java_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")
load(
    "@io_bazel_rules_go//proto/wkt:well_known_types.bzl",
    "PROTO_RUNTIME_DEPS",
    "WELL_KNOWN_TYPES_APIV2",
)

package(
    default_visibility =
        ["//visibility:public"],
)

proto_library(
    name = "validate_proto",
    srcs = ["validate.proto"],
    deps = [
        "@com_google_protobuf//:descriptor_proto",
        "@com_google_protobuf//:duration_proto",
        "@com_google_protobuf//:timestamp_proto",
    ],
)

cc_proto_library(
    name = "validate_cc",
    deps = [":validate_proto"],
)

py_proto_library(
    name = "validate_py",
    srcs = ["validate.proto"],
    deps = ["@com_google_protobuf//:protobuf_python"],
)

go_proto_compiler(
    name = "go_protoc",
    deps = PROTO_RUNTIME_DEPS + WELL_KNOWN_TYPES_APIV2,
)

go_proto_library(
    name = "go_default_library",
    compilers = [":go_protoc"],
    importpath = "github.com/envoyproxy/protoc-gen-validate/validate",
    proto = ":validate_proto",
    deps = [
        "@com_github_golang_protobuf//protoc-gen-go/descriptor",
        "@com_github_golang_protobuf//ptypes/duration",
        "@com_github_golang_protobuf//ptypes/timestamp",
    ],
)

cc_library(
    name = "cc_validate",
    hdrs = ["validate.h"],
)

java_proto_library(
    name = "validate_java",
    deps = [":validate_proto"],
)
