package main

import (
	"fmt"
	"github.com/gibraan/GolangTrain/05_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
