package main

import (
	"flag"
	"log"

	"errors"

	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hashicorp/packer/common/json"
	"github.com/matryer/filedb"
	"github.com/yurakawa/goblueprints/chapter8/backup"
)

type path struct {
	Path string
	Hash string
}

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			log.Fatal(fatalErr)
		}
	}()

	var (
		interval = flag.Int("interval", 10, "チェックの間隔(秒単位)")
		archive  = flag.String("archive", "archive", "アーカイブの保存先")
		dbpath   = flag.String("db", "./db", "filedbデータベースへのパス")
	)
	flag.Parse()

	m := &backup.Monitor{
		Destination: *archive,
		Archiver:    backup.ZIP,
		Paths:       make(map[string]string),
	}

	db, err := filedb.Dial(*dbpath)
	if err != nil {
		fatalErr = err
		return
	}
	defer db.Close()
	col, err := db.C("paths")
	if err != nil {
		fatalErr = err
		return
	}

	var path path
	col.ForEach(func(_ int, data []byte) bool {
		if err := json.Unmarshal(data, &path); err != nil {
			fatalErr = err
			return true
		}
		m.Paths[path.Path] = path.Hash
		return false //胥吏を続行する
	})
	if fatalErr != nil {
		return
	}
	if len(m.Paths) < 1 {
		fatalErr = errors.New("パスがありません。backup ツールを使って追加してください")
		return
	}

	check(m, col)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
Loop:
	for {
		select {
		case <-time.After(time.Duration(*interval) * time.Second):
			check(m, col)
		case <-signalChan:
			// 終了
			fmt.Println()
			log.Printf("終了します...")
			break Loop
		}
	}
}
