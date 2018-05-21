package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	sum       int64
	path, ext string
)

func init() {
	var sPath = flag.String("p", `./`, `Use -p <Condition> 输入執行的目錄。`)
	var sExt = flag.String("e", `go`, `Use -e <Condition> 扩展名。`)
	flag.Parse()
	path = *sPath
	ext = *sExt
}

/**
 * 主程序
 */
func main() {
	files := 0
	filepath.Walk(path, func(filename string, info os.FileInfo, err error) error {
		if getExt(filename) == ext {
			//fmt.Println(path)
			n, e := countFileLine(filename)
			if e == nil {
				//fmt.Println("lines: ", n)
				sum = sum + n
				files = files + 1
			}
		}
		return nil
	})
	fmt.Println("total lines: ", sum, "\t files:", files)
}

//获取得文件的扩展名，最后一个.后面的内容
func getExt(f string) (ext string) {
	index := strings.LastIndex(f, ".")
	data := []byte(f)
	for i := index + 1; i < len(data); i++ {
		ext = ext + string([]byte{data[i]})
	}
	return
}

//计算文件行数
func countFileLine(name string) (count int64, err error) {
	fmt.Print(name, "\t\t\t")
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return
	}
	count = 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			count++
		}
	}
	fmt.Println("line:", count)
	return
}
