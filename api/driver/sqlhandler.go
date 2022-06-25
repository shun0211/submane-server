package driver

import (
	"api/adapter/database"
	"api/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlHandler struct {
	Conn *gorm.DB
}

// SqlHandlerの実装部分
// FUCK: database.SqlHandlerが戻り値で指定されているのに対して、関数内ではSqlHandlerのポインタ型を返していてOKなのかが分からない
func NewSqlHandler() database.SqlHandler {
	dsn := "host=postgres user=gorm password=gorm dbname=submane_db port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}) // &gormはポインタ型
	if err != nil {
		panic(err.Error())
	}

	// NOTE: Auto Migration
	// conn.Migrator().DropTable("subscriptions")
	conn.AutoMigrate(domain.User{}, domain.Subscription{})

	sqlHandler := &SqlHandler{Conn: conn}
	return sqlHandler
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Order("created_at desc").Find(out, where...)
}

func(handler *SqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

// NOTE: SELECT文以外はExec関数を使う
func(handler *SqlHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(sql, values...)
}

// NOTE: SELECT文はRaw関数を使う
func(handler *SqlHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Raw(sql, values...)
}

func(handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

func(handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

func(handler *SqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

func(handler *SqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}

func(handler *SqlHandler) Paginate(page domain.Page) *gorm.DB {
	offset := (page.Page - 1) * page.PerPage
	return handler.Conn.Offset(offset).Limit(page.PerPage)
}
