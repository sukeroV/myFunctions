package myFunctions

import (
	"database/sql"
	"fmt"
	"reflect"
)

func SwitchAllTypr(iV interface{}) {
	// 这里的 v 是类型的值
	switch nameiV := iV.(type) {
	case bool:
		fmt.Printf("bool\n")
	case string:
		fmt.Printf("string\n")
	case int, int64:
		fmt.Printf("int or int 64\n")
	case float32, float64:
		fmt.Printf("float32 or 64\n")
	default:
		fmt.Printf("what Unknown is %v\n", reflect.TypeOf(nameiV))
	}
}

// 传入DSN
// 返回连接
func MySQL_Connect(dsn string) *sql.DB {
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
