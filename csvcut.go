package main

import (
	"gopkg.in/urfave/cli.v1"
	"os"
	"io"
	"github.com/pkg/errors"
	"encoding/csv"
	"unicode/utf8"
	"strings"
	"strconv"
	"fmt"
)

const Version = "0.1"

func main() {
	app := cli.NewApp()
	app.Name = "csvcut"
	app.Usage = "cut(1) for comma(*) seperated values"
	app.Action = run
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "d, delimiter",
			Value: ",",
			Usage: "delimiter used in the input the CSV, defaults to ','",
		},
		/*
		cli.StringFlag{
			Name: "D, out-delimiter",
			Value: "\t",
			Usage: "delimiter used in the output, defaults to tab",
		},
		*/
		cli.StringFlag{
			Name: "f, fields",
			Usage: "Fields to put put index (starting with 1, not 0), comma seperated",
		},
	}

	app.Run(os.Args)
}

type options struct {
	delimiter rune
	fields []int
}

func (o options) String() string {
	return fmt.Sprintf("{options delimiter:%v, fields:%v}", o.delimiter, o.fields)
}

func run(ctx *cli.Context) error {
	rawFields := ctx.String("f")
	fields := []int{}
	for _, f := range strings.Split(rawFields, ",") {
		fieldIndex, err := strconv.Atoi(f)
		if err != nil {
			return errors.Wrapf(err, "unable to convert %s to an integer", f)
		}
		fields = append(fields, fieldIndex - 1)
	}
	comma, _ := utf8.DecodeRune([]byte(ctx.String("d")))
	opts := options{
		delimiter: comma,
		fields: fields,
	}

	if !ctx.Args().Present() {
		// use stdin!
		process(opts, os.Stdin, os.Stdout)
		return nil
	} else {
		// loop through args and treat them as files.
		for i := 0; i < ctx.NArg(); i++ {
			path := ctx.Args().Get(i)
			in, err := os.Open(path)
			if err != nil {
				return errors.Wrapf(err, "error opening %s", path)
			}
			process(opts, in, os.Stdout)
		}
	}
	return nil
}

func process(opts options, in io.Reader, out io.Writer) error {
	reader := csv.NewReader(in)
	reader.Comma = opts.delimiter

	writer := csv.NewWriter(out)
	writer.Comma = rune('\t')

	idx := 0
	for {
		idx++
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				writer.Flush()
				return nil
			}
			return errors.Wrapf(err, "error reading line %d", idx)
		}
		if len(opts.fields) == 0 {
			// all fields
			writer.Write(row)
		} else  {
			// specific fields
			outRow := []string {}
			for i := 0; i < len(opts.fields); i++ {
				outRow = append(outRow, string(row[opts.fields[i]]))
			}
			writer.Write(outRow)
		}
	}
}
