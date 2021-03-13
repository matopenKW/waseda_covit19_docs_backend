package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
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
	db.LogMode(true)
	defer db.Close()

	srv, err := GetImpl(db, args[0])
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

func GetImpl(db *gorm.DB, pm string) (service.CmdServiceImpl, error) {
	switch pm {
	case "upload_histories_csv":
		if os.Getenv("GOOGLE_DRIVE_WORK_FOLDER") == "" {
			return nil, fmt.Errorf("gd work folder env empty")
		}
		return service.NewUploadHistoriesCsvImpl(db)
	default:
		return nil, fmt.Errorf(fmt.Sprintf("not service pm=%s", pm))
	}
}
