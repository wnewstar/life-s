package conf

type ConfMysql struct {
	Env 				string
	Dsn 				string
	MaxIdleConnNum		int
	MaxOpenConnNum		int
}

var MysqlWrite *ConfMysql = &ConfMysql{ Env: "deve", Dsn: "life:12345678@tcp(127.0.0.1:3306)/life?charset=utf8", MaxIdleConnNum: 3, MaxOpenConnNum: 10 }

var MysqlQuery *ConfMysql = &ConfMysql{ Env: "deve", Dsn: "life:12345678@tcp(127.0.0.1:3306)/life?charset=utf8", MaxIdleConnNum: 3, MaxOpenConnNum: 10 }
