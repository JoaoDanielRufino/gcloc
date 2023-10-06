# GCloc

A cli written in Go inspired by [cloc](https://github.com/AlDanial/cloc) to count the lines, blank lines and comment lines of source code.

## Installation

You can install from the stable release by clicking [here](https://github.com/JoaoDanielRufino/gcloc/releases/latest)

## Usage

### Basic example

```
$ gcloc .
  Language | Files | Lines | Blank lines | Comments | Code lines
-----------+-------+-------+-------------+----------+-------------
  Golang   |    26 |  1588 |         253 |       12 |       1323
  HTML     |     2 |   576 |          82 |       16 |        478
  YAML     |     4 |   108 |           8 |        0 |        100
  Makefile |     1 |    22 |           3 |        0 |         19
-----------+-------+-------+-------------+----------+-------------
   Total   |  33   | 2294  |     346     |    28    |    1920
-----------+-------+-------+-------------+----------+-------------
```

### Excluding directories

```
$ gcloc jazzy-bot -e=node_modules --order-by-comment
   Language  | Files | Lines | Blank lines | Comments | Code lines
-------------+-------+-------+-------------+----------+-------------
  YAML       |     1 |    55 |           8 |        2 |         45
  JavaScript |    26 |  1181 |           0 |        0 |       1181
  TypeScript |    24 |  1180 |         189 |        0 |        991
-------------+-------+-------+-------------+----------+-------------
    Total    |  51   | 2416  |     197     |    2     |    2217
-------------+-------+-------+-------------+----------+-------------
```

### From git remote repository

```
$ gcloc github.com/JoaoDanielRufino/gcloc
  Language | Files | Lines | Blank lines | Comments | Code lines
-----------+-------+-------+-------------+----------+-------------
  Golang   |    32 |  3086 |         310 |       12 |       2764
  HTML     |     1 |   167 |          14 |       16 |        137
  YAML     |     4 |   110 |           8 |        0 |        102
  Makefile |     1 |    22 |           3 |        0 |         19
-----------+-------+-------+-------------+----------+-------------
   Total   |  38   | 3385  |     335     |    28    |    3022
-----------+-------+-------+-------------+----------+-------------
```

### Compressed files

Supported archive formats from [go-getter](https://github.com/hashicorp/go-getter#unarchiving)

```
$ gcloc ~/Documents/gcloc.tar.gz
  Language | Files | Lines | Blank lines | Comments | Code lines
-----------+-------+-------+-------------+----------+-------------
  Golang   |    32 |  3086 |         310 |       12 |       2764
  HTML     |     1 |   167 |          14 |       16 |        137
  YAML     |     4 |   110 |           8 |        0 |        102
  Makefile |     1 |    22 |           3 |        0 |         19
-----------+-------+-------+-------------+----------+-------------
   Total   |  38   | 3385  |     335     |    28    |    3022
-----------+-------+-------+-------------+----------+-------------
```

### Exporting results

```
$ gcloc github.com/JoaoDanielRufino/gcloc --report-formats=prompt,json
   Language  | Files | Lines | Blank lines | Comments | Code lines
-------------+-------+-------+-------------+----------+-------------
  Golang     |    34 |  3591 |         364 |       12 |       3215
  Markdown   |     1 |   169 |          25 |        0 |        144
  HTML       |     1 |   167 |          14 |       16 |        137
  YAML       |     4 |   110 |           8 |        0 |        102
  Json       |     1 |    20 |           0 |        0 |         20
  Makefile   |     1 |    22 |           3 |        0 |         19
  Plain Text |     1 |     1 |           0 |        0 |          1
-------------+-------+-------+-------------+----------+-------------
    Total    |  43   | 4080  |     414     |    28    |    3638
-------------+-------+-------+-------------+----------+-------------
json report exported to /home/joaodaniel/Documents/gcloc/cmd/gcloc/result.json
```

### GitHub Actions integration

```yaml
name: GCloc

on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
      steps:
      - uses: actions/checkout@v3
      - name: Run GCloc
        run: docker run -t -v $PWD:/gcloc joaodanielrufino/gcloc /gcloc --report-formats=prompt,json -o /gcloc
      - name: Archive gcloc result
        uses: actions/upload-artifact@v3
        with:
          name: gcloc-result
          path: result.json
```

### Via Docker

From [dockerhub](https://hub.docker.com/r/joaodanielrufino/gcloc)

```
docker run --rm -t -v $PWD:/gcloc joaodanielrufino/gcloc /gcloc
```

### Flags

```
$ gcloc -h
GCloc is a simple tool to count lines of code of many programming languages

Usage:
  gcloc [flags]
  gcloc [command]

Available Commands:
  help        Help about any command
  languages   Show gcloc supported languages

Flags:
      --by-file                      Show results by file
  -e, --exclude strings              Exclude directories or files from being scanned
      --exclude-extensions strings   Exclude extensions from being scanned
  -h, --help                         help for gcloc
      --include-extensions strings   Include the extensions to be scanned
      --order string                 Sorting order <ASC,DESC> (default "DESC")
      --order-by-blank               Show results ordered by blank lines
      --order-by-code                Show results ordered by lines of code
      --order-by-comment             Show results ordered by comments
      --order-by-file                Show results ordered by file count
      --order-by-lang                Show results ordered by language
      --order-by-line                Show results ordered by lines count
      --output-name string           Name of report output (default "result")
  -o, --output-path string           Path where the report will be exported
      --report-formats strings       Report formats in which the results will be exported (prompt, json) (default [prompt])
  -v, --version                      version for gcloc

Use "gcloc [command] --help" for more information about a command.
```

## Supported languages

To show all supported languages use the subcommand `languages`

```
$ gcloc languages
      Language     |   Extensions   | Single Comments | Multi Line Comments
-------------------+----------------+-----------------+----------------------
  ActionScript     |  .as           |  //             |  /* */
  Assembly         |  .asm          |  // ; # @ | !   |  /* */
  Bash             |  .bash .sh     |  #              |
  C Header         |  .h            |  //             |  /* */
  C#               |  .cs           |  //             |  /* */
  C++              |  .cpp .cc      |  //             |  /* */
  C++ Header       |  .hh .hpp      |  //             |  /* */
  COBOL            |  .cbl          |  * /            |
  CSS              |  .css          |  //             |  /* */
  Clojure          |  .clj          |  ; ;; #_        |
  CoffeeScript     |  .coffee       |  #              |  ### ###
  Dart             |  .dart         |  // ///         |  /* */
  Elixir           |  .ex .exs      |  #              |
  Erlang           |  .erl          |  % %%           |
  Haskell          |  .hs           |  --             |
  Java             |  .java         |  //             |  /* */
  JavaScript       |  .js .jsx      |  //             |  /* */
  Json             |  .json         |                 |
  Jupyter Notebook |  .ipynb        |  #              |
  Kotlin           |  .kt .kts      |  //             |  /* */
  Lisp             |  .lsp .lisp    |  ;              |
  Lua              |  .lua          |  --             |  --[[ ]]
  Makefile         |  Makefile      |  #              |
  Markdown         |  .md .markdown |                 |
  Maven            |  .maven        |  <!--           |  <!-- -->
  PHP              |  .php          |  // #           |  /* */
  Perl             |  .pl           |  #              |  = =cut
  Plain Text       |  .txt .text    |                 |
  PowerShell       |  .ps1          |  #              |
  Processing       |  .pde          |  //             |  /* */
  Protocol Buffers |  .proto        |  //             |
  PRQL             |  .prql         |  #              |
  Python           |  .py           |  #              |  """ """
  R                |  .r .R         |  #              |
  Rego             |  .rego         |  #              |
  Ruby             |  .rb           |  #              |  =begin =end
  Rust             |  .rs           |  //             |  /* */
  SQL              |  .sql          |  --             |  /* */
  Scala            |  .scala        |  //             |  /* */
  Scss             |  .scss         |  //             |  /* */
  Swift            |  .swift        |  //             |  /* */
  TypeScript       |  .ts .tsx      |  //             |  /* */
  Vim              |  .vim          |  "              |
  Vue              |  .vue          |  <!--           |  <!-- -->
  XML              |  .xml .XML     |  <!--           |  <!-- -->
  YAML             |  .yaml .yml    |  #              |
  Zsh              |  .zsh          |  #              |
```
