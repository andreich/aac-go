load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "io_bazel_rules_go",
    commit = "54a0c697c263612730c0366eead88a819af7a04e",
    remote = "https://github.com/bazelbuild/rules_go.git",
    shallow_since = "1580238866 -0500",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()
# Uncomment below to get the basic example up and running.
#
# git_repository(
#     name = "bazel_gazelle",
#     commit = "0d378ccadef1b527e3b927aabdeaae38f5d46156",
#     remote = "https://github.com/bazelbuild/bazel-gazelle.git",
# )
# 
# load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
# 
# gazelle_dependencies()
# 
# load("@bazel_gazelle//:deps.bzl", "go_repository")
# 
# go_repository(
# 	name = "com_github_youpy_goriff",
#   importpath = "github.com/youpy/go-riff",
#   commit = "557d78c11efbdcdf178bf52723cfba91f02e0bb2",
# )
# 
# go_repository(
# 	name = "com_github_youpy_gowav",
#   importpath = "github.com/youpy/go-wav",
#   commit = "b63a9887d320becede259f24ba8ba7b2459659da",
# )
# 
