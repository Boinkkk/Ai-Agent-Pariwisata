package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Parameter Argon

const (
	time    = 1
	memory  = 64 * 1024
	threads = 4
	keyLen  = 32
)

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)

	_, err := rand.Read(salt)

	return salt, err
}

func HashPassword(password string) (string, error) {
	salt, err := generateSalt()

	if err != nil {
		fmt.Println("Something Error with Generating Salt...")
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return b64Salt + "." + b64Hash, nil
}

func VerifyPassword(password, encodeHash string) bool {
	fmt.Println("Input:", password)
	fmt.Println("Length:", len(password))
	fmt.Println("Stored hash:", encodeHash)
	parts := strings.Split(encodeHash, ".")

	if len(parts) != 2 {
		fmt.Println("Something Wrong With hash")
		return false
	}

	salt, _ := base64.RawStdEncoding.DecodeString(parts[0])
	hash, _ := base64.RawStdEncoding.DecodeString(parts[1])

	fmt.Printf("Salt: %x\n", salt)

	newHash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	fmt.Printf("Stored: %x\n", hash)
	fmt.Printf("New   : %x\n", newHash)

	fmt.Println(hash, newHash)

	return subtle.ConstantTimeCompare(hash, newHash) == 1
}
