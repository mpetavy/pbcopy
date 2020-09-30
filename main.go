package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/mpetavy/common"
	"io/ioutil"
	"os"
)

var enc = flag.String("enc", common.DefaultConsoleEncoding(), "character encoding")
var output = flag.Bool("o", false, "output from clipboard")

func init() {
	common.Init(false, "1.0.0", "", "2018", "Pass text to/from clipboard", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, nil, nil, run, 0)
}

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
	defer common.Done()

	common.Run(nil)
}
