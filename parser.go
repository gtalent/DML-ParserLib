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

func fixAlignment(text string) string {
	text = strings.Replace(text, "<center>", "<pre style=\"text-align:center\">", -1)
	text = strings.Replace(text, "</center>", "</pre>", -1)
	text = strings.Replace(text, "<right>", "<pre style=\"text-align:right\">", -1)
	text = strings.Replace(text, "</right>", "</pre>", -1)
	text = strings.Replace(text, "<left>", "<pre style=\"text-align:left\">", -1)
	text = strings.Replace(text, "</left>", "</pre>", -1)
	text = strings.Replace(text, "<justify>", "<pre style=\"text-align:justify\">", -1)
	text = strings.Replace(text, "</justify>", "</pre>", -1)
	return string(text)
}

/*
 * Swaps the DML tags in the given input for appropriate HTML tags in the return value.
 * Takes:
 *      doc - the text of the document to convert
 * Returns:
 *      the given text with all special characters escaped and the DML tags swapped for the appropriate HTML tags
 */
func ParseDoc(text string) string {
	text = fixAlignment(text)
	return text
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
