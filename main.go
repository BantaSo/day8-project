package main

import (
	"day8-project/config"
	"day8-project/route"
)

func main() {
	config.StartDB()
	route.StartRoute().Run(":8080")
}
