# static
Opinionated static site compiler written in golang

## Description

static is a command tool to aide in building web sites where content is generally unchanged, a blog with commenting and API backend that keeps a lot of state server-side is not a great use case for static.  However, if you'd like a simple way of creating web content using a concise templating language then static is for you.  You don't need to write any javascript code or learn any heavyweight frameworks.  All content is compiled ahead of time and can be served from various web servers like apache or reverse proxies like nginx.  Content and layout should be the focus of your site and static takes care of all meanial tasks so you can focus on the more important things.

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
index.amber
indexbar.amber
layout.amber
menubar.amber
.static.clean
.static.conf
`

### `static config`

Lists configuration stored in .staticrc

`static config`

```javascript
{"exclude":"indexbar.amber,footer.amber, layout.amber"}
```

`static config --exclude="indexbar.amber, footer.amber, layout.amber"`

#### .static.conf

Contains a json structure of all configuration parameters.

Parameter | Description
---|---
exclude | Templates to exclude during compilation

### `static compile`

Converts all .amber files to .html

Parameter | Description
---|---
dir | Location of source .amber files
out | Location to store compiled .html files, directory must exist and user must have permissions to write.

`static compile --dir=test --out=production`

### `static clean`

Removes all compiled output (html files) along with .static.clean and .static.conf.

#### .static.clean

Contains a json structure of all compiled files.

### `static version`

Shows version of static.

### `static server`

Runs a mini-http server with content

## FAQ

### Why not just use Hugo?

Hugo is a great tool also for creating static sites, 

## TODO

- [x] generate .static file that lists all the compiled output files
- [x] configuration file to avoid certain file types like indexbar.amber
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