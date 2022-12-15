package main

import (
	_ "embed"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
)

//go:embed template.html
var redirectTemplate string

func usage() {
	fmt.Println("vanity")
	fmt.Printf("\nUsage:\n")
	fmt.Printf("\tFlags:\n")
	fmt.Printf("\t-h\t - display usage information\n")
	fmt.Printf("\t-p\t - optional flag to indicate if the repo is private\n")

	fmt.Printf("\n\tArguments:\n")
	fmt.Printf("\tvanity [-p] [pkg] [url]\n")
	fmt.Printf("\tpkg \t - your desired package import path\n")
	fmt.Printf("\turl \t - url to where your code is actually hosted\n")
}

func generateHTML(pkg, url, protocol string, out io.WriteCloser) error {
	t, err := template.New("redirectTemplate").Parse(redirectTemplate)
	if err != nil {
		return err
	}
	err = t.Execute(out, struct {
		Package  string
		URL      string
		Protocol string
	}{
		pkg,
		url,
		protocol,
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	log.SetOutput(os.Stderr)
	flag.Usage = usage
	p := flag.Bool("p", false, "")
	flag.Parse()

	if len(os.Args) < 3 {
		usage()
		os.Exit(0)
	}

	protocol := "https"

	if *p {
		protocol = "ssh"
	}

	err := generateHTML(os.Args[1], os.Args[2], protocol, os.Stdout)
	if err != nil {
		log.Fatalf("Error: failed to generate HTML\n%v", err)
	}
}
