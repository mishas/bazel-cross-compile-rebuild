workspace(name = "com_github_basel_croll_compile_rebuild")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "2b1641428dff9018f9e85c0384f03ec6c10660d935b750e3fa1492a281a53b0f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.29.0/rules_go-v0.29.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.29.0/rules_go-v0.29.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "ebf1d8ceaa736e85a4e2becf81de297e0be6277085ac9a5f41a4dc5cb6175fbe",
    strip_prefix = "bazel-gazelle-36dfc0fe91acf0271d1a5eea1dad0a409e46ca60",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/archive/36dfc0fe91acf0271d1a5eea1dad0a409e46ca60.zip"],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:repositories.bzl", "go_repositories")

# gazelle:repository_macro repositories.bzl%go_repositories
go_repositories()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
    nogo = "@//:my_nogo",
    version = "1.17.1",
)

gazelle_dependencies()
