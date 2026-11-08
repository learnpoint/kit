package version

import (
	"errors"
	"flag"
	"fmt"
	"io"

	"github.com/learnpoint/kit/internal/cli"
)

func Command(v string) cli.Command {
	return cli.Command{
		Name:    "version",
		Summary: "Print the version",
		Help:    printHelp,
		Run: func(args []string, stdout, stderr io.Writer) int {
			return run(v, args, stdout, stderr)
		},
	}
}

func printHelp(w io.Writer) {
	fmt.Fprint(w, `Usage: kit version

Print the version of kit.

Also available as: kit -v, kit --version
`)
}

func run(v string, args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("version", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.Usage = func() {} // we print help ourselves, so it goes to the right stream
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			printHelp(stdout)
			return 0
		}
		// flag has already printed the error, e.g. "flag provided but not defined: -x".
		printHelp(stderr)
		return 2
	}
	if fs.NArg() > 0 {
		fmt.Fprintln(stderr, "kit version: unexpected arguments")
		printHelp(stderr)
		return 2
	}

	fmt.Fprintln(stdout, v)
	return 0
}
