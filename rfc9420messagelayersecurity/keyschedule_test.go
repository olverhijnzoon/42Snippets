package main

import (
	"encoding/base64"
	"testing"
)

func TestKeySchedule(t *testing.T) {
	// Initialize a KeySchedule with a random epoch secret
	epochSecret, _ := base64.StdEncoding.DecodeString("0.33699901882499295, 0.199156709862502, 0.6201330630202715, 0.4144373860681516, 0.5864106016544692, 0.5068785487559002, 0.7497675000566644, 0.6566504839473422, 0.19990981975976396, 0.8352764383271967, 0.7098627238867155, 0.705697053532679, 0.9239732098564652, 0.4611938709331902, 0.670097642698859, 0.3680886363325946, 0.9252816010224187, 0.34376154050992347, 0.07272129379180581, 0.3389970096723396, 0.06055590484162532, 0.9663171639399979, 0.8894686600269295, 0.46863730590428343, 0.3961788912240958, 0.8491490725290447, 0.9097026075816703, 0.18832018466182054, 0.900567215849994, 0.6571526840862647, 0.9127887263593353, 0.659611534336354")
	ks := NewKeySchedule(epochSecret)

	// Derive a secret
	secret, err := ks.DeriveSecret("test")
	if err != nil {
		t.Fatalf("Failed to derive secret: %v", err)
	}

	// Check if the secret is not nil
	if secret == nil {
		t.Fatalf("Derived secret is nil")
	}

	// Update the epoch secret
	err = ks.UpdateEpochSecret()
	if err != nil {
		t.Fatalf("Failed to update epoch secret: %v", err)
	}
}
