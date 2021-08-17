package db

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"leeter/pkg/logging"
)

type Config struct {
	Alias    string
	Driver   string
	Host     string
	Port     string
	Name     string
	Charset  string
	User     string
	Password string
	Timezone string
}

const dbSource = "%s:%s@tcp(%s:%s)/%s?charset=%s"

var (
	defaultDBConfig Config
)

func Init() {
	setUp()
}

func setUp() {
	initConfig()
	dbConfig := getDBConfig()
	registerDbByConfig(dbConfig)
	registerOrmByConfig(dbConfig)
}

// initConfig 初始化配置
func initConfig() {
	defaultDBConfig.Alias, _ = web.AppConfig.String("db.alias")
	defaultDBConfig.Driver, _ = web.AppConfig.String("db.driver")
	defaultDBConfig.Host, _ = web.AppConfig.String("db.host")
	defaultDBConfig.Port, _ = web.AppConfig.String("db.port")
	defaultDBConfig.User, _ = web.AppConfig.String("db.user")
	defaultDBConfig.Password, _ = web.AppConfig.String("db.password")
	defaultDBConfig.Name, _ = web.AppConfig.String("db.name")
	defaultDBConfig.Charset, _ = web.AppConfig.String("db.charset")
	defaultDBConfig.Timezone, _ = web.AppConfig.String("db.timezone")
}

// 获取数据库配置
func getDBConfig() Config {
	return defaultDBConfig
}

// registerDb 注册数据库
func registerDbByConfig(dbConfig Config) {
	var err error
	strDBSource := fmt.Sprintf(dbSource, dbConfig.User, dbConfig.Password, dbConfig.Host,
		dbConfig.Port, dbConfig.Name, dbConfig.Charset)
	err = orm.RegisterDataBase(dbConfig.Alias, dbConfig.Driver, strDBSource)
	if err != nil {
		logging.ErrorConfig("registerDb", nil, nil, "register database fail", err)
	}

	if web.BConfig.RunMode == web.DEV {
		orm.Debug = true
	}
}

// registerOrmByConfig 注册 orm
func registerOrmByConfig(dbConfig Config) orm.Ormer {
	return orm.NewOrmUsingDB(dbConfig.Alias)
}

// GetORM 获取orm，beego debug模式返回新orm对象,生产模式返回缓存orm对象,所以应用层无需缓存orm，保存dbAlias即可
func GetORM() orm.Ormer {
	return getDefaultORM()
}

func getDefaultORM() orm.Ormer {
	return orm.NewOrmUsingDB(defaultDBConfig.Alias)
}
