package config

import "testing"

func TestClearAllData(t *testing.T) {
	GetApprovalDbConfig("test")
	GetDb()
	ClearAllData()
}
