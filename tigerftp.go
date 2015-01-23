package main

import (
	"code.google.com/p/ftp4go"
	"log"
	"os"
)

const FTP_URL = "ftp2.census.gov"

func DownloadAddrFeat(fips string) []string {

	ftpClient := ftp4go.NewFTP(0)

	_, err := ftpClient.Connect(FTP_URL, ftp4go.DefaultFtpPort, "")
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	defer ftpClient.Quit()

	_, loginErr := ftpClient.Login("anonymous", "", "")
	if loginErr != nil {
		log.Fatal(loginErr)
		os.Exit(-1)
	}

	_, cwdError := ftpClient.Cwd("/geo/tiger/TIGER2014/ADDRFEAT")
	if cwdError != nil {
		log.Fatal(cwdError)
		os.Exit(-1)
	}

	zipList, listError := ftpClient.Nlst()
	if listError != nil {
		log.Fatal(listError)
		os.Exit(-1)
	}

	l := make([]string, len(zipList))
	i := 0

	for _, zipFile := range zipList {
		if fips == zipFile[8:10] {
			l[i] = zipFile
			log.Print("Downloading ", zipFile)
			ftpClient.DownloadFile(zipFile, zipFile, false)
			i += 1
		}
	}

	s := l[0:i]
	return s
}
