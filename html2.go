package main

import (
	"fmt"
	"strings"
	"os"
	"golang.org/x/net/html"
	"flag"
)
// Helper function to pull the href attribute from a Token
func getHref(t html.Token) (ok bool, href string) {
	// Iterate over token attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}
	// "bare" return will return the variables (ok, href) as 
    // defined in the function definition
	return
}
func  main() {
	htmlFile := flag.String("i", "index.html", "read html file from, default is index.html")
//	output := flag.String("o", "stdout", "output the markdown to, default is stdout")
	flag.Parse()
	var vals string
	// read html file
	bs, err := os.ReadFile(*htmlFile)
	if err != nil{
		fmt.Println("failed to open html file, use -i to specify the html file")
		os.Exit(1)
	}
	// tokenize html
	tkn := html.NewTokenizer(strings.NewReader(string(bs)))
var isP bool
var isH1 bool
var isH2 bool
var isH3 bool
var isStrong bool
var isCode bool
var isA bool
var isLi bool
var isEm bool
var ok bool
var url string
i := 0
for i != 1{
	tt := tkn.Next()
	switch{
	case tt == html.ErrorToken:
	i = 1
	case tt == html.StartTagToken:
	t := tkn.Token()
	isP = t.Data == "p"
	isH1 = t.Data == "h1"
	isH2 = t.Data == "h2"
	isH3 = t.Data == "h3"
	isStrong = t.Data == "strong"
	isEm = t.Data == "em"
	isCode = t.Data == "code"
	isA = t.Data == "a"
	isLi = t.Data == "li"
	if !isA{
	ok = false
	}
	if isA{
	ok, url = getHref(t)
	}
	case tt == html.TextToken:
	    t := tkn.Token()
	    if isH1{
		    vals = fmt.Sprintf("%s\n# %s\n", vals, t.Data)
	    }
	    isH1 = false
	    if isH2{
		    vals = fmt.Sprintf("%s\n## %s\n", vals, t.Data)
	    }
	    isH2 = false
	 if isH3{
		    vals = fmt.Sprintf("%s\n### %s\n", vals, t.Data)
	    }
	    isH3 = false

	    if isP {
		    vals = fmt.Sprintf("%s %s", vals, t.Data)
            }

	    if isStrong {
		    vals = fmt.Sprintf("%s **%s**", vals, t.Data)
		    isP = true
            }
	    isStrong = false
	   if isEm {
		    vals = fmt.Sprintf("%s *%s*", vals, t.Data)
		    isP = true
            }
	    isEm = false

	    if isCode {
		    vals = fmt.Sprintf("%s ``%s``", vals, t.Data)
		    isP = true
            }
	    isCode = false
	    if isA {
		    trimval := strings.TrimSpace(t.Data)
		    if !ok{
		    vals = fmt.Sprintf("%s %s", vals, trimval)
		    } else{
		    vals = fmt.Sprintf("%s [%s](%s)", vals,trimval,url)
	    }
	   isP = true
	            }
		    isA = false
		if isLi{
		    vals = fmt.Sprintf("%s + %s\n", vals, t.Data)
	    }
	    isLi = false

        }
	}
	fmt.Println(vals)
}

