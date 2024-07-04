package main

import (
	"fmt"

	"github.com/extractings/meds-reminder/internals/api"
	"github.com/extractings/meds-reminder/master"
)

func init() {
	// Connnections and initializers
	api.LoadEnvVariables()
	api.LoadRedis()

	// Main App 
	master.LoadMedsApp()
}

func main() {
	fmt.Println("start of something cool fr fr")
}