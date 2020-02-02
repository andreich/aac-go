load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
  name = "aac-go",
  visibility = ["//visibility:public"],
  srcs = ["encode.go"],
  deps = ["//aacenc:aacenc"],
  importpath = "github.com/andreich/aac-go",
  cgo = True,
)

