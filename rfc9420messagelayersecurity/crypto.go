package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"errors"
)

func EncryptMessage(message *Message, publicKey *rsa.PublicKey) ([]byte, error) {
	if message == nil || publicKey == nil {
		return nil, errors.New("invalid parameters for encryption")
	}
	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		[]byte(message.Content),
		nil,
	)
	if err != nil {
		return nil, errors.New("encryption failed")
	}
	return ciphertext, nil
}

func DecryptMessage(ciphertext []byte, privateKey *rsa.PrivateKey) (string, error) {
	if ciphertext == nil || privateKey == nil {
		return "", errors.New("invalid parameters for decryption")
	}
	plaintext, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		ciphertext,
		nil,
	)
	if err != nil {
		return "", errors.New("decryption failed")
	}
	return string(plaintext), nil
}
