package main

import (
	"fmt"
	"gapproval/approval/db/config"
	"gcoresys/common/mysql"
	"flag"
	"gcoresys/common"
	"firstGoLang/model"
)

func main() {
	common.DefineDbMigrateCommonFlag()
	env := common.DefineRunTimeCommonFlag()
	action := flag.Lookup("action").Value.String()
	switch action {
	case "create":
		doCreate(env)
	case "drop":
		doDrop(env)
	case "migrate":
		doMigrate(env)
	}
}

func doCreate(env string) {
	fmt.Println("do create")
	dbConfig := config.GetApprovalDbConfig(env)
	mysql.CreateDB(dbConfig)
}

func doDrop(env string) {
	fmt.Println("do drop")
	dbConfig := config.GetApprovalDbConfig(env)
	mysql.DropDB(dbConfig)
}

func doMigrate(env string) {
	fmt.Println("do migrate")
	config.GetApprovalDbConfig(env)
	db := config.GetDb()
	db.AutoMigrate(&model.ApprovalKaty{})
}
