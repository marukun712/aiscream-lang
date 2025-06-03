package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	INCREMENT  = "歩夢ちゃん"
	DECREMENT  = "四季ちゃん"
	NEXT       = "ルビィちゃん"
	PREVIOUS   = "アイスクリーム!"
	LOOP_START = "はーい"
	LOOP_END   = "何が好き?"
	READ       = "あ・な・た"
	WRITE      = "叫びましょ"
)

func main() {
	if len(os.Args) == 2 {
		bytes, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		code := string(bytes)

		code = strings.ReplaceAll(code, "\n", "")
		re := regexp.MustCompile(`\s+`)
		code = re.ReplaceAllString(code, "")

		code = strings.ReplaceAll(code, "+", INCREMENT)
		code = strings.ReplaceAll(code, "-", DECREMENT)
		code = strings.ReplaceAll(code, ">", NEXT)
		code = strings.ReplaceAll(code, "<", PREVIOUS)
		code = strings.ReplaceAll(code, "[", LOOP_START)
		code = strings.ReplaceAll(code, "]", LOOP_END)
		code = strings.ReplaceAll(code, ",", READ)
		code = strings.ReplaceAll(code, ".", WRITE)

		fmt.Print(code)

		os.WriteFile("dist/program.aiscream", []byte(code), 0644)
	} else {
		fmt.Println("Usage: parser program.bf")
	}
}
