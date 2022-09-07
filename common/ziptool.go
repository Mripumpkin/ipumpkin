package common

import (
	"io"
	"os"
	"path"

	ezip "github.com/alexmullins/zip"
)

func Zipfile(filename []string, destfile string, passwd string) error {
	zipfile, err := os.Create(destfile)
	if err != nil {
		return err
	}
	defer zipfile.Close()
	zipWriter := ezip.NewWriter(zipfile)
	defer zipWriter.Close()
	// gzip.NewWriter(zipfile)
	for _, file := range filename {
		if err = AddfileToZip(zipWriter, file, passwd); err != nil {
			return err
		}
	}
	return nil
}

func AddfileToZip(zipfile *ezip.Writer, file string, passwd string) error {
	newfile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer newfile.Close()
	fs, err := newfile.Stat()
	if err != nil {
		return err
	}
	header, err := ezip.FileInfoHeader(fs)
	if err != nil {
		return err
	}
	header.Name = path.Base(file)
	header.SetPassword(passwd)
	zipWriter, err := zipfile.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(zipWriter, newfile)
	return err
}
