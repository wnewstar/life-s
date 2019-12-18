package model

import (
    "os"
    "fmt"
    "life/conf"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Base struct {
    Id                          uint                        `json:"id" gorm:"AUTO_INCREMENT"`
    Status                      uint                        `json:"status"`
    UserDelete                  uint                        `json:"userd"`
    UserCreate                  uint                        `json:"userc"`
    UserUpdate                  uint                        `json:"useru"`
    TimeDelete                  uint64                      `json:"timed"`
    TimeCreate                  uint64                      `json:"timec"`
    TimeUpdate                  uint64                      `json:"timeu"`
}

var Tm uint

var DbQuery *gorm.DB
var DbWrite *gorm.DB

var Musers map[uint]User

func init() {
    initDbMysql()
    Musers = make(map[uint]User)
}

func initDbMysql() {
    var err error

    DbQuery, err = gorm.Open("mysql", conf.MysqlQuery.Dsn)

    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
    if conf.MysqlQuery.Env == "deve" {
        DbQuery.LogMode(true)
    }

    DbQuery.SingularTable(true)
    DbQuery.DB().SetMaxIdleConns(conf.MysqlQuery.MaxIdleConnNum)
    DbQuery.DB().SetMaxOpenConns(conf.MysqlQuery.MaxOpenConnNum)

    DbWrite, err = gorm.Open("mysql", conf.MysqlWrite.Dsn)

    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
    if conf.MysqlWrite.Env == "deve" {
        DbWrite.LogMode(true)
    }

    DbWrite.SingularTable(true)
    DbWrite.DB().SetMaxIdleConns(conf.MysqlWrite.MaxIdleConnNum)
    DbWrite.DB().SetMaxOpenConns(conf.MysqlWrite.MaxOpenConnNum)
}
