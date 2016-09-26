package main

import (
	"encoding/base64"
	"fmt"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func main() {

	// encode
	hello := "fpdm=3100144130&fphm=10049402&kpje=94017.09&kprq=2015-02-09"

	x1, err := AesEncrypt([]byte(hello))

	debyte := base64Encode(x1)
	fmt.Println(string(debyte))
	// decode
	enbyte, err := base64Decode([]byte(debyte))
	if err != nil {
		fmt.Println(err.Error())
	}

	if hello != string(enbyte) {
		fmt.Println("hello is not equal to enbyte")
	}

	fmt.Println(string(enbyte))

	// x2, err := AesEncrypt(enbyte)

	// fmt.Println(string(x2))
}

const (
	aesTable = "ywlSRb80TaCQ4b7b"
)

var (
	aesBlock       cipher.Block
	ErrAESTextSize = errors.New("ciphertext is not a multiple of the block size")
	ErrAESPadding  = errors.New("cipher padding size error")
)

func init() {

	var err error
	aesBlock, err = aes.NewCipher([]byte(aesTable))

	if err != nil {
		panic(err)
	}
}

// AES解密
func AesDecrypt(src []byte) ([]byte, error) {

	// 长度不能小于aes.Blocksize
	if len(src) < aes.BlockSize*2 || len(src)%aes.BlockSize != 0 {
		return nil, ErrAESTextSize
	}

	srcLen := len(src) - aes.BlockSize
	decryptText := make([]byte, srcLen)
	iv := src[srcLen:]
	mode := cipher.NewCBCDecrypter(aesBlock, iv)
	mode.CryptBlocks(decryptText, src[:srcLen])
	paddingLen := int(decryptText[srcLen-1])

	if paddingLen > 16 {
		return nil, ErrAESPadding
	}

	return decryptText[:srcLen-paddingLen], nil
}

// AES加密
func AesEncrypt(src []byte) ([]byte, error) {

	padLen := aes.BlockSize - (len(src) % aes.BlockSize)
	for i := 0; i < padLen; i++ {
		src = append(src, byte(padLen))
	}

	srcLen := len(src)
	encryptText := make([]byte, srcLen+aes.BlockSize)
	iv := encryptText[srcLen:]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(aesBlock, iv)
	mode.CryptBlocks(encryptText[:srcLen], src)
	return encryptText, nil

}
