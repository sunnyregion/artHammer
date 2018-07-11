package gofmt

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
)

func GetPath() string {
	return *path
}

// 判断是否是go文件
// 是go文件返回true，否则返回false
func IsGoFile(f os.FileInfo) bool {
	filename := f.Name()

	return !f.IsDir() &&
		!strings.HasPrefix(filename, ".") &&
		strings.HasSuffix(filename, ".go")
}

func WalkDir(walkFn filepath.WalkFunc) {
	parameters := flag.Args()

	if len(parameters) == 0 {
		parameters = append(parameters, *path)
	}

	for i := 0; i < len(parameters); i++ {
		path := parameters[i]

		switch dir, err := os.Stat(path); {
		case err != nil:
			report(err)
		case dir.IsDir():
			filepath.Walk(path, walkFn)
		default:
			walkFn(path, dir, nil)
		}
	}
}
