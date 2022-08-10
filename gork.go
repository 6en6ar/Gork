package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
)

var gurl string

func GenerateLink(dork string, domain string) {
	gurl := "https://www.google.com/search?q=" + url.QueryEscape(dork)
	// write to .html
	f, err := os.OpenFile(domain+".html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("[ - ] Error opening file..")
	}
	if _, err := f.Write([]byte("<li><a href=\"" + gurl + "\" target=\"_blank\"" + ">" + gurl + "</a></li>")); err != nil {
		fmt.Println("[ - ] Error writing to file..")
	}
	if err := f.Close(); err != nil {
		fmt.Println("[ - ] Error closing file..")
	}

}
func ReadDorks(domain string) {
	file, err := os.Open("dorks.txt")
	if err != nil {
		fmt.Println("[ - ] Error opening file..")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dork := scanner.Text() + " site:" + domain
		GenerateLink(dork, domain)
	}

}
func CreateHtml(domain string) {
	f, err := os.Create(domain + ".html")
	if err != nil {
		fmt.Println("[ - ] Error creating file")
	}
	defer f.Close()

}

var domain = flag.String("d", "", "domain for dorking")

func usage() {
	fmt.Printf("usage : -d <DOMAIN>\n")
	os.Exit(0)
}
func main() {

	flag.Usage = usage
	flag.Parse()
	if flag.NFlag() == 0 {
		usage()
		os.Exit(1)
	}

	banner := `
 _______  _______  ______    ___   _ 
|       ||       ||    _ |  |   | | |
|    ___||   _   ||   | ||  |   |_| |
|   | __ |  | |  ||   |_||_ |      _|
|   ||  ||  |_|  ||    __  ||     |_ 
|   |_| ||       ||   |  | ||    _  |
|_______||_______||___|  |_||___| |_| . . . .

	Coded by: 6en6ar 3:)

`
	fmt.Printf("Gorking...")
	fmt.Println("hey")
	fmt.Println(banner)
	CreateHtml(*domain)
	fmt.Println("[ + ] File " + *domain + ".html Created")
	ReadDorks(*domain)
	var env = runtime.GOOS
	fmt.Println("[ + ] Opening file in your browser ...")
	switch env {
	case "linux":
		_ = exec.Command("/usr/bin/xdg-open", *domain+".html").Start()
	case "windows":
		_ = exec.Command("C:\\Windows\\System32\\rundll32", "url.dll,FileProtocolHandler", *domain+".html").Start()
	}

}
