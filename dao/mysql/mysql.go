package mysql

import (
	"fmt"
	"go-web-app/settings"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
	//	viper.GetString("mysql.user"),
	//	viper.GetString("mysql.password"),
	//	viper.GetString("mysql.host"),
	//	viper.GetInt("mysql.port"),
	//	viper.GetString("mysql.dbname"),
	//)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}
