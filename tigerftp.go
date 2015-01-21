package main

import (
	"code.google.com/p/ftp4go"
	"fmt"
	"os"
)

const FTP_URL = "ftp2.census.gov"

func downloadAddrFeat(fips string) {

	ftpClient := ftp4go.NewFTP(0)

	_, err := ftpClient.Connect(FTP_URL, ftp4go.DefaultFtpPort, "")
	if err != nil {
		fmt.Println("ERROR: Could not connect to FTP server")
		os.Exit(-1)
	}

	defer ftpClient.Quit()

	_, loginerr := ftpClient.Login("anonymous", "", "")
	if loginerr != nil {
		fmt.Println("ERROR: Could not connect to FTP server")
		os.Exit(-1)
	}

	_, cwderror := ftpClient.Cwd("/geo/tiger/TIGER2014/ADDRFEAT")
	if cwderror != nil {
		fmt.Println("ERROR: Could not change to directory")
		os.Exit(-1)
	}

	zipList, listError := ftpClient.Nlst()
	if listError != nil {
		fmt.Println("ERROR: Could not list directory contents", listError)
		os.Exit(-1)
	}

	for _, zipElem := range zipList {
		if fips == zipElem[8:10] {
			fmt.Println("Downloading", zipElem)
			ftpClient.DownloadFile(zipElem, zipElem, false)
		}
	}

}
