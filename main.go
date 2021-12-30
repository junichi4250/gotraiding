package main

import (
	"fmt"
	"gotraiding/config"
	"gotraiding/utils"
	"log"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.LogFile)

	utils.LoggingSetting(config.Config.LogFile)
	log.Println("test")
}
