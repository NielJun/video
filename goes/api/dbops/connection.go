package dbops

import (
  "database/sql"
  _"github.com/go-sql-driver/mysql"
)

var (
  dbConn *sql.DB
  err    error
)

//数据库连接初始化  会被自动回调
func init() {
  dbConn, err = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/video_server?charset=utf8&parseTime=true")
  if err != nil {
    panic(err.Error())
  }
  println("---------->")
}

// 数据库连接文件
