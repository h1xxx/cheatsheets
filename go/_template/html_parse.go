package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

const HTML = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8"/>
        <title>selected attribute</title>
    </head>
    <body>
        <form method="GET">
            <input type="submit" value="submit"/>
        </form>
    </body>
</html>
`

func main() {
	z, _ := html.Parse(strings.NewReader(HTML))
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "input" {
			for _, a := range n.Attr {
				if a.Key == "value" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(z)
}
