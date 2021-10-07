package client

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/icehubin/futu-go/pb/common"

	"github.com/icehubin/futu-go/logger"

	"github.com/icehubin/futu-go/adapt"
)

func (c *Client) bodyEncrypt(protoID uint32, body []byte) []byte {
	if !c.encrypt {
		return body
	}
	if protoID == adapt.ProtoID_InitConnect {
		return c.rsaEncrypt(body)
	} else {
		return c.aesEncrypt(body)
	}
}

func (c *Client) bodyDecrypt(protoID uint32, body []byte) []byte {
	if !c.encrypt {
		return body
	}
	if protoID == adapt.ProtoID_InitConnect {
		return c.rsaDecrypt(body)
	} else {
		return c.aesDecrypt(body)
	}
}

//---------- Todo encode and decode

func (c *Client) rsaEncrypt(body []byte) []byte {
	return RsaEncrypt(body, c.rsa_key)
}

func (c *Client) rsaDecrypt(body []byte) []byte {
	return RsaDecrypt(body, c.rsa_key)
}

func (c *Client) aesEncrypt(body []byte) []byte {
	switch c.encAlgo {
	case int32(common.PacketEncAlgo_PacketEncAlgo_FTAES_ECB):
		//not work
		return AesFTECBEncrypt(body, []byte(c.server.ConnAESKey))
	case int32(common.PacketEncAlgo_PacketEncAlgo_AES_CBC):
		//worked
		return AesCBCEncrypt(body, []byte(c.server.ConnAESKey), []byte(c.server.AesCBCiv))
	}
	logger.WithFields(logger.Fields{
		"method": "Client.aesEncrypt",
	}).Warning("encrypt failed reson of encAlgo not support")
	return body
}

func (c *Client) aesDecrypt(body []byte) []byte {
	switch c.encAlgo {
	case int32(common.PacketEncAlgo_PacketEncAlgo_FTAES_ECB):
		//not work
		return AesFTECBDecrypt(body, []byte(c.server.ConnAESKey))
	case int32(common.PacketEncAlgo_PacketEncAlgo_AES_CBC):
		//worked
		return AesCBCDecrypt(body, []byte(c.server.ConnAESKey), []byte(c.server.AesCBCiv))
	}
	logger.WithFields(logger.Fields{
		"method": "Client.aesDecrypt",
	}).Warning("encrypt failed reson of encAlgo not support")
	return body
}

//aes cbc加密
func AesCBCEncrypt(body []byte, key []byte, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.WithFields(logger.Fields{
			"method": "AesCBCEncrypt",
		}).Warning("encrypt failed reson of key load failed")
		return body
	}
	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		logger.WithFields(logger.Fields{
			"method": "aesCBCEncrypt",
		}).Warning("decrypt failed reson of iv lenth error")
		return body
	}
	body = PKCS7Padding(body, blockSize)
	//采用AES加密方法中CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(body))
	//执行加密
	blocMode.CryptBlocks(crypted, body)
	return crypted
}

//aes cbc解密
func AesCBCDecrypt(encrypted_body []byte, key []byte, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.WithFields(logger.Fields{
			"method": "aesCBCDecrypt",
		}).Warning("decrypt failed reson of key load failed")
		return encrypted_body
	}
	//获取块大小
	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		logger.WithFields(logger.Fields{
			"method": "aesCBCDecrypt",
		}).Warning("decrypt failed reson of iv lenth error")
		return encrypted_body
	}
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, iv)
	body := make([]byte, len(encrypted_body))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(body, encrypted_body)
	//去除填充字符串
	body, err = PKCS7UnPadding(body)
	if err != nil {
		logger.WithFields(logger.Fields{
			"method": "aesCBCDecrypt",
		}).Warning("decrypt failed reson of PKCS7UnPadding failed")
		return encrypted_body
	}
	return body
}

//aes FT ECB 加密
func AesFTECBEncrypt(body []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.WithFields(logger.Fields{
			"method": "aesFTECBEncrypt",
		}).Warning("encrypt failed reson of key load failed")
		return body
	}
	blockSize := block.BlockSize()
	mod_len := len(body) % blockSize
	if mod_len != 0 {
		body = append(body, bytes.Repeat([]byte{byte(0)}, blockSize-mod_len)...)
	}
	crypted := make([]byte, 0)
	tmpData := make([]byte, blockSize)
	for index := 0; index < len(body); index += blockSize {
		block.Encrypt(tmpData, body[index:index+blockSize])
		crypted = append(crypted, tmpData...)
	}
	crypted = append(crypted, bytes.Repeat([]byte{byte(0)}, blockSize-1)...)
	crypted = append(crypted, byte(mod_len))

	return crypted
}

//aes FT ECB 解密
func AesFTECBDecrypt(encrypted_body []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.WithFields(logger.Fields{
			"method": "aesFTECBEncrypt",
		}).Warning("encrypt failed reson of key load failed")
		return encrypted_body
	}
	blockSize := block.BlockSize()
	mod_len := int(encrypted_body[len(encrypted_body)-1])
	cut_len := blockSize - mod_len
	encrypted_body = encrypted_body[:len(encrypted_body)-blockSize]
	body := make([]byte, 0)
	tmpData := make([]byte, blockSize)
	for index := 0; index < len(encrypted_body); index += blockSize {
		block.Decrypt(tmpData, encrypted_body[index:index+blockSize])
		body = append(body, tmpData...)
	}
	body = body[:len(body)-cut_len]
	return body
}

func RsaEncrypt(body []byte, key []byte) []byte {
	block, _ := pem.Decode(key)
	if nil == block {
		panic("private key error")
	}
	encrypt_body := make([]byte, 0)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.PublicKey

	//最大加密长度
	maxLen := privateKey.Size() - 11

	for len(body) > maxLen {
		buff, err := rsa.EncryptPKCS1v15(rand.Reader, &publicKey, body[:maxLen])
		if err != nil {
			panic(err)
		}
		body = body[maxLen:]
		encrypt_body = append(encrypt_body, buff...)
	}
	if len(body) > 0 {
		buff, err := rsa.EncryptPKCS1v15(rand.Reader, &publicKey, body)
		if err != nil {
			panic(err)
		}
		encrypt_body = append(encrypt_body, buff...)
	}

	if err != nil {
		panic(err)
	}
	return encrypt_body
}

func RsaDecrypt(encrypted_body []byte, key []byte) []byte {
	block, _ := pem.Decode(key)
	body := make([]byte, 0)
	if block != nil {
		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			panic(err)
		}
		size := privateKey.Size()
		if len(encrypted_body) < size || len(encrypted_body)%size != 0 {
			panic("err encrypted body")
		}
		for len(encrypted_body) >= size {
			buff, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encrypted_body[:size])
			if err != nil {
				panic(err)
			}
			body = append(body, buff...)
			encrypted_body = encrypted_body[size:]
		}
	}
	return body
}

//PKCS7 填充模式
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//填充的反向操作，删除填充字符串
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}
