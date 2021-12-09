package _0301_basics_

import (
	httpRequests "go-web/01-basics/01-httpRequests"
	httptestPackage "go-web/01-basics/02-httptestPackage"
	structuringApps "go-web/01-basics/03-structuringApps"
)

func Demos() {
	httpRequests.Demo()
	httptestPackage.DemoTestHTTPCalls()
	httptestPackage.DemoUseTestServer()
	structuringApps.Demo()
}
