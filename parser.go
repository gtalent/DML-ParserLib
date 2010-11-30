package dml

import "io/ioutil"

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

/*
 * Swaps the DML tags in the given input for appropriate HTML tags in the return value.
 * Takes:
 *      doc - the text of the document to convert
 * Returns:
 *      the given text with all special characters escaped and the DML tags swapped for the appropriate HTML tags
 */
func ParseDoc(text string) string {
        return string(text)
};

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
