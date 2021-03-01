package service

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"google.golang.org/api/drive/v3"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

const (
	ParentsPath     = "1_C0EVIK-WeK0WSMF1eIDOXLkE6UhnmNJ"
	HistoryFileName = "history_%s.csv"
	MasterDataID    = "1sJ7V19QDgN4o9BNrUpBE-G_MUYmoqiVZEWl7mmNyCRU"
)

type UploadHistoriesCsvImpl struct {
	dbRepo repository.Repository
	gdRepo repository.GoogleDriveRepository
}

func NewUploadHistoriesCsvImpl(db *gorm.DB) (*UploadHistoriesCsvImpl, error) {
	return &UploadHistoriesCsvImpl{
		dbRepo: repository.NewDbRepository(db),
		gdRepo: repository.NewGoogleDriveRepository(),
	}, nil
}

func (s *UploadHistoriesCsvImpl) SetParam(args []string) error {
	return nil
}

func (s *UploadHistoriesCsvImpl) Validate() error {
	return nil
}

func (s *UploadHistoriesCsvImpl) Execute() error {
	con, err := s.dbRepo.NewConnection()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	client, err := s.gdRepo.GetClient()
	if err != nil {
		return err
	}
	srv, err := client.GetService()
	if err != nil {
		return err
	}

	lu, err := con.LatestLastUpload()
	if err != nil {
		return err
	}

	err = srv.Delete(lu.DriveID)
	if err != nil {
		log.Println(err.Error())
	}

	// TODO　一時的にテキストを取得
	//	buf, _ := ioutil.ReadFile("test.csv")
	buf, err := createHistoryFile()
	if err != nil {
		return err
	}
	r := bytes.NewReader(buf)

	weekDay := time.Now().Weekday()
	ret, err := srv.Create(r, &drive.File{
		OriginalFilename: fmt.Sprintf(HistoryFileName, weekDay.String()),
		Name:             fmt.Sprintf(HistoryFileName, weekDay.String()),
		Description:      "test",
		Parents:          []string{ParentsPath},
		MimeType:         "text/csv",
	})
	if err != nil {
		return err
	}

	err = con.RunTransaction(func(tx repository.Transaction) error {
		err := tx.UpdateLastUpload(&model.LastUpload{
			DriveID: ret.Id,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func createHistoryFile() ([]byte, error) {

	// file, err := srv.Files.Get("1sJ7V19QDgN4o9BNrUpBE-G_MUYmoqiVZEWl7mmNyCRU").Do()
	// if err != nil {
	// 	return nil, err
	// }
	// log.Println(file)

	testData := []string{"1", "1", "1", "1", "1"}

	var csv string
	for _, s := range testData {
		csv += s + ","
	}
	return []byte(csv), nil
}
