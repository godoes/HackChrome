package utils

import (
	"crypto/aes"
	"crypto/cipher"
)

func AesGCMDecrypt(encrypted, key, nounce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode, _ := cipher.NewGCM(block)
	origData, err := blockMode.Open(nil, nounce, encrypted, nil)
	if err != nil {
		return nil, err
	}
	return origData, nil
}
