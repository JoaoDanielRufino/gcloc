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
