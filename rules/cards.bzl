"""card: one PNG with a message on it, rendered by Typst."""

load("@typst-bazel//typst:defs.bzl", "typst_document")

def card(name, message, **kwargs):
    """Renders //cards:card.typ with `message` into <name>.png."""
    typst_document(
        name = name,
        main = "//cards:card.typ",
        format = "png",
        inputs = {"message": message},
        deps = ["//cards:card"],
        **kwargs
    )
