package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	INCREMENT  = "歩夢ちゃん"    // +
	DECREMENT  = "四季ちゃん"    // -
	NEXT       = "ルビィちゃん"   // >
	PREVIOUS   = "アイスクリーム!" // <
	LOOP_START = "はーい"      // [
	LOOP_END   = "何が好き?"    // ]
	READ       = "あ・な・た"    // ,
	WRITE      = "叫びましょ"    // .
)

func main() {
	// 引数が1つ(プログラムファイル)でなければ使用法を表示して終了
	if len(os.Args) == 2 {
		bytes, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		code := string(bytes)

		// 改行と空白を削除
		code = strings.ReplaceAll(code, "\n", "")
		re := regexp.MustCompile(`\s+`)
		code = re.ReplaceAllString(code, "")

		runes := []rune(string(code))
		// 認識する命令を並べたリスト
		programs := []string{INCREMENT, DECREMENT, NEXT, PREVIOUS, LOOP_START, LOOP_END, READ, WRITE}
		// トークン化された命令列
		tokens := []string{}

		// それぞれの命令文字列が出現するかを確認し、マッチしたらトークンとして追加
		for i := 0; i < len(runes); i++ {
			for _, program := range programs {
				// 最初の1文字が一致するか確認
				if string(runes[i]) == string([]rune(program)[0]) {
					// program と同じ長さの文字列を切り取って比較
					text := string(runes[i : i+len([]rune(program))])
					if text == string([]rune(program)) && i+len([]rune(program)) <= len(runes) {
						tokens = append(tokens, text)
					}
				}
			}
		}

		buffer := make([]byte, 30000)
		ptr := 0 // メモリポインタ

		// トークンを順番に実行
		for i := 0; i < len(tokens); i++ {
			token := tokens[i]

			switch token {
			case INCREMENT:
				// 現在のセルをインクリメント
				buffer[ptr]++

			case DECREMENT:
				// 現在のセルをデクリメント
				buffer[ptr]--

			case NEXT:
				// ポインタを右へ
				if ptr < len(buffer) {
					ptr++
				}

			case PREVIOUS:
				// ポインタを左へ
				if ptr > 0 {
					ptr--
				}

			case LOOP_START:
				// [ と同じ：現在のセルが 0 なら対応する ] の位置までジャンプ
				if buffer[ptr] == 0 {
					n := 0
					for {
						i++
						if tokens[i] == LOOP_START {
							n++
						} else if tokens[i] == LOOP_END {
							// ネストが 0 のときが本来の対応括弧
							if n == 0 {
								break
							}
							n--
						}
					}
				}

			case LOOP_END:
				// ] と同じ：現在のセルが 0 でなければ対応する [ まで戻る
				if buffer[ptr] != 0 {
					n := 0
					for {
						i--
						if tokens[i] == LOOP_END {
							n++
						} else if tokens[i] == LOOP_START {
							// ネストが 0 のときが対応する括弧
							if n == 0 {
								break
							}
							n--
						}
					}
				}

			case READ:
				// 入力1文字を読み取って現在のセルに格納
				buf := []byte{0}
				os.Stdin.Read(buf)
				buffer[ptr] = buf[0]

			case WRITE:
				// 現在のセルの内容を1文字出力
				fmt.Print(string(buffer[ptr]))
			}
		}
	} else {
		fmt.Println("Usage: aiscream program.aiscream")
	}
}
