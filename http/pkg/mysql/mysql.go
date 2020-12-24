package mysql

// 参考 http://go-database-sql.org/connection-pool.html
import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	Master DbConfig
	Slave  DbConfig
}

type DbConfig struct {
	Host         string
	User         string
	Password     string
	Database     string
	MaxopenConns int
	MaxidleConns int
}

//主库连接池
var MysqlMasterPool *gorm.DB

//从库连接池
var MysqlSlavePool *gorm.DB

//初始化mysql 入口
func NewMySqlPool(c *Config) error {
	var err error
	MysqlMasterPool, err = getMysqlPool(c.Master)
	if err != nil {
		return err
	}
	MysqlSlavePool, err = getMysqlPool(c.Slave)
	if err != nil {
		return err
	}

	return nil
}

func getMysqlPool(dbConfig DbConfig) (*gorm.DB, error) {
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Database)

	log.Println("连接mysql: ", dns)

	pool, err := gorm.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	//在进程退出时释放mysql连接池  在入口处，应用程序退出时
	//defer pool.DB().Close()

	// mysql Pool 根据系统状况自行配置
	//pool.DB().SetConnMaxLifetime(time.Duration(mysqlSection.Key("maxlifetime").MustInt(30)) * time.Second)

	//pool.DB().SetMaxIdleConns(mysqlSection.Key("maxidleConns").MustInt(10))

	//pool.SetMaxOpenConns(n)  //不限制数据库最大并发数

	pool.DB().SetConnMaxLifetime(time.Second * 60 * 60) //运行1小时

	pool.DB().SetMaxIdleConns(dbConfig.MaxidleConns)

	pool.DB().SetMaxOpenConns(dbConfig.MaxopenConns)

	err = pool.DB().Ping()
	if err != nil {
		return nil, err
	}

	return pool, nil
}

/**
 * 释放连接池
 */
func CloseMysqlPool() error {

	var err error
	if MysqlMasterPool.DB() != nil {
		err = MysqlMasterPool.DB().Close()
		if err != nil {
			return err
		}
	}
	if MysqlSlavePool.DB() != nil {
		err = MysqlSlavePool.DB().Close()
		if err != nil {
			return err
		}
	}
	return nil
}
