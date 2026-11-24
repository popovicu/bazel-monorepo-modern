// One card: a message on a dark background. The message comes from Bazel,
// through sys.inputs.
#let message = sys.inputs.at("message", default: "Hello, Typst!")
#let accent = rgb("#D9480F")

#set page(width: 600pt, height: 315pt, margin: 32pt, fill: rgb("#14171A"))
#set text(font: "JetBrains Mono", weight: "extrabold", fill: white)

#text(size: 14pt, fill: accent)[\/\/ built by bazel]
#v(1fr)
#text(size: 40pt, message)
#v(12pt)
#box(width: 48pt, height: 6pt, fill: accent)
