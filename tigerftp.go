package main

import (
	"code.google.com/p/ftp4go"
	"fmt"
	"os"
)

const FTP_URL = "ftp2.census.gov"

func downloadAddrFeat(fips string) {
	fmt.Println("Downloading state: " + fips)

	ftpClient := ftp4go.NewFTP(0)

	_, err := ftpClient.Connect(FTP_URL, ftp4go.DefaultFtpPort, "")
	if err != nil {
		fmt.Println("ERROR: Could not connect to FTP server")
		os.Exit(-1)
	}

	//defer ftpClient.Quit()

	_, loginerr := ftpClient.Login("anonymous", "", "")
	if loginerr != nil {
		fmt.Println("ERROR: Could not connect to FTP server")
		os.Exit(-1)
	}

	pwd, err := ftpClient.Pwd()
	if err != nil {
		fmt.Println("The Pwd command failed")
		os.Exit(-1)
	}

	fmt.Println("The current folder is", pwd)

	ftpClient.Quit()
}
