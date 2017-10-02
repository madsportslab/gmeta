# static
Opinionated static site compiler written in golang

## Description

static is for sites where content is generally unchanged, a blog with commenting and
API backend that keeps a lot of state server-side is not a great use case for static.
However, if you'd like a simple way of creating web content using a templating language, then static is for you.  You don't need to write any javascript code or learn any heavyweight frameworks.  All content is compiled ahead of time and can be served from various web servers like apache or reverse proxies like nginx.  Content and formatting are key to static sites so let static do all the meanial tasks while you focus on the
content and layout.

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
assets/css/site.css
assets/js/config.js
assets/js/site.js
index.haml
indexbar.haml
layout.haml
menubar.haml
`

### `static compile`
### `static clean`
### `static version`
### `static server`

## TODO

- [ ] generate .static file that lists all the compiled output files
- [ ] configuration file to avoid certain file types like indexbar.amber
- [ ] source and destination
- [ ] clone skeleton from github
- [ ] upload to github pages
- [ ] upload to s3
- [ ] documentation for deploying to web servers
- [ ] automatically generate web server configuration, e.g. nginx.conf
- [ ] docker file
- [ ] automatically download assets like material-ui, bootstrap, etc
- [ ] markdown support
- [ ] push to github repository