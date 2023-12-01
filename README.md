# @crashmax/xlsx-splitter

<a href="https://www.npmjs.com/package/@crashmax/xlsx-splitter">
  <img alt="npm" src="https://img.shields.io/npm/v/@crashmax/xlsx-splitter?color=orange">
</a>

> Splits an XLSX file into multiple XLSX files.

## Install
```bash
pnpm i -g @crashmax/xlsx-splitter
```

## Usage
```bash
xlsx-splitter --file document.xlsx --rows 1000
```

## Arguments

- `--file` or `-f`: The path to the XLSX file to split.
- `--rows` or `-r`: The number of rows per output file.
- `--output` or `-o` (optional): The path to the output folder.
- `--offset` or `-s` (optional): The offset of the first row in the input file.
