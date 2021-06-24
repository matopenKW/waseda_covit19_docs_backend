package service

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"google.golang.org/api/drive/v3"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

const (
	HistoryFileName = "history_%s.csv"
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

	buf, err := createHistoryFile(con)
	if err != nil {
		return err
	}
	r := bytes.NewReader(buf)
	weekDay := time.Now().Weekday()
	ret, err := srv.Create(r, &drive.File{
		OriginalFilename: fmt.Sprintf(HistoryFileName, weekDay.String()),
		Name:             fmt.Sprintf(HistoryFileName, weekDay.String()),
		Description:      "test",
		Parents:          []string{os.Getenv("GOOGLE_DRIVE_WORK_FOLDER")},
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

func createHistoryFile(con repository.Connection) ([]byte, error) {

	us, err := con.ListUser()
	if err != nil {
		return nil, err
	}

	uMap := make(map[model.UserID]*model.User)
	for _, v := range us {
		uMap[v.ID] = v
	}

	var csv string
	header := []string{"氏名", "日時", "内容", "経路（行き）", "経路（帰り）"}
	csv += strings.Join(header, ",")
	csv += "\n"

	aps, err := con.ListActivityPrograms(repository.ActivityProgramFilter{})
	if err != nil {
		return nil, err
	}

	var activeDatas []string
	for _, v := range aps {
		u := uMap[v.UserID]
		activeDatas := append(activeDatas, u.Name, fmt.Sprintf("%s / %s~%s", v.Datetime, v.StartTime, v.EndTime), "", v.OutwardTrip, v.ReturnTrip)
		csv += strings.Join(activeDatas, ",")
		csv += "\n"
	}

	return []byte(csv), nil
}
