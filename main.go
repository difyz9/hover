package main

import (
	"github.com/go-flutter-desktop/hover/cmd"
	"fmt"
)

var VERSION = "v0.0.1"
var AUTHOR = "dell"
var buildTime = "2025-04-01 10:50:01"

func main() {

    fmt.Println("version :", VERSION)
    fmt.Println("author :", AUTHOR)
    fmt.Println("buildTime :", buildTime)
	cmd.Execute()
}
