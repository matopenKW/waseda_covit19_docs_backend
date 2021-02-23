package main

import (
	"flag"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/cmd/upload"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

func main() {
	flag.Parse()
	exec(flag.Args())
}

func exec(args []string) {
	if len(args) == 0 {
		panic("not args")
	}

	db, err := repository.NewDbConnection()
	if err != nil {
		panic(fmt.Sprintf("db error; %s", err.Error()))
	}
	defer db.Close()

	switch args[0] {
	case "upload_histories_csv":
		err = upload.UploadHistoriesCsv(db, args[1:]...)
	}

	if err != nil {
		panic(err.Error())
	}

	log.Println("Succsess")
}
