package test

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	// "time"
)

func TestRun(t *testing.T) {
	//创建数据对象 当前并未创建实际连接
	db, err := sql.Open("mysql", "root:lyc123456@tcp(127.0.0.1:3306)/test?charset=utf8") // 数据库配置
	// 判断是否连接成功
	if err != nil {
		panic(err)
	}
	// 检查数据库是否可用
	err = db.Ping()
	// 判断是否连接成功
	if err != nil {
		panic(err)
	}
	// 执行 顺序
	// t.Run("TestMySQLCreate",TestMySQLCreate)
	// t.Run("TestMysqlInsert", TestMysqlInsert)
	// t.Run("TestMysqlSelectQuery", TestMysqlSelectQuery)
	t.Run("TestMysqlSelectPrepare", TestMysqlSelectPrepare)
	// t.Run("TestMysqlUpdate",TestMysqlUpdate)
}

func TestMySQLCreate(t *testing.T) {
	/*
		@params	string 驱动名称，为避免混淆推荐与包名相同
		@params string 数据源URL
	*/
	db, err := sql.Open("mysql", "root:lyc123456@tcp(127.0.0.1:3306)/test?charset=utf8") // 数据库配置
	// 判断是否连接成功
	if err != nil {
		log.Print(err.Error())
	}
	// 关闭数据库
	defer db.Close() //  延迟释放连接资源
	// 检查数据库是否可用
	err = db.Ping()
	// 判断是否连接成功
	if err != nil {
		log.Print(err.Error())
	}
	// 创建一个gin引擎
	stmt, err := db.Prepare("CREATE TABLE person (id int NOT NULL AUTO_INCREMENT, first_name varchar(40), last_name varchar(40) , birthday  varchar(40) , PRIMARY KEY (id));")
	//	判断是否创建成功
	if err != nil {
		log.Println(err.Error())
	}
	//  执行查询
	_, err = stmt.Exec()
	// 判断是否执行成功
	if err != nil {
		log.Print("发生错误：", err.Error())
	} else {
		log.Printf("Person Table successfully migrated....")
	}

}

/*
	查询数据
*/
func TestMysqlSelectQuery(t *testing.T) {
	db, err := sql.Open("mysql", "root:lyc123456@tcp(127.0.0.1:3306)/test?charset=utf8") // 数据库配置
	// 判断是否连接成功
	if err != nil {
		log.Print(err.Error())
	}
	// 关闭数据库
	defer db.Close()
	// query
	rows, err := db.Query("select * from person order by id")// 执行语句且无返回，调用完后会自动释放连接。
	// 判断是否执行成功
	if err != nil {
		log.Println("db query error , " + err.Error())
	} else {
		log.Println("db query success")
		// 循环遍历结果集
		var (
			id         int32
			first_name string
			last_name  string
			birthday   string
		)
		// 循环读取数据
		for rows.Next() {
			err = rows.Scan(&id, &first_name, &last_name, &birthday)
			if err != nil {
				log.Println("db scan error , " + err.Error())
			} else {
				log.Printf("id=%d,first_name=%s,last_name=%s,birthday=%s", id, first_name, last_name, birthday)
			}
		}
	}
	// 判断是否执行成功
	if err != nil {
		log.Print(err.Error())
	} else {
		log.Println("Person Table successfully migrated....")
	}
}
func TestMysqlSelectPrepare(t *testing.T) {
	db, err := sql.Open("mysql", "root:lyc123456@tcp(127.0.0.1:3306)/test?charset=utf8") // 数据库配置
	// 判断是否连接成功
	if err != nil {
		log.Print(err.Error())
	}
	// 关闭数据库
	defer db.Close()
	// Prepare
	stmt, err := db.Prepare("SELECT * FROM person;")// 执行语句且无返回，调用完后会自动释放连接。
	//	判断是否创建成功
	if err != nil {
		log.Println(err.Error())
	}
	//  执行查询
	_, err = stmt.Exec() // 执行语句且无返回，调用完后会自动释放连接。

	// 判断是否执行成功
	if err != nil {
		log.Print(err.Error())
	} else {
		log.Println("Person Table successfully migrated....")
		log.Println("stmt:", stmt) // stmt: &{0xc00007fad0 SELECT * FROM person; <nil> {{0 0} 0 0 0 0} <nil> 0xc00002e8c0 <nil> {0 0} false [{0xc0001321b0 0xc00002e8c0}] 0}
	}
}

/*
	修改数据
*/
func TestMysqlUpdate(t *testing.T) {
	// db, err := sql.Open("mysql", "root:lyc123456@tcp(127.0.0.1:3306)/test?charset=utf8")// 数据库配置
	// // 判断是否连接成功
	// if err != nil {
	//     fmt.Print(err.Error())
	// }
	// // 关闭数据库
	// defer db.Close()

}

/*
	插入数据
*/
func TestMysqlInsert(t *testing.T) {
	insert("INSERT INTO person (first_name, last_name, birthday) VALUES ('James', 'Bond', 'aaaa');")
}

//insert 插入数据表
func insert(query string) int64 {
	db, err := sql.Open("mysql", "root:lyc123456@tcp(127.0.0.1:3306)/test?charset=utf8") // 数据库配置
	// 判断是否连接成功
	if err != nil {
		log.Print(err.Error())
	}
	// 关闭数据库
	defer db.Close()

	//执行语句获取结果集
	result, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	// 获取插入的主键
	last_insert_id, err := result.LastInsertId()
	// 判断是否执行成功
	if err != nil {
		log.Fatal(err)
	}
	// 返回插入的主键
	log.Println(last_insert_id)
	// 返回影响的行数
	return last_insert_id
}

/*
	删除数据
*/
func TestMysqlDelete(t *testing.T) {
	db, err := sql.Open("mysql", "root:lyc123456@tcp(127.0.0.1:3306)/test?charset=utf8") // 数据库配置
	// 判断是否连接成功
	if err != nil {
		log.Print(err.Error())
	}
	// 关闭数据库
	defer db.Close()

}
