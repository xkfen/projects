package config

import (
	"gcoresys/common/mysql"
	"gcoresys/common/mysql/connection"
	"github.com/jinzhu/gorm"
	"fmt"
)

// 记录当前运行时的数据库配置
var curDbConfig *mysql.DbConfig

// 获取mysql数据库配置，此处还可以加到配置文件去获取配置的逻辑
func GetApprovalDbConfig(env string) *mysql.DbConfig {
	curDbConfig = mysql.NewDbConfig()
	switch env {
	case "prod":
		curDbConfig.DbName = "qy_approval_prod"
	case "test":
		curDbConfig.DbName = "qy_approval_test"
	default:
		curDbConfig.DbName = "qy_approval_dev"
	}
	return curDbConfig
}

// 获取当前服务的数据库连接
func GetDb() *gorm.DB {
	if curDbConfig == nil {
		panic("请先初始化数据库配置，调用：GetAccountingDbConfig方法")
	}
	return connection.GetDb(curDbConfig)
	//tmp := GetProdDb()
	//tmp.LogMode(true)
	//return  tmp
}

// 清空当前数据库的所有数据
func ClearAllData() {
	connection.ClearAllData(curDbConfig)
}






// ====================== 数据割接用 ===============================


var targetDb *gorm.DB
func GetTargetDb() *gorm.DB {
	if targetDb == nil {
		tmpDb, err  := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "qy", "hello_qy", "172.16.0.74", "3306", "qy_approval_prod"))
		if err != nil {
			panic(err.Error())
		}
		targetDb = tmpDb
	}
	return targetDb
}





var prodDb *gorm.DB

func GetProdDb() *gorm.DB {
	if prodDb == nil {
		tmpDb, err  := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "qy", "hello_qy", "172.16.1.90", "3306", "qy_approval_prod"))
		if err != nil {
			panic(err.Error())
		}
		prodDb = tmpDb
	}
	return prodDb
}

var rstDb *gorm.DB

func GetRstDb() *gorm.DB {
	if rstDb == nil {
		tmpDb, err  := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "qyread", "qyZonbXAs1mYUf", "rm-bp17gz8kz0z46y46ao.mysql.rds.aliyuncs.com", "3306", "qiyuan"))
		if err != nil {
			panic(err.Error())
		}
		rstDb = tmpDb
	}
	return rstDb
}



var test *gorm.DB

func GetTestDb() *gorm.DB {
	if test == nil {
		tmpDb, err  := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "qy", "hello_qy", "172.16.0.102", "3306", "qy_approval_prod"))
		if err != nil {
			panic(err.Error())
		}
		test = tmpDb
	}
	test.LogMode(true)
	return test
}