package orm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/os/gfile"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path"
)

// 在其它model的实体类中可直接调用
var DB *gorm.DB

func InitDB(conf *ini.File) *gorm.DB {
	var err error

	// 数据库的类型
	dbType := conf.Section("").Key("db_type").String()

	// Mysql配置信息
	mysqlName := conf.Section("mysql").Key("db_name").String()
	mysqlUser := conf.Section("mysql").Key("db_user").String()
	mysqlPwd := conf.Section("mysql").Key("db_pwd").String()
	mysqlHost := conf.Section("mysql").Key("db_host").String()
	mysqlPort := conf.Section("mysql").Key("db_port").String()
	mysqlCharset := conf.Section("mysql").Key("db_charset").String()

	// sqlite3配置信息
	sqliteName := conf.Section("sqlite3").Key("db_name").String()

	var dataSource string
	switch dbType {
	case "mysql":
		dataSource = mysqlUser + ":" + mysqlPwd + "@tcp(" + mysqlHost + ":" +
			mysqlPort + ")/" + mysqlName + "?charset=" + mysqlCharset

		DB, err = gorm.Open(dbType, dataSource)
	case "sqlite3":
		dataSource = "database" + string(os.PathSeparator) + sqliteName
		if !gfile.Exists(dataSource) {
			os.MkdirAll(path.Dir(dataSource), os.ModePerm)
			os.Create(dataSource)
		}
		DB, err = gorm.Open(dbType, dataSource)
	}

	if err != nil {
		DB.Close()
		log.Fatal("数据库连接异常：", err)
	}

	// 设置连接池，空闲连接
	DB.DB().SetMaxIdleConns(50)
	// 打开链接
	DB.DB().SetMaxOpenConns(100)

	// 表明禁用后缀加s
	DB.SingularTable(true)

	return DB
}
