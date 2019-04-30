package model

import (
    "os"
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

var Db *gorm.DB
var Musers map[uint]User

func init() {
    Connect(true)
    Musers = make(map[uint]User)
}

func Close() {
    Db.Close()
}

func Connect(real bool) {
    var err error

    if (real && Db != nil) {
        Close()
    }
    if (real || Db == nil) {
        Db, err = gorm.Open("mysql", os.Getenv("MYSQL_DSN_LIFE"))

        if err == nil { Db.SingularTable(true) } else { panic("[ERROR] [MySQL] connect failed, check and repair it") }
    }
}
