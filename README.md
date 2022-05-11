# GCloc

A cli written in Go inspired by [cloc](https://github.com/AlDanial/cloc) to count the lines, blank lines and comment lines of source code.

## Installation

You can install from the stable release by clicking [here](https://github.com/JoaoDanielRufino/gcloc/releases/latest)

## Usage

### Basic example

```
$ gcloc .
  Language | Lines | Blank lines | Comments | Code lines
-----------+-------+-------------+----------+-------------
  Golang   |  1503 |         248 |       12 |       1243
  HTML     |   576 |          82 |       16 |        478
  YAML     |   108 |           8 |        0 |        100
  Makefile |    22 |           3 |        0 |         19
-----------+-------+-------------+----------+-------------
   Total   | 2209  |     341     |    28    |    1840
-----------+-------+-------------+----------+-------------
```

### Excluding directories

```
$ gcloc jazzy-bot -e=node_modules --order-by-comment
   Language  | Lines | Blank lines | Comments | Code lines
-------------+-------+-------------+----------+-------------
  YAML       |    55 |           8 |        2 |         45
  JavaScript |  1181 |           0 |        0 |       1181
  TypeScript |  1180 |         189 |        0 |        991
-------------+-------+-------------+----------+-------------
    Total    | 2416  |     197     |    2     |    2217
-------------+-------+-------------+----------+-------------
```

### Via Docker

From [dockerhub](https://hub.docker.com/repository/docker/joaodanielrufino/gcloc)

```
docker run --rm -v $PWD:/tmp joaodanielrufino/gcloc /tmp
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
      --by-file                      Show result by file
  -e, --exclude strings              Exclude directories or files from being scanned
      --exclude-extensions strings   Exclude extensions from being scanned
  -h, --help                         help for gcloc
  -o, --order string                 Sorting order <ASC,DESC> (default "DESC")
      --order-by-blank               Show results ordered by blank lines
      --order-by-code                Show results ordered by lines of code
      --order-by-comment             Show results ordered by comments
      --order-by-lang                Show results ordered by language
      --order-by-line                Show results ordered by lines count
  -v, --version                      version for gcloc

Use "gcloc [command] --help" for more information about a command.
```

## Supported languages

To show all supported languages use the subcommand `languages`

```
$ gcloc languages
   Language  | Single Comments | Multi Line Comments
-------------+-----------------+----------------------
  Assembly   |  // ; # @ | !   |  /* */
  C          |  //             |  /* */
  C Header   |  //             |  /* */
  C++        |  //             |  /* */
  C++ Header |  //             |  /* */
  Golang     |  //             |  /* */
  HTML       |                 |  <!-- -->
  Haskell    |  --             |
  Java       |  //             |  /* */
  JavaScript |  //             |  /* */
  Makefile   |  #              |
  PHP        |  // #           |  /* */
  Prolog     |  %              |
  Python     |  #              |  """ """
  TypeScript |  //             |  /* */
  YAML       |  #              |
```
