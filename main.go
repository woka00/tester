package main

import (
	"fmt"
	"git/feature1"
	"git/feature2"
	"git/feature_postgres/simple_connection"
)

func main() {
	fmt.Println("Hello, Git!")
	feature1.Feature1()
	feature2.Feature2()
	simple_connection.CheckConnection()
}
