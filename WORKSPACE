workspace(name = "com_github_istio_api")

git_repository(
    name = "io_bazel_rules_go",
    commit = "eba68677493422112dd25f6a0b4bbdb02387e5a4", # Aug 1, 2017
    remote = "https://github.com/bazelbuild/rules_go.git",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "go_repository")

go_repositories()

go_repository(
    name = "com_github_golang_protobuf",
    commit = "8ee79997227bf9b34611aee7946ae64735e6fd93",
    importpath = "github.com/golang/protobuf",
)

http_archive(
    name = "com_github_google_protobuf",
    sha256 = "2a25c2b71c707c5552ec9afdfb22532a93a339e1ca5d38f163fe4107af08c54c",
    strip_prefix = "protobuf-3.2.0",
    url = "https://github.com/google/protobuf/archive/v3.2.0.tar.gz",
)
