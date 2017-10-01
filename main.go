package main;

import(
	//"fmt"

	"github.com/fatih/color"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (

	CMD_INIT			= "init"
	CMD_COMPILE   = "compile"
	CMD_VERSION		= "version"

)

const (
	AMBER   	= ".amber"
	MARKDOWN	= ".markdown"
	MD      	= ".md"
	PWD				= "."
	VERSION 	= "0.1.0"
)

var (

	cmdInit			= kingpin.Command("init", "Initialize site")
	cmdVersion	=	kingpin.Command("version", "static version")

	cmdCompile	=	kingpin.Command("compile", "Compile templates")
	cmdDir			= cmdCompile.Flag(
		"in", "Source file location").Short('d').String()
	cmdOut	    = cmdCompile.Flag(
		"out", "Output location").Short('o').String()
	cmdForce    = cmdCompile.Flag(
		"force", "Overwrite if files exist?").Short('f').Bool()

)

func main() {
	
	switch kingpin.Parse() {
	case CMD_INIT:
		color.Green("Creating static site contents...")
	case CMD_VERSION:
		color.Green("static v%s", VERSION)
	case CMD_COMPILE:
		color.Yellow("Compiling...")
		compile()
	}
	
} // main
