package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/atotto/clipboard"
	//"github.com/sipt/GoJsoner"
)

var h bool

const help = `
本命令行只支持单个参数 这个参数必须是合法文件名
`

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.Usage = usage
}

func usage() {
	//只有输入 -h 才可以被展示
	if h {
		fmt.Fprintf(os.Stderr, `帮助说明:`+help)
		flag.PrintDefaults()
	}
}

func main() {
	var originSourceFileName string
	flag.Parse()
	flag.Usage()

	time.Now()
	//if len(os.Args) == 1 {
	//	fmt.Println("请输入文件名,仅支持单个")
	//	return
	//}
	var content string

	if len(os.Args) > 1 {
		originSourceFileName = os.Args[1]
		contentBytes, _ := ioutil.ReadFile(originSourceFileName)
		content = string(contentBytes)
	} else {
		//非法文件或者不存在的文件 就去读取剪切板
		content, _ = clipboard.ReadAll()
	}

	//借用外部包删除注释
	//result, err := GoJsoner.Discard(content)

	var first = removeSpace([]byte(content)) //fmt.Println(first)

	result := removeComments(first)

	newContentBytes := remoteInvalidComma(string(result))

	//输出结果到新文件
	//ioutil.WriteFile("json_"+originSourceFileName, newContentBytes, 0777)

	//打印结果
	fmt.Println(string(newContentBytes))

}

//删除不合文法的末尾逗号 使之符合正确的json格式
func remoteInvalidComma(result string) []byte {
	//逗号右边的符号判断
	symbol := [2]byte{'}', ']'}
	var newContentBytes []byte
	newContentBytes = []byte(result)
	for po, ch := range newContentBytes {
		if ch == ',' && inArray(newContentBytes[po+1], symbol) {
			newContentBytes = append(newContentBytes[:po], newContentBytes[po+1:]...)
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

var keySymbol []byte = []byte{'[', '{', '}', ']', '"', ',', '\n'}

//去掉空格
func removeSpace(content []byte) []byte {
	var result []byte
	var canAppendSpace byte
	for _, singleChar := range content {
		if singleChar == '"' {
			canAppendSpace = canAppendSpace ^ 1
		}
		if canAppendSpace == 0 && singleChar == ' ' || singleChar == '\t' {
			continue
		}
		result = append(result, singleChar)
	}
	return result
}

//去掉注释 和 换行符
func removeComments(content []byte) []byte {
	var result []byte
	start := 0
	end := 0
	last := 0
	for index, singleChar := range content {

		//正常末尾注释
		if index > 0 && inKeySymbol(content[index-1]) && singleChar == '/' {
			end = index
			result = append(result, content[start:end]...)
		}
		//单行注释 以及空行情况
		if singleChar == '\n' {
			last = start
			start = index
			// 本行无注释 两个 \n 中间不存在正常末尾注释的情况
			if end <= last {
				result = append(result, content[last:start]...)
			}
		}
	}

	//去掉换行符
	for index, line := range result {
		if line == '\n' {
			result = append(result[:index], result[index+1:]...)
		}
	}

	//拼接尾字符
	result = append(result, content[len(content)-1])

	if start == 0 {
		result = content
	}
	return result
}

func inKeySymbol(needle byte) bool {
	for _, item := range keySymbol {
		if item == needle {
			return true
		}
	}
	return false
}
