# static
Opinionated static site compiler written in golang

## Description

Static is for sites where content is generally unchanged, a blog with commenting and
API backend that keeps a lot of state server-side is not a great use case for static.
However, if you'd like a simple way of creating web content using a concise templating language that allows greater levels of re-use and control over the look and feel, then static is for you.

There are 2 motivations for static:

1.  Compile .amber or .markdown files into .html
1.  Provide browser/client side routing

## Requirements

1.  golang 1.9+
1.  `go get github.com/fatih/color`
1.  `go get gopkg.in/alecthomas/kingpin.v2`
1.  `go get github.com/russross/blackfriday`
1.  `go get github.com/eknkc/amber`

## Usage

### Initialize static site

`static init`

_creates the following hierarchy and files in the current directory:_

`
./css
./js
./index.haml
./layout.haml
`

### Compile templates

`static`
