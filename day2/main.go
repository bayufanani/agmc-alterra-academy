package main

import (
	"agmc/day2/config"
	"agmc/day2/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
