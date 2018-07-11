package main

import (
	"fmt"

	"github.com/sunnyregion/artHammer/gotool"
)

/**
 * Main function
 */
func main() {
	fmt.Println("Checking begin...")

	goFmt := gotool.GoFmt{}
	goVet := gotool.GoVet{}
	goLint := gotool.GoLint{}
	goLines := gotool.GoLines{0, 0}

	goFmt.Check()
	goVet.Check()
	goLint.Check()
	goLines.Check()

	fmt.Println("Checking end\n")
	fmt.Println()
}
