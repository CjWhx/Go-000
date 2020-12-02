package model

import (
	"database/sql"
	"fmt"
	"geek.com/lesson3-4/homework/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// 用户对象表
type User struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// 获取数据库连接
func getDB() *sql.DB {
	fmt.Println("连接数据库")
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.USER_NAME, conf.USER_PASSWORD, conf.DB_HOST, conf.DB_PORT, conf.DB_NAME)
	db, err := sql.Open("mysql", dsn)
	fmt.Println("init db:", db)
	if err != nil {
		panic(err)
	}

	// ping 确保联通
	dbConnectErr := db.Ping()
	if dbConnectErr != nil {
		panic("An error occurred when ping")
	}

	return db
}

// 查询数据
func QueryUser() ([]User, error) {

	// 获取数据连接，并保证数据库关闭
	db := getDB()
	defer db.Close()

	persons := make([]User, 0)

	// sql语句
	rows, err := db.Query("SELECT id,name,age FROM users;")
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while querying")
	}

	for rows.Next() {
		var id int64
		var name string
		var age int
		rows.Scan(&id, &name, &age)
		persons = append(persons, User{id, name, age})
	}

	if len(persons) == 0 {
		not_found_error := errors.New("not foud user data")
		return nil, errors.Wrap(not_found_error, "not found data")
	}

	return persons, nil

}
