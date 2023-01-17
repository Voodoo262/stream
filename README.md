# Stream package

This is a fork of [github.com/ghemawat/stream](https://github.com/ghemawat/stream) that allows data types other than `string`

Package stream provides filters that can be chained together in a manner
similar to Unix pipelines.  A simple example that prints all go files
under the current directory:

	stream.Run(
		stream.Find("."),
		stream.Grep(`\.go$`),
		stream.WriteLines(os.Stdout),
	)

## Installation

~~~~
go get github.com/Voodoo262/stream
~~~~

See godoc for further documentation and examples.

*   [godoc.org/github.com/ghemawat/stream](http://godoc.org/github.com/ghemawat/stream)
