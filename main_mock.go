package main

import (
	"github.com/stretchr/testify/mock"
)

type svcMock struct {
	mock.Mock
}

func (m *svcMock) ParsePlatRawToString(plat string) int {
	args := m.Called(plat)
	return args.Int(0)
}

func (m *svcMock) PlatTypeCheck(numb int) string {
	args := m.Called(numb)
	return args.String(0)
}
