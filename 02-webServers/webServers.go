package _0302_webServers_

import (
	createWebServer "go-web/02-webServers/01-createWebServer"
	createFileServer "go-web/02-webServers/02-createFileServer"
	handlers "go-web/02-webServers/03-handlers"
	multipleHandlers "go-web/02-webServers/04-multipleHandlers"
)

func Demos() {
	createWebServer.Demo()
	createFileServer.Demo()
	handlers.Demo()
	multipleHandlers.Demo()
}
