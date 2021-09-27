package mfile

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

type FileContent struct {
	content  string
	FileInfo fs.FileInfo
}

func (fc *FileContent) Content() string {

	return fc.content
}

func (fc *FileContent) LastModified() string {

	filestat := fc.FileInfo
	timee := filestat.ModTime()
	dateTime := timee.Format("2006-01-02 15:04:05")
	return dateTime
}

func (fc *FileContent) ModeN() string {

	filestat := fc.FileInfo
	modeNumeric := fmt.Sprintf("%04o", filestat.Mode().Perm())
	return modeNumeric
}

func (fc *FileContent) ModeT() string {

	filestat := fc.FileInfo
	return filestat.Mode().String()
}

func GetFileContent(filePath string) (*FileContent, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	return &FileContent{content: string(data), FileInfo: fileInfo}, nil
}
