//解析格式化go
// 1.去掉每行逗号之后的空格使之成为 ,//...\n {//...\n [//...\n, }//...\n ]//...\n紧密相接的格式
// 2.去掉每行末尾与之匹配的格式
package main

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
		if index > 1 {
			//正常末尾注释
			if inKeySymbol(content[index-1]) && singleChar == '/' {
				end = index
				result = append(result, content[start:end]...)
			}
			//单行注释
			if singleChar == '\n' {
				last = start
				start = index
				// 本行无注释 两个 \n 中间不存在正常末尾注释的情况
				if end < last {
					result = append(result, content[last:start]...)
				}
			}
		}
	}

	if start == 0 {
		result = content
	} else {
		result = append(result, content[len(content)-1])
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
