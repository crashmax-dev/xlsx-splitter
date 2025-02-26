package main

import (
	"log"
	"os"
	"xlsx-splitter/internal/utils"
	xlsxParser "xlsx-splitter/internal/xlsx-parser"

	"github.com/urfave/cli/v2"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "xlsx file path",
				Required: true,
				Action:   utils.ValidateXlsxFile,
			},
			&cli.IntFlag{
				Name:     "rows",
				Aliases:  []string{"r"},
				Usage:    "number of rows to split",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "output",
				Aliases:  []string{"o"},
				Usage:    "output folder path",
				Required: false,
				Action:   utils.ValidateOutputPath,
			},
			&cli.IntFlag{
				Name:     "offset",
				Aliases:  []string{"s"},
				Usage:    "offset row number",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			file := c.String("file")
			rows := c.Int("rows")
			offset := c.Int("offset")

			output := c.String("output")
			if output == "" {
				output = cwd
			}

			xlsxParser := xlsxParser.New(file, output)
			xlsxParser.ReadTable(rows, offset)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
