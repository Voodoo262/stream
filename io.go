package stream

import (
	"bufio"
	"io"
	"os"
)

// Cat emits each line from each named file in order. If no arguments
// are specified, Cat copies its input to its output.
func Cat(filenames ...string) Filter[string] {
	return FilterFunc[string](func(arg Arg[string]) error {
		if len(filenames) == 0 {
			for s := range arg.In {
				arg.Out <- s
			}
			return nil
		}
		for _, f := range filenames {
			file, err := os.Open(f)
			if err == nil {
				err = splitIntoLines(file, arg)
				file.Close()
			}
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// WriteLines prints each input item s followed by a newline to
// writer; and in addition it emits s.  Therefore WriteLines()
// can be used like the "tee" command, which can often be useful
// for debugging.
func WriteLines(writer io.Writer) Filter[string] {
	return FilterFunc[string](func(arg Arg[string]) error {
		for s := range arg.In {
			if _, err := writer.Write(append([]byte(s), '\n')); err != nil {
				return err
			}
			arg.Out <- s
		}
		return nil
	})
}

// ReadLines emits each line found in reader.
func ReadLines(reader io.Reader) Filter[string] {
	return FilterFunc[string](func(arg Arg[string]) error {
		return splitIntoLines(reader, arg)
	})
}

func splitIntoLines(rd io.Reader, arg Arg[string]) error {
	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		arg.Out <- scanner.Text()
	}
	return scanner.Err()
}
