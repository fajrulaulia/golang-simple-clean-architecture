package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLalinGenapPlatMobilGenap(t *testing.T) {
	mockSvc := new(svcMock)

	mockSvc.On("ParsePlatRawToString", "BL2LK").Return(1)
	mockSvc.On("PlatTypeCheck", 1).Return("GENAP")

	svc := MainService{mockSvc}
	result := svc.Validate("BL2LK")

	t.Log("ceritanya Mobil ini genap, jadi aman saja, sistem validasi, tinggal pulici ngasih izin atau tidak")
	assert.Equal(t, result, "Plat anda Genap")

	mockSvc.AssertExpectations(t)
}

func TestLalinGenapPlatMobilMelanggar(t *testing.T) {
	mockSvc := new(svcMock)

	mockSvc.On("ParsePlatRawToString", "B1LK").Return(1)
	mockSvc.On("PlatTypeCheck", 1).Return("GANJIL")

	svc := MainService{mockSvc}
	result := svc.Validate("B1LK")

	t.Log("ceritanya oknum mobil plat B banyak tingkah")
	assert.NotEqual(t, result, "Plat anda Genap")

	mockSvc.AssertExpectations(t)
}
