package localstorage

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	e "main/domain/errors"
	m "main/domain/model"
	rep "main/repository"

	"github.com/google/uuid"
)

type LocalStore struct {
	chatFilesPath     string
	homeworkFilesPath string
	solutionFilesPath string
	filestoragePath   string
}

func NewLocalStore(cfp string, hfp string, sfp string, fsp string) rep.FileStoreInterface {
	for _, path := range []string{cfp, hfp, sfp} {
		if err := os.MkdirAll(fsp+path, os.ModePerm); err != nil {
			log.Fatalln(e.StacktraceError(err))
		}
	}

	return &LocalStore{
		chatFilesPath:     cfp,
		homeworkFilesPath: hfp,
		solutionFilesPath: sfp,
		filestoragePath:   fsp,
	}
}

func (s *LocalStore) UploadFile(file *m.Attach) (string, error) {
	fileExt := ""
	switch file.MimeType {
	case "image/jpeg":
		fileExt = ".jpg"
	case "image/png":
		fileExt = ".png"
	case "image/svg+xml":
		fileExt = ".svg"
	case "application/pdf":
		fileExt = ".pdf"
	default:
		return "", e.StacktraceError(errors.New("error: " + file.MimeType + " is not allowed file extension"))
	}

	filePath := ""
	switch file.Dest {
	case "homework":
		filePath = s.filestoragePath + s.homeworkFilesPath
	case "solution":
		filePath = s.filestoragePath + s.solutionFilesPath
	case "chat":
		filePath = s.filestoragePath + s.chatFilesPath
	default:
		return "", e.StacktraceError(errors.New("error wrong destination"))
	}

	fileName := filePath + "/" + uuid.New().String() + fileExt

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", e.StacktraceError(err)
	}
	defer f.Close()

	resp, err := http.Get(file.FileURL)
	if err != nil {
		return "", e.StacktraceError(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", e.StacktraceError(err)
	}
	return fileName, nil
}
