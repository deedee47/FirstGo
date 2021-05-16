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

	//Variable declararion - var varName dataType
	var sampleData1 string = "Hello Deedee ";
	var number int = 8642943024;

	fmt.Printf("%s, %d", sampleData1, number);

	//Use other functions
	randomDataLib();
}

func randomDataLib(){
        fmt.Println("Running the TestApp")
        fmt.Println(randomdata.SillyName())
}