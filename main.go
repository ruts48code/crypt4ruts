package crypt4ruts

import (
	"crypto/aes"
	"encoding/base64"
)

func MooEncrypt(data, key []byte) []byte {
	result := make([]byte, len(data))
	dummy := byte(0)
	dummy2 := byte(0)
	dummy3 := byte(0)
	dummy4 := byte(0)
	for i := range data {
		result[i] = data[i]
		result[i] = encode8(result[i], dummy)
		result[i] = encode8(result[i], dummy2)
		result[i] = encode8(result[i], dummy3)
		result[i] = encode8(result[i], dummy4)
		for j := 0; j < len(key); j++ {
			k := key[j]
			result[i] = encode8(result[i], k)
		}
		dummy4 = dummy3
		dummy3 = dummy2
		dummy2 = dummy
		dummy = result[i]
	}
	return result
}

func MooEncryptBase64(data []byte, key []byte) string {
	return base64.StdEncoding.EncodeToString(MooEncrypt(data, key))
}

func MooDecrypt(data, key []byte) []byte {
	result := make([]byte, len(data))
	dummy := byte(0)
	dummy2 := byte(0)
	dummy3 := byte(0)
	dummy4 := byte(0)
	ndummy := byte(0)
	for i := range data {
		result[i] = data[i]
		ndummy = data[i]
		for j := len(key) - 1; j >= 0; j-- {
			k := key[j]
			result[i] = decode8(result[i], k)
		}
		result[i] = decode8(result[i], dummy4)
		result[i] = decode8(result[i], dummy3)
		result[i] = decode8(result[i], dummy2)
		result[i] = decode8(result[i], dummy)
		dummy4 = dummy3
		dummy3 = dummy2
		dummy2 = dummy
		dummy = ndummy
	}
	return result
}

func MooHash(data, key, salt []byte) string {
	encode1 := MooEncrypt(data, key)
	encode2 := MooEncrypt(data, salt)
	encode3 := MooEncrypt(encode1, encode2)
	return string(salt) + base64.StdEncoding.EncodeToString(encode3)
}

func AESEncrypt(key []byte, data []byte) []byte {
	c, err := aes.NewCipher(key)
	if err != nil {
		return make([]byte, 0)
	}
	out := make([]byte, len(data))
	c.Encrypt(out, data)
	return out
}

func AESDecrypt(key []byte, data []byte) []byte {
	c, err := aes.NewCipher(key)
	if err != nil {
		return make([]byte, 0)
	}
	out := make([]byte, len(data))
	c.Decrypt(out, data)
	return out
}
