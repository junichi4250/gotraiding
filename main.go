package main

import (
	"fmt"
	"gotraiding/app/controllers"
	"gotraiding/config"
)

func main() {
	fmt.Println(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
