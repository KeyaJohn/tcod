package database

import (
	"bytes"
	"fmt"
	"github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"strconv"
	"tcod/tools/config"
)

var (
	DbType   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
)

var Eloquent *gorm.DB
var MysqlConn string

type Mysql struct {
}

func (e *Mysql) Setup() {
	var err error
	var db Database
	db = new(Mysql)
	MysqlConn = db.GetConnect()

	Eloquent, err = db.Open(DbType, MysqlConn)
	if err != nil {
		panic(fmt.Errorf("%s connect error %v", DbType, err))
	} else {
		seelog.Info(fmt.Errorf("%s connect success!", DbType))
	}

	if Eloquent.Error != nil {
		panic(fmt.Errorf("database error %v", Eloquent.Error))
	}

	Eloquent.LogMode(true)
}

func (e *Mysql) Open(dbType string, conn string) (db *gorm.DB, err error) {
	return gorm.Open(dbType, conn)
}

func (e *Mysql) GetConnect() string {

	DbType = config.DatabaseConfig.Dbtype
	Host = config.DatabaseConfig.Host
	Port = config.DatabaseConfig.Port
	Name = config.DatabaseConfig.Name
	Username = config.DatabaseConfig.Username
	Password = config.DatabaseConfig.Password

	var conn bytes.Buffer
	conn.WriteString(Username)
	conn.WriteString(":")
	conn.WriteString(Password)
	conn.WriteString("@tcp(")
	conn.WriteString(Host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(Port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(Name)
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=1000ms")
	return conn.String()
}
