package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sipt/GoJsoner"
)

const help = `
本命令行只支持一个参数 这个参数必须是合法文件名
`

func main() {
	var originSourceFileName string

	flag.Usage = func() {
		fmt.Println(help)
	}

	if len(os.Args) == 1 {
		fmt.Println("请输入文件名")
		return
	}

	originSourceFileName = os.Args[1]

	contentBytes, fileError := ioutil.ReadFile(originSourceFileName)

	//非法文件或者不存在的文件
	if fileError != nil {
		fmt.Print(fileError)
	}

	content := string(contentBytes)

	//借用外部包删除注释
	result, err := GoJsoner.Discard(content)

	newContentBytes := remoteInvalidComma(result)

	//输出结果到新文件
	ioutil.WriteFile("json_"+originSourceFileName, newContentBytes, 0777)

	//打印结果
	fmt.Println(string(newContentBytes))
	if err != nil {
		fmt.Println(err)
	}
}

//删除不合文法的末尾逗号 使之符合正确的json格式
func remoteInvalidComma(result string) []byte {
	//逗号右边的符号判断
	symbol := [2]byte{'}', ']'}
	var newContentBytes []byte
	newContentBytes = []byte(result)
	for po, ch := range newContentBytes {
		if ch == ',' && inArray(newContentBytes[po+1], symbol) {
			newContentBytes[po] = 0
		}
	}
	return newContentBytes
}

//校验右侧符号
func inArray(needle byte, array [2]byte) bool {
	for _, item := range array {
		if item == needle {
			return true
		}
	}
	return false
}