package config

import "testing"

func TestClearAllData(t *testing.T) {
	GetOcrDbConfig("test")
	GetDb()
	ClearAllData()
}
