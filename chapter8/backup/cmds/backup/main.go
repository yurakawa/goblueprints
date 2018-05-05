package main

import (
	"errors"
	"flag"
	"log"
	"strings"

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
	case "add":
	case "remove":
	}
}
