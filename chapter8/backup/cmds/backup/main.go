package main

import (
	"errors"
	"flag"
	"log"
	"strings"

	"fmt"

	"github.com/hashicorp/packer/common/json"
	"github.com/matryer/filedb"
)

/*
  backup command
  usage:
    backup -db=./backupdata.db add {path} [{path} {path}...]
    backup -db=./backupdata.db remove {path} [{path} {path}...]
    backup -db=./backupdata.db list
*/

type path struct {
	Path string
	Hash string
}

func (p path) String() string {
	return fmt.Sprintf("%s [%s]", p.Path, p.Hash)
}

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			flag.PrintDefaults()
			log.Fatalln(fatalErr)
		}
	}()

	var (
		dbpath = flag.String("db", "./backupdata", "データベースのディレクトリへのパス")
	)
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fatalErr = errors.New("エラー; コマンドを指定してください")
		return
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

	switch strings.ToLower(args[0]) {
	case "list":
		var path path
		col.ForEach(func(i int, data []byte) bool {
			err := json.Unmarshal(data, &path)
			if err != nil {
				fatalErr = err
				return true
			}
			fmt.Printf("= %s\n", path)
			// Returning false is the result of following the API of filedb.
			// When returning true, processing to subsequent items is canceled,
			// so false is returned so that all items are displayed.
			return false
		})
	case "add":
		if len(args[1:]) == 0 {
			fatalErr = errors.New("追加するパスを指定してください")
			return
		}
		for _, p := range args[1:] {
			path := &path{Path: p, Hash: "まだアーカイブされていません"}
			if err := col.InsertJSON(path); err != nil {
				fatalErr = err
				return
			}
			fmt.Printf("+ %s\n", path)
		}
	case "remove":
	}
}
