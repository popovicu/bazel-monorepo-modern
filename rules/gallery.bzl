"""gallery: an HTML page that shows a set of PNG images."""

def _gallery_impl(ctx):
    out = ctx.actions.declare_file(ctx.attr.out)
    images = ctx.files.images

    args = ctx.actions.args()
    args.add("--out", out)
    args.add_all(images)

    ctx.actions.run(
        executable = ctx.executable._tool,
        arguments = [args],
        inputs = images,
        outputs = [out],
        mnemonic = "Gallery",
        progress_message = "Building the gallery page %{label}",
    )

    return [DefaultInfo(files = depset([out]))]

gallery = rule(
    implementation = _gallery_impl,
    attrs = {
        "images": attr.label_list(
            doc = "Targets whose files are the PNG images to show.",
            allow_files = [".png"],
            mandatory = True,
        ),
        "out": attr.string(
            doc = "Name of the HTML file to write.",
            default = "index.html",
        ),
        "_tool": attr.label(
            default = "//tools/gallery",
            executable = True,
            cfg = "exec",
        ),
    },
)
