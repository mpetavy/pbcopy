package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/atotto/clipboard"
	"github.com/mpetavy/common"
)

var enc = flag.String("enc", common.DefaultConsoleEncoding(), "character encoding")
var output = flag.Bool("o", false, "output from clipboard")

func run() error {
	if *output {
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

		return nil
	} else {
		reader := bufio.NewReader(os.Stdin)
		b, err := ioutil.ReadAll(reader)
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

		return nil
	}
}

func main() {
	defer common.Cleanup()

	common.NoBanner = true

	common.New(&common.App{"pbcopy", "1.0.0", "2018", "Pass text to/from clipboard", "mpetavy", common.APACHE, "https://github.com/mpetavy/pbcopy", false, nil,nil, nil, run, time.Duration(0)}, nil)
	common.Run()
}
