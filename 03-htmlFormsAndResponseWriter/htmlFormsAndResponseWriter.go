package _0303_htmlFormsAndResponseWriter_

import (
	htmlForms "go-web/03-htmlFormsAndResponseWriter/01-htmlForms"
	responseWriter "go-web/03-htmlFormsAndResponseWriter/02-responseWriter"
	responseHeaders "go-web/03-htmlFormsAndResponseWriter/03-responseHeaders"
	returningJSON "go-web/03-htmlFormsAndResponseWriter/04-returningJSON"
)

func Demos() {
	htmlForms.Demo()
	responseWriter.Demo()
	responseHeaders.Demo()
	returningJSON.Demo()
}
