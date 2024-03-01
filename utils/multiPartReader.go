package utils

import (
	"bufio"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type MultiPartReader struct {
	reader *multipart.Reader
}

// @TODO: set application type check and field
type FilePart struct {
	Field string
}

func GetPartReader(req *http.Request) (*MultiPartReader, error) {
	reader, err := req.MultipartReader()
	if err != nil {
		return nil, err
	}

	partReader := &MultiPartReader{
		reader: reader,
	}

	return partReader, nil
}

func (r *MultiPartReader) NextTextPart(fieldName string) (string, error) {
	p, err := r.reader.NextPart()
	if err != nil {
		return "", err
	}

	if p.FormName() != fieldName {
		return "", errors.New("Expected Field Name " + fieldName)
	}

	text := make([]byte, 512)
	n, err := p.Read(text)
	if err != nil && err != io.EOF {
		return "", err
	}

	return string(text[:n]), nil
}

func (r *MultiPartReader) NextFilePart(filePart FilePart) (string, error) {
	p, err := r.reader.NextPart()
	if err != nil {
		return "", err
	}

	if p.FormName() != filePart.Field {
		return "", errors.New("Expected Field Name " + filePart.Field)
	}

	f, err := os.CreateTemp("./temp", "r*.m4a")
	if err != nil {
		return "", err
	}

	buf := bufio.NewReader(p)
	_, err = io.Copy(f, buf)
	if err != nil {
		return "", err
	}

	fileName := strings.Split(f.Name(), "\\")[1]

	f.Close()
	return fileName, nil
}
