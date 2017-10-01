# static
Opinionated static site compiler written in golang

## Description

static is for sites where content is generally unchanged, a blog with commenting and
API backend that keeps a lot of state server-side is not a great use case for static.
However, if you'd like a simple way of creating web content using a templating language, then static is for you.  You don't need to write any javascript code or learn any heavyweight frameworks.  All content is compiled ahead of time and can be served with various web servers.  Content and formatting are key to static sites so keep the focus
and let static do the rest for you.

## Features

1.  Compile .amber, .markdown, .md files into .html
1.  Allow for testing of static content before staging
1.  Facilitate staging of sites

## Requirements

1.  golang 1.9+
1.  `go get github.com/fatih/color`
1.  `go get gopkg.in/alecthomas/kingpin.v2`
1.  `go get github.com/russross/blackfriday`
1.  `go get github.com/eknkc/amber`

## Usage

### `static init`

_creates the following hierarchy and files in the current working directory:_

`
./assets/css
./assets/js
./index.haml
./layout.haml
`

### `static compile`
### `static clean`
### `static version`
### `static server`
