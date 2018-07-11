package gotool

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sunnyregion/artHammer/gotool/gofmt"
)

// GoLines getting the go file number and file lines number
type GoLines struct {
	Files int64 // files number
	Lines int64 // total lines number
}

// GoLines parameters
var (
	path = flag.String("path", `.`, `Use -p <Condition> 输入執行的目錄。`)
	ext  = flag.String("ext", `go`, `Use -e <Condition> 扩展名。`)
)

// Getting the file lines number and return the file lines number
func getFileLines(filename string) (count int64, err error) {
	fmt.Print(filename, "\t\t\t")

	count = 0
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			count++
		}
	}

	fmt.Println("line:", count)

	return
}

// Get the every file's lines number
func (g *GoLines) walkFileFunc(filename string, fi os.FileInfo, err error) error {
	if gofmt.IsGoFile(fi) {
		// getting the file lines number
		n, e := getFileLines(filename)

		if e == nil {
			g.Lines += n
			g.Files++
		}

		return e
	}

	return nil
}

// Implement IGoCheck.Check interface, getting the go files' lines number
func (g *GoLines) Check() {
	fmt.Println("Getting lines begin...")

	flag.Parse()

	gofmt.WalkDir(g.walkFileFunc)

	fmt.Println("total lines: ", g.Lines, "\t files:", g.Files)
	fmt.Println("Getting lines end\n")
	fmt.Println()
}
