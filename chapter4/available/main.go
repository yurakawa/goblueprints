package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// 指定されたドメイン名の詳細についてWHOISサーバに問い合わせる
// ->情報なし未使用
// しかしwhoisサーバごとに形式が異なる

// exists whoisServerで指定されたサーバのポート43に対して接続を開く
//        接続先にドメイン名と\r\n(復帰と改行)を送信
func exists(domain string) (bool, error) {
	const whoisServer = "com.whois-servers.net"
	conn, err := net.Dial("tcp", whoisServer+":43")
	if err != nil {
		return false, err
	}
	defer conn.Close()
	conn.Write([]byte(domain + "\r\n"))
	scanner := bufio.NewScanner(conn)
	// レスポンスを1行ずつ読み込む
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), "no match") {
			return false, nil
		}
	}
	return true, nil
}

var marks = map[bool]string{true: "○", false: "☓"}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		domain := s.Text()
		fmt.Print(domain, " ")
		exist, err := exists(domain)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(marks[!exist])
		time.Sleep(1 * time.Second)

	}
}
