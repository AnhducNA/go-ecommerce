package initialize

import (
	"fmt"
	"time"

	"github.com/AnhducNA/go-ecommerce/global"
	"github.com/AnhducNA/go-ecommerce/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrPanic(err error, msg string) {
	if err != nil {
		global.Logger.Error(msg, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.MySQL
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s?charset=utf8mb4&parseTime=True&loc=Local"
	var str = fmt.Sprintf(dsn, m.Url)
	db, err := gorm.Open(mysql.Open(str), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrPanic(err, "InitMysql error")
	global.Logger.Info("InitMysql mysql connect success")
	global.Mdb = db

	// set Pool
	SetPool()
	migrateTables()

}

func SetPool() {
	m := global.Config.MySQL
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("Mysql connect error: %s::", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		fmt.Println("AutoMigrate error: ", err)
	}
}
