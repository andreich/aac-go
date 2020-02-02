load("@io_bazel_rules_go//go:def.bzl", "go_library")

cc_library(
  name = "voAAC",
  srcs = glob([
    "external/aacenc/src/*.c",
    "external/aacenc/src/*.h",
  ]),
  hdrs = glob(["external/aacenc/include/*.h"]),
  strip_include_prefix = "external/aacenc/include",
  copts = ["-std=gnu99", "-DUSE_DEFAULT_MEM"],
)

go_library(
  name = "aacenc",
  visibility = ["//visibility:public"],
  srcs = ["aacenc.go", "aacenc_cgo.go", "bridge.c"],
  importpath = "github.com/andreich/aac-go/aacenc",
  cgo = True,
  copts = ["-std=gnu99", "-I external/aacenc/include", "-DUSE_DEFAULT_MEM"],
  cdeps = [":voAAC"],
)
