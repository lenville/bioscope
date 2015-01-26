package even

/*
 * go build & go install this file then "Even" can be used in other files
 */

// 函数首字母大写 - 可导出函数
func Even(i int) bool {
	return i%2 == 0
}

// 函数首字母小写 - 私有函数
func odd(i int) bool {
	return i%2 == 1
}
