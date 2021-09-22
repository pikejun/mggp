package main

import (
	"fmt"
	"github.com/pikejun/mggp/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pikejun/mggp/global"
	"github.com/pikejun/mggp/pkg/setting"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Printf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Printf("init.setupDBEngine err: %v", err)
	}
}

func main() {
	// log.Printf("ServerSetting:%#v", global.ServerSetting)
	// log.Printf("DatabaseSetting:%#v", global.DatabaseSetting)
	gin.SetMode(global.ServerSetting.RunMode) // gin 的运行模式
	router := config.NewRouter()

	// 自定义 http.Server
	s := &http.Server{
		Addr:         ":" + global.ServerSetting.HttpPort, // 监听端口
		Handler:      router,                              // 处理程序
		ReadTimeout:  global.ServerSetting.ReadTimeout,    // 允许读取的最大时间
		WriteTimeout: global.ServerSetting.WriteTimeout,   // 允许写入的最大时间
		// MaxHeaderBytes: 1 << 20,          // 请求头的最大字节数
		MaxHeaderBytes: 1024 * 1024, // 请求头的最大字节数
	}

	s.ListenAndServe()

	fmt.Println("Server listen on ",global.ServerSetting.HttpPort)
}

// 读取配置
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

// 初始化 DB
func setupDBEngine() error {
	var err error
	global.DBEngine, err = NewDBEngine(global.DatabaseSetting)
	global.DBEngine.Logger.LogMode(0)
	if err != nil {
		return err
	}
	return nil
}


// 创建 DB 实例
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", databaseSetting.Username, databaseSetting.Password, databaseSetting.Host, databaseSetting.DBName, databaseSetting.Charset, databaseSetting.ParseTime)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ Logger:logger.Default.LogMode(logger.Info),})
	db.Logger.LogMode(logger.Info)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(global.DatabaseSetting.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(global.DatabaseSetting.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, nil
}

