package main

import (
	"fmt"

	randomdata "github.com/Pallinder/go-randomdata"
)

//helps to display on console

//entry point to program; auto called when the fo file is
//executed
func main(){
	fmt.Println("Home Page");
	randomDataLib();
}

func randomDataLib(){
        fmt.Println("Running the TestApp")
        fmt.Println(randomdata.SillyName())
}