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

		runes := []rune(string(code))

		programs := []string{INCREMENT, DECREMENT, NEXT, PREVIOUS, LOOP_START, LOOP_END, READ, WRITE}

		tokens := []string{}

		for i := 0; i < len(runes); i++ {
			for _, program := range programs {
				if string(runes[i]) == string([]rune(program)[0]) {
					text := string(runes[i : i+len([]rune(program))])
					if text == string([]rune(program)) && i+len([]rune(program)) <= len(runes) {
						tokens = append(tokens, text)
					}
				}
			}
		}

		buffer := make([]byte, 30000)
		ptr := 0

		for i := 0; i < len(tokens); i++ {
			token := tokens[i]

			switch token {
			case INCREMENT:
				buffer[ptr]++
			case DECREMENT:
				buffer[ptr]--
			case NEXT:
				if ptr < len(buffer) {
					ptr++
				}
			case PREVIOUS:
				if ptr > 0 {
					ptr--
				}
			case LOOP_START:
				if buffer[ptr] == 0 {
					n := 0
					for {
						i++
						if tokens[i] == LOOP_START {
							n++
						} else if tokens[i] == LOOP_END {
							if n == 0 {
								break
							}
							n--
						}
					}
				}
			case LOOP_END:
				if buffer[ptr] != 0 {
					n := 0
					for {
						i--
						if tokens[i] == LOOP_END {
							n++
						} else if tokens[i] == LOOP_START {
							if n == 0 {
								break
							}
							n--
						}
					}
				}
			case READ:
				buf := []byte{0}
				os.Stdin.Read(buf)
				buffer[ptr] = buf[0]
			case WRITE:
				fmt.Print(string(buffer[ptr]))
			}
		}
	} else {
		fmt.Println("Usage: aiscream program.aiscream")
	}
}
