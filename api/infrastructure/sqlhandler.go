package infrastructure

import (
	"api/interfaces/controllers/database"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlHandler struct {
	Conn *gorm.DB
}

// SqlHandlerの実装部分
// FUCK: database.SqlHandlerが戻り値で指定されているのに対して、関数内ではSqlHandlerのポインタ型を返していてOKなのかが分からない
func NewSqlHandler() database.SqlHandler {
	dsn := "host=postgres user=root password=password dbname=submane_db port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // &gormはポインタ型
	if err != nil {
		panic(err.Error())
	}
	sqlHandler := &SqlHandler{Conn: conn}
	return sqlHandler
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
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
