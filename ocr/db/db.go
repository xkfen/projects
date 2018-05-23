package main

import (
	"fmt"
	"gcoresys/common/mysql"
	"gpreview/db/config"
	"flag"
	"gcoresys/common"
	"gpreview/model"
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
	dbConfig := config.GetPreviewDbConfig(env)
	mysql.CreateDB(dbConfig)
}

func doDrop(env string) {
	fmt.Println("do drop")
	dbConfig := config.GetPreviewDbConfig(env)
	mysql.DropDB(dbConfig)
}

func doMigrate(env string) {
	fmt.Println("do migrate")
	config.GetPreviewDbConfig(env)
	db := config.GetDb()
	db.AutoMigrate(
					&model.PreChannelApprovalOrder{},
	               )
}
