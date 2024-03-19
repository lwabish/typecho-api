package utils

import "regexp"

func Must(e error) {
	if e != nil {
		panic(e)
	}
}

func RemoveFrontMatter(markdown string) string {
	// 正则表达式匹配三个连字符之间的所有内容
	re := regexp.MustCompile(`^---\n([\s\S]*?)\n---\n`)
	return re.ReplaceAllString(markdown, "")
}
