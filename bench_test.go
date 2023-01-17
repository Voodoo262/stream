package stream_test

import (
	"github.com/Voodoo262/stream"

	"fmt"
	"os"
	"testing"
)

func BenchmarkSingle(b *testing.B) {
	stream.Run(stream.Repeat("", b.N))
}

func BenchmarkFive(b *testing.B) {
	f := stream.FilterFunc[string](func(arg stream.Arg[string]) error {
		for s := range arg.In {
			arg.Out <- s
		}
		return nil
	})
	stream.Run[string](stream.Repeat("", b.N), f, f, f, f)
}

func BenchmarkWrite(b *testing.B) {
	f, err := os.Create("/dev/null")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	stream.Run(
		stream.Repeat("hello", b.N),
		stream.WriteLines(f),
	)
}

func BenchmarkSample(b *testing.B) {
	stream.Run(
		stream.Repeat("hello", b.N),
		stream.Sample[string](10),
	)
}

func BenchmarkSort(b *testing.B) {
	stream.Run[string](
		stream.Repeat("hello", b.N),
		stream.Sort(),
	)
}

func BenchmarkSort3(b *testing.B) {
	stream.Run[string](
		stream.Repeat("the 3 musketeers", b.N),
		stream.Sort().Num(2).Text(1).Text(3),
	)
}

func BenchmarkCmd(b *testing.B) {
	stream.Run(
		stream.Repeat("hello", b.N),
		stream.Command("cat"),
	)
}

func BenchmarkXargs1(b *testing.B) {
	stream.Run[string](
		stream.Repeat("hello", b.N),
		stream.Xargs("true").LimitArgs(1),
	)
}
