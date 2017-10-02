package main;

import(
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (

	CMD_CLEAN     = "clean"
	CMD_COMPILE   = "compile"
	CMD_INIT			= "init"
	CMD_SERVER    = "server"
	CMD_VERSION		= "version"

)

const (
	AMBER   					= ".amber"
	MARKDOWN					= ".markdown"
	MD      					= ".md"
)

const (
	CLEAN_FILE        = ".static.json"
	CONFIG_FILE       = ".staticrc"
	DEFAULT_ADDRESS   = "localhost"
	DEFAULT_PORT      = "8888"
	PWD								= "."
	VERSION 					= "0.1.0"
)

var (

	cmdClean    = kingpin.Command("clean", "Clean up site")
	cmdInit			= kingpin.Command("init", "Initialize site")
	cmdVersion	=	kingpin.Command("version", "static version")

	cmdServer		=	kingpin.Command("server", "Starts up local server for testing")
	cmdPort			=	cmdServer.Flag("port", "Port used by static server").String()

	cmdCompile	=	kingpin.Command("compile", "Compile templates")
	cmdDir			= cmdCompile.Flag(
		"in", "Source file location").Short('d').String()
	cmdOut	    = cmdCompile.Flag(
		"out", "Output location").Short('o').String()
	cmdForce    = cmdCompile.Flag(
		"force", "Overwrite if files exist?").Short('f').Bool()

)

func address() string {

	if len(*cmdPort) == 0 {
		return fmt.Sprintf("%s:%s", DEFAULT_ADDRESS, DEFAULT_PORT)
	} else {
		return fmt.Sprintf("%s:%s", DEFAULT_ADDRESS, *cmdPort)
	}
	
} // address

func initRouter() *mux.Router {

	router := mux.NewRouter()

	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets",
		http.FileServer(http.Dir("./assets"))))
	
	router.HandleFunc("/", genericHandler)
	router.HandleFunc("/{resource}", genericHandler)

	return router

} // initRouter

func main() {
	
	switch kingpin.Parse() {
	case CMD_CLEAN:
		color.Green("Deleting files...")
		cleanFiles()
		
	case CMD_INIT:
		color.Green("Creating static site contents...")
	case CMD_SERVER:
		color.Green("Starting server on %s...", address())

		router := initRouter()

		color.Red("Fatal: %s", http.ListenAndServe(address(), router))

	case CMD_VERSION:
		color.Green("static v%s", VERSION)
	case CMD_COMPILE:
		color.Yellow("Compiling...")
		compile()
	}
	
} // main
