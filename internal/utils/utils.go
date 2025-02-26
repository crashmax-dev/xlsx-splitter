package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func ValidateOutputPath(ctx *cli.Context, path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return nil
	}

	return fmt.Errorf("path is not a directory: %s", path)
}

func ValidateXlsxFile(ctx *cli.Context, path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	fileExt := filepath.Ext(fileInfo.Name())
	if fileExt != ".xlsx" {
		return fmt.Errorf("file is not an xlsx file: %s", path)
	}

	return nil
}
