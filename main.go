package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mozillazg/go-pinyin"
)

// 将中文字符串转换为拼音
func convertToPinyin(text string) string {
	// 默认
	a := pinyin.NewArgs()
	res := ""
	for _, v := range pinyin.Pinyin(text, a) {
		res += v[0]
	}
	return res
}

func openTxt2Ennames(prefix *string) {
	// 打开文件
	file, err := os.Open("names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建文件
	f, err := os.Create("names_en.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 读取文件
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := scanner.Text()
		pinyin := convertToPinyin(name)
		if prefix != nil {
			fmt.Fprintln(f, *prefix+pinyin)
		} else {
			fmt.Fprintln(f, pinyin)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	prefix := flag.String("prefix", "", "prefix for pinyin names")
	flag.Parse()
	openTxt2Ennames(prefix)
}
