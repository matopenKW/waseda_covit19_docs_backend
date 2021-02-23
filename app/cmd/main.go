package main

import (
	"flag"
	"log"

	_ "github.com/lib/pq"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/cmd/upload"
)

func main() {
	flag.Parse()
	exec(flag.Args())
}

func exec(args []string) {
	if len(args) == 0 {
		panic("not args")
	}

	var err error
	switch args[0] {
	case "upload_histories_csv":
		err = upload.UploadHistoriesCsv(args[1:]...)
	}

	if err != nil {
		panic(err.Error())
	}

	log.Println("Succsess")
}
