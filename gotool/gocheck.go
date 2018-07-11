package gotool

import (
	"fmt"

	"github.com/sunnyregion/artHammer/gotool/gofmt"
	"github.com/sunnyregion/artHammer/gotool/golint"
	"github.com/sunnyregion/artHammer/gotool/govet"
)

// IGoCheck check interface
type IGoCheck interface {
	Check()
}

// GoFmt formatting the go files
type GoFmt struct{}

func (g *GoFmt) Check() {
	fmt.Println("Formatting file begin...")

	gofmt.GoFmtMain()

	fmt.Println("Formatting file end")
	fmt.Println()
}

// GoVet go vet tool
type GoVet struct{}

func (g *GoVet) Check() {
	fmt.Println("Govet begin...")

	govet.GoVetMain()

	fmt.Println("Govet end")
	fmt.Println()
}

// GoLint golint tool
type GoLint struct{}

func (g *GoLint) Check() {
	fmt.Println("Golint begin...")

	golint.GoLintMain()

	fmt.Println("Golint end")
	fmt.Println()
}
