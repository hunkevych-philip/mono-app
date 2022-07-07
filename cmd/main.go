package main

import (
	"flag"
	"fmt"
	"github.com/hunkevych-philip/mono-app/pkg/handler"
	"github.com/hunkevych-philip/mono-app/pkg/service"
	"github.com/hunkevych-philip/mono-app/pkg/service/excel"
	"github.com/hunkevych-philip/mono-app/pkg/service/mono"
	"log"
	"os"
)

//flag descriptions
var (
	helpFlagDesc   = "Prints descriptions of available flags."
	xTokenFlagDesc = "Your personal token obtained on https://api.monobank.ua/."
)

//flag declarations
var (
	helpFlag   = flag.Bool("help", false, helpFlagDesc)
	xTokenFlag = flag.String("x-token", "", helpFlagDesc)
)

func init() {
	//shorthand notations
	flag.BoolVar(helpFlag, "h", false, helpFlagDesc)
	flag.StringVar(xTokenFlag, "x", "", xTokenFlagDesc)
	// TODO: Add start date flag + end date flag
	// TODO: Add account choice flag
	// TODO: Add a flag for excel filename and location

	//custom flag usage
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage: %s [OPTIONS] -x \"X-Token\"\n\n", os.Args[0])
		fmt.Fprintf(os.Stdout, "Options:\n")
		fmt.Fprintf(os.Stdout, "-h, --help\t\t%s\n", helpFlagDesc)
		fmt.Fprintf(os.Stdout, "-x, --x-token\t\t%s\n", xTokenFlagDesc)
		os.Exit(0)
	}
}

func main() {
	flag.Parse()

	if *helpFlag {
		flag.Usage() // will print usage and exit
	}

	if *xTokenFlag == "" {
		printlnfStdOut("X-Token is required to proceed.")
		flag.Usage()
	}

	var (
		s = service.NewService(mono.NewMonoService(), excel.NewExcelService())
		h = handler.NewHandler(s)
	)

	if err := h.Go(*xTokenFlag, "", ""); err != nil {
		log.Fatalln(err)
	}
}

func printlnfStdOut(msg string) {
	if _, err := fmt.Fprintln(os.Stdout, msg); err != nil {
		log.Fatalf("Failed to send message to stdout: %s\n", msg)
	}
}
