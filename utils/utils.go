package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

//goland:noinspection GoUnusedExportedFunction
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func CopyFile(source, dest string) bool {
	if source == "" || dest == "" {
		log.Println("source or dest is null")
		return false
	}

	sourceOpen, err := os.Open(source)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer func(sourceOpen *os.File) {
		_ = sourceOpen.Close()
	}(sourceOpen)

	destOpen, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer func(destOpen *os.File) {
		_ = destOpen.Close()
	}(destOpen)

	_, copyErr := io.Copy(destOpen, sourceOpen)
	if copyErr != nil {
		log.Println(copyErr.Error())
		return false
	} else {
		return true
	}
}

func Merge(res1, res2 map[string]map[string]string) map[string]map[string]string {
	for k, v := range res2 {
		if _, ok := res1[k]; ok {
			if len(v["password"]) > 0 && len(res1[k]["password"]) == 0 {
				res1[k]["password"] = v["password"]
			}
		} else {
			res1[k] = v
		}
	}

	return res1
}

func FormatOutput(totalRes map[string]map[string]string, pwdDb string) error {
	for k, v := range totalRes {
		fmt.Printf("====================\n")
		fmt.Printf("Url: %s\nUsername: %s\nPassword:%s\n\n", k, v["username"], v["password"])
	}

	fmt.Printf("\nTotal Auth: %d", len(totalRes))
	err := os.Remove(pwdDb)
	if err != nil {
		return err
	}
	return nil
}
