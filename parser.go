package dml

import (
	"io/ioutil"
	"strings"
)

var defaultCSS = `
	pre {
		white-space: pre-wrap;
	}`
var defaultBottom = "</body></html>"

/*
 * Returns the default top of the document.
 * Takes:
 *	css - the CSS settings for this document
 * Returns:
 *	the default top of the document
 */
func defaultTop(title, css string) string {
	return "<html><head><style type=\"text/css\">" + css + "</style><link rel=\"stylesheet\" type=\"text/css\" href=\"default.css\" /><title>" + title + "</title></head><body>"
}

/*
 * Gets the default CSS for the documents.
 * Returns:
 * 	the default CSS for the documents.
 */
func getCSS() string {
	text, err := ioutil.ReadFile("default.css")
	if err != nil {
		return defaultCSS
	}
	return string(text)
}

func getTop(title, css string) string {
	file, err := ioutil.ReadFile("top.html")
	if err != nil {
		return defaultTop(title, css)
	}
	return string(file)
}

func getBottom() string {
	file, err := ioutil.ReadFile("bottom.html")
	if err != nil {
		return defaultBottom
	}
	return string(file)
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
	text = "<pre>" + fixAlignment(text) + "</pre>"
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
	return getTop(title, getCSS()) + ParseDoc(input) + getBottom()
}

/*
 * Swaps the DML tags in the given input for appropriate HTML tags and embed the new HTML in HTML tags as the return value.
 * Takes:
 *      doc - the text of the document to convert
 *	css - the text of css document
 * Returns:
 *      the given text with all special characters escaped and the DML tags swapped for the appropriate HTML tags
 */
func ToHTMLCSS(title, input, css string) string {
	return getTop(title, css)+ ParseDoc(input) + getBottom()
}
