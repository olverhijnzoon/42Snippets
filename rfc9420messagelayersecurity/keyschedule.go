package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
)

type KeySchedule struct {
	EpochSecret []byte
}

func NewKeySchedule(epochSecret []byte) *KeySchedule {
	return &KeySchedule{
		EpochSecret: epochSecret,
	}
}

func (ks *KeySchedule) DeriveSecret(label string) ([]byte, error) {
	h := hmac.New(sha256.New, ks.EpochSecret)
	_, err := h.Write([]byte(label))
	if err != nil {
		return nil, errors.New("failed to derive secret")
	}
	return h.Sum(nil), nil
}

func (ks *KeySchedule) UpdateEpochSecret() error {
	newEpochSecret, err := ks.DeriveSecret("MLS 1.0 Epoch Secret")
	if err != nil {
		return errors.New("failed to update epoch secret")
	}
	ks.EpochSecret = newEpochSecret
	return nil
}
