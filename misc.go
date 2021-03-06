package otp

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
)

func GenerateToken(size int) string {
	return randomString(size)
}

func randomString(size int) string {
	var bytes = make([]byte, size)
	rand.Read(bytes)
	
	return base32.StdEncoding.EncodeToString(bytes)
}

func counterToBytes(counter uint64) (text []byte) {
	text = make([]byte, 8)
	for i := 7; i >= 0; i-- {
		text[i] = byte(counter & 0xff)
		counter = counter >> 8
	}
	return
}

func hmacSHA1(key, text []byte) []byte {
	H := hmac.New(sha1.New, key)
	H.Write([]byte(text))
	return H.Sum(nil)
}

func truncate(hash []byte) int {
	offset := int(hash[len(hash)-1] & 0xf)
	return ((int(hash[offset]) & 0x7f) << 24) |
		((int(hash[offset+1] & 0xff)) << 16) |
		((int(hash[offset+2] & 0xff)) << 8) |
		(int(hash[offset+3]) & 0xff)
}
