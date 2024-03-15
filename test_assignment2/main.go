package main

import (
	"test_assignment2/routers"
)

func main(){
	PORT := ":8080"

	routers.StartServer().Run(PORT)
}