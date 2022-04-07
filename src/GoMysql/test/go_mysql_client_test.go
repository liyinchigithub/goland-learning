package test

import (
	"database/sql" // go数据库泛接口
	"log"          // 日志输出
	"testing"      // 单元测试

	_ "github.com/go-sql-driver/mysql" // 数据库驱动	因为没有直接使用该包中的对象，所以在导入包前面加上了下划线
	// "time"
)

/*
	1.下载驱动：go get -u github.com/go-sql-driver/mysql
	2.函数：
		（1）Open(指定数据库,指定数据源) 用于连接数据库，指定数据源"账号:密码@tcp(数据库ip和端口)/数据库名称?charset=utf8"
		（2）Ping() 检查数据源名称是否合法可用
		（3）SetMaxOpenConns(设置最大连接数)
		（4）SetMaxIdleConns(设置最大空闲连接数)
		（5）Query(查询语句)	多行查询
		（6）QueryRow(查询语句)	单行查询
		（7）Exec(执行语句)	执行语句  用于执行一次命令（插入、删除 、更新、查询）
		（8）Prepare(预编译语句)	预编译语句 会将sql语句发送给mysql服务器端，返回一个好准备好的状态，用于之后的查询和命令，返回值可以同时执行多个查询和命令
		（9）事务 Begin()  Commit()  Rollback()
*/

// var db *sql.DB // 创建数据(数据库操作句柄)

func initDB() (err error) {
	//创建数据对象 当前并未创建实际连接
	db, err := sql.Open("mysql", "root:lyc123456@tcp(127.0.0.1:3306)/test?charset=utf8") // 数据库配置
	// 判断是否连接成功
	if err != nil {
		// panic(err)
		return err
	}
	// 检查数据库是否可用
	err = db.Ping()
	// 判断是否连接成功
	if err != nil {
		// panic(err)
		return err
	} else {
		log.Println("Connected to MySQL successfully！")
	}
	return nil
}

func TestRun(t *testing.T) {
	initDB()
	// 单元测试方法执行顺序
	// t.Run("TestMySQLCreate", TestMySQLCreate)
	// t.Run("TestMysqlInsert", TestMysqlInsert)
	t.Run("TestMysqlSelectQuery", TestMysqlSelectQuery)
	t.Run("TestMysqlSelectQueryRow", TestMysqlSelectQueryRow)
	// t.Run("TestMysqlSelectPrepare", TestMysqlSelectPrepare)
	// t.Run("TestMysqlUpdate",TestMysqlUpdate)
}

func TestMySQLCreate(t *testing.T) {
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
		// panic(err)
	} else {
		log.Println("Connected to MySQL successfully！")
	}
	// 预处理
	stmt, err := db.Prepare("CREATE TABLE person2 (id int NOT NULL AUTO_INCREMENT, first_name varchar(40), last_name varchar(40) , birthday  varchar(40) ,PRIMARY KEY (id));")
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

// 多行查询
func TestMysqlSelectQuery(t *testing.T) {
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
		// panic(err)
	} else {
		log.Println("Connected to MySQL successfully！")
	}
	// Query 查询多行数据
	rows, err := db.Query("select * from person order by id") // 执行语句且无返回，调用完后会自动释放连接。
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
			// 调用Scan方法将行数据保存到变量中
			err = rows.Scan(&id, &first_name, &last_name, &birthday) // 每个遍历行都会调用Scan方法
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

// 多单行查询
func TestMysqlSelectQueryRow(t *testing.T) {
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
		// panic(err)
	} else {
		log.Println("Connected to MySQL successfully！")
	}
	type Person struct {
		id         int32
		first_name string
		last_name  string
		birthday   string
	}
	var p Person
	// Query 查询多行数据
	db.QueryRow("select * from `person` order by id").Scan(&p.id, &p.first_name, &p.last_name, &p.birthday) // 执行语句且无返回，调用完后会自动释放连接。
	// 判断是否执行成功
	if err != nil {
		log.Println("db query error , " + err.Error())
	}
	// 循环读取数据
	// 调用Scan方法将行数据保存到变量中
	if err != nil {
		log.Println("db scan error , " + err.Error())
	} else {
		log.Printf("id=%d,first_name=%s,last_name=%s,birthday=%s", p.id, p.first_name, p.last_name, p.birthday)
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
	stmt, err := db.Prepare("SELECT * FROM person;") // 执行语句且无返回，调用完后会自动释放连接。
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
	db, err := sql.Open("mysql", "root:lyc123456@tcp(127.0.0.1:3306)/test?charset=utf8") // 数据库配置
	// 判断是否连接成功
	if err != nil {
		log.Print(err.Error())
	}
	// 关闭数据库
	defer db.Close()

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
