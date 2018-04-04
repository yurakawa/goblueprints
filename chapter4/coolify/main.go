package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	duplicateVowel = true  // 母音を重ねる
	removeVowel    = false // 母音を削除する
)

// randBool trueかfalseをランダムで返す
func randBool() bool {
	return rand.Intn(2) == 0
}

// coolifyが返還を行う確率は50%
//
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := []byte(s.Text())
		if randBool() {
			var vI = -1
			for i, char := range word {
				switch char {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					if randBool() {
						vI = i
					}
				}
			}
			if vI >= 0 {
				switch randBool() {
				case duplicateVowel:
					word = append(word[:vI+1], word[vI:]...) // スライスに続けてドットを3つ記述すると、スライス中の各項目を独立した引数として渡すことが出来る
				case removeVowel:
					word = append(word[:vI], word[vI+1:]...)
				}
			}
		}
		fmt.Println(string(word))
	}
}
