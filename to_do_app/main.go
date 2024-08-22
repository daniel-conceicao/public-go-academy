package main

import (
	"goAcademy/todoApp/keyValueStore"
	"goAcademy/todoApp/webserver"
)

func main() {
	go keyValueStore.Run()
	webserver.Run()
}
