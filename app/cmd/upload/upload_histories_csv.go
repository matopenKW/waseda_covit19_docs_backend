package upload

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
	"google.golang.org/api/drive/v3"
)

func UploadHistoriesCsv(args ...string) error {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer db.Close()

	err = execute(db)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func execute(db *gorm.DB) error {
	repo := repository.NewGoogleDriveRepository()

	client, err := repo.GetClient()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	srv, err := client.GetService()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// TODO 前回登録したファイルIDを取得する
	beforeDriveID := "14LFyvFU8BxCtYpZytbM96AKa-zeJvSIV"
	err = srv.Delete(beforeDriveID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// TODO　一時的にテキストを取得
	buf, _ := ioutil.ReadFile("test.csv")

	r := bytes.NewReader(buf)
	ret, err := srv.Create(r, &drive.File{
		OriginalFilename: "test.csv",
		Name:             "test.csv",
		Description:      "test",
		Parents:          []string{"1i12pYvnABx5Ov_xV53IDufOJhSZrBwDu"},
		MimeType:         "text/csv",
	})
	if err != nil {
		return err
	}

	fmt.Println(ret)
	return nil
}
