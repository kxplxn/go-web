package _0405_dataStorage_

import (
	structs "go-web/05-dataStorage/01-structs"
	gobs "go-web/05-dataStorage/02-gobs"
	databases "go-web/05-dataStorage/03-databases"
	sqlMappers "go-web/05-dataStorage/04-sqlMappers"
)

func Demos() {
	structs.Demo()
	gobs.Demo()
	databases.Demo()
	sqlMappers.Demo()
}
