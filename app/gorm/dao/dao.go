package dao

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sblog/system/global"
)

var masterDB *gorm.DB

func init() {
	global.Init()
	var err error

	//拼接dsn,基本都是固定死的 &timeout=30s
	//viper.GetString("database.timeout"),
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.dbname"),
		viper.GetString("database.charset"))

	//log.Println(dsn)

	masterDB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   viper.GetString("database.prefix"), // 表名前缀，`User` 对应的表名是 `tb_users`
			SingularTable: true,                               // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user` todo 一定要加这个要不然表后面都会加个s
		},
	})

	if err != nil {
		panic(fmt.Sprintf("gorm.Open err :%v,", err))
	}

}

// Close 关闭数据库
func Close() {
	dbSql, _ := masterDB.DB()
	_ = dbSql.Close()
}

//列表参数
type listParam struct {
	Page     int `json:"page" xml:"page" form:"page"`
	PageSize int `json:"page_size" xml:"page_size" form:"page_size"`
}
