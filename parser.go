package dml

import (
	"io/ioutil"
	"strings"
)

var defaultCSS = `
	h1 {
		text-align:center
	}
	h2 {
		text-align:center
	}
	h3 {
		text-align:center
	}
	h4 {
		text-align:center
	}
	h5 {
		text-align:center
	}
	h6 {
		text-align:center
	}
	pre {
		white-space: pre-wrap;
	}`

func LoadDefaultCSS(file string) bool {
	text, err := ioutil.ReadFile("default.css")
	if err != nil {
		return false
	}
	defaultCSS = string(text)
	return true
}

/*
 * Swaps the DML tags in the given input for appropriate HTML tags in the return value.
 * Takes:
 *      doc - the text of the document to convert
 * Returns:
 *      the given text with all special characters escaped and the DML tags swapped for the appropriate HTML tags
 */
func ParseDoc(text string) string {
	text = strings.Replace(text, "<center>", "<p style=\"text-align:center\">", -1)
	text = strings.Replace(text, "</center>", "</p>", -1)
	text = strings.Replace(text, "<right>", "<p style=\"text-align:right\">", -1)
	text = strings.Replace(text, "</right>", "</p>", -1)
	text = strings.Replace(text, "<left>", "<p style=\"text-align:left\">", -1)
	text = strings.Replace(text, "</left>", "</p>", -1)
	text = strings.Replace(text, "<justify>", "<p style=\"text-align:justify\">", -1)
	text = strings.Replace(text, "</justify>", "</p>", -1)
	return string(text)
}

/*
 * Swaps the DML tags in the given input for appropriate HTML tags and embed the new HTML in HTML tags as the return value.
 * Takes:
 *      doc - the text of the document to convert
 * Returns:
 *      the given text with all special characters escaped and the DML tags swapped for the appropriate HTML tags
 */
func ToHTML(title, input string) string {
	file, err := ioutil.ReadFile("default.css")
	var css string
	if err != nil {
		css = defaultCSS
	} else {
		css = string(file)
	}
	top := "<html><head><style type=\"text/css\">" + css + "</style><link rel=\"stylesheet\" type=\"text/css\" href=\"default.css\" /><title>" + title + "</title></head><body>"
	bottom := "</body></html>"
	return top + ParseDoc(input) + bottom
}
