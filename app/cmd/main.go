package main

import (
	"flag"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/cmd/service"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

func main() {
	flag.Parse()
	if err := exec(flag.Args()); err != nil {
		log.Println(err.Error())
		panic("error")
	}
	log.Println("Succsess")
}

func exec(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("not args")
	}

	db, err := repository.NewDbConnection()
	if err != nil {
		return fmt.Errorf("db error; %s", err.Error())
	}
	defer db.Close()

	srv, err := GetImpl(args[0])
	if err != nil {
		return err
	}

	err = srv.SetParam(args[1:])
	if err != nil {
		return err
	}

	err = srv.Validate()
	if err != nil {
		return err
	}

	err = srv.Execute()
	if err != nil {
		return err
	}
	return nil
}

func GetImpl(pm string) (service.CmdServiceImpl, error) {
	switch pm {
	case "upload_histories_csv":
		return service.NewUploadHistoriesCsvImpl()

	default:
		return nil, fmt.Errorf(fmt.Sprintf("not service pm=%s", pm))
	}
}
