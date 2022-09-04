package myFunctions

import (
	"database/sql"
	"fmt"
)

func ConnectMySQL(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn) //不检验密码，用户名是否正确
	if err != nil {
		fmt.Printf("dsn 连接出错(格式有错)(用户名:密码@tcp(127.0.0.1:3306)/gotest)：%v , eer is %v\n", dsn, err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Ping err = %v\n", err)
		return nil
	}
	fmt.Printf("数据库连接成功\n")
	return db
}
