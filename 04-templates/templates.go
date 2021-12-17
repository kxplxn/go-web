package _0304_templates_

import (
	parsingAndExecuting "go-web/04-templates/01-parsingAndExecuting"
	actions "go-web/04-templates/02-actions"
	arguments "go-web/04-templates/03-arguments"
	flags "go-web/04-templates/04-flags"
	variables "go-web/04-templates/05-variables"
	pipelines "go-web/04-templates/06-pipelines"
	functions "go-web/04-templates/07-functions"
)

func Demos() {
	parsingAndExecuting.Demo()
	actions.Demo()
	arguments.Demo()
	flags.Demo()
	variables.Demo()
	pipelines.Demo()
	functions.Demo()
}
