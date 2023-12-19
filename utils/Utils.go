package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

// PKCS7Padding fills plaintext as an integral multiple of the block length
func PKCS7Padding(p []byte, blockSize int) []byte {
	pad := blockSize - len(p)%blockSize
	padText := bytes.Repeat([]byte{byte(pad)}, pad)
	return append(p, padText...)
}

func AESCBCEncrypt(p []byte) (string, error) {
	key := []byte("lWTF^_FQS9PKMX!LrUfKkj5WkUUv9Sxs")
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	p = PKCS7Padding(p, block.BlockSize())
	ciphertext := make([]byte, len(p))
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	blockMode.CryptBlocks(ciphertext, p)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func GetFileMd5Values(filePath string) (string, error) {
	// 读取文件内容
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// 计算MD5哈希值
	hash := md5.Sum(data)
	// 将哈希值转换为十六进制字符串
	md5sum := hex.EncodeToString(hash[:])
	return md5sum, nil
}
