package main

import (
	"bufio"
	"bytes"
	"embed"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/atotto/clipboard"
	"github.com/mpetavy/common"
)

var enc = flag.String("enc", common.DefaultConsoleEncoding(), "character encoding")
var filename = flag.String("f", "", "read from file")
var output = flag.Bool("o", false, "output from clipboard")

//go:embed go.mod
var resources embed.FS

func init() {
	common.Init("", "", "", "", "Pass text to/from clipboard", "", "", "", &resources, nil, nil, run, 0)
}

func run() error {
	switch {
	case *filename != "":
		ba, err := os.ReadFile(*filename)
		if common.Error(err) {
			return err
		}

		err = clipboard.WriteAll(string(ba))
		if common.Error(err) {
			return err
		}
	case *output:
		t, err := clipboard.ReadAll()
		if err != nil {
			return err
		}

		if len(t) > 0 {
			_, err = fmt.Fprintf(os.Stdout, "%s", t)
			if err != nil {
				return err
			}
		}
	default:
		reader := bufio.NewReader(os.Stdin)
		b, err := io.ReadAll(reader)
		if err != nil {
			return err
		}

		if *enc != "" {
			b, err = common.ToUTF8(bytes.NewReader(b), *enc)
			if err != nil {
				return err
			}
		}

		err = clipboard.WriteAll(string(b))
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	common.Run(nil)
}
