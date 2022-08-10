package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
)

var gurl string

func GenerateLink(dork string) {
	gurl := "https://www.google.com/search?q=" + url.QueryEscape(dork)
	// write to .html
	fmt.Println(gurl)
	fmt.Println("-------------------------------------------------------------------")

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
		GenerateLink(dork)
	}

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
	ReadDorks(*domain)
}
