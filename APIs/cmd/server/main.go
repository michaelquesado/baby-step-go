package main

import "github.com/michaelquesado/baby-step-go/APIs/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
