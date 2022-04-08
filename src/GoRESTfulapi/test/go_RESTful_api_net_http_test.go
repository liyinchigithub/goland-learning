package test

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

/*
	[官方文档]
	https://github.com/gin-gonic/gin
*/
func TestRESTfulAPI(t *testing.T) {
	// 数据库配置
	db, err := sql.Open("mysql", "root:passapp@tcp(127.0.0.1:3306)/test?charset=utf8") // 参数1：数据库驱动名称，参数2：数据库连接字符串(/数据库名称) 参数3：可选参数
	// 判断是否连接成功
	if err != nil {
		fmt.Print(err.Error())
	}
	// 关闭数据库
	defer db.Close()
	// make sure connection is available
	// 执行查询
	err = db.Ping()
	// 判断是否连接成功
	if err != nil {
		fmt.Print(err.Error())
	}
	// 创建一个gin实例
	type Person struct {
		Id         int
		First_Name string
		Last_Name  string
	}
	// 创建一个gin引擎
	router := gin.Default()
	// 添加一个路由，并且使用GET方法
	router.GET("/person/:id", func(c *gin.Context) {
		var (
			person Person
			result gin.H
		)
		// 获取参数
		id := c.Param("id")                                                                  // 获取URL参数
		row := db.QueryRow("select id, first_name, last_name from person where id = ?;", id) // 查询数据
		err = row.Scan(&person.Id, &person.First_Name, &person.Last_Name)                    // 查询数据
		// 判断是否查询成功
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": person,
				"count":  1,
			}
		}
		// 返回JSON
		c.JSON(http.StatusOK, result)
	})
	// GET all persons
	// 添加一个路由，并且使用GET方法
	router.GET("/persons", func(c *gin.Context) {
		// 创建一个数组
		var (
			person  Person
			persons []Person
		)
		// 执行查询
		rows, err := db.Query("select * from person;") // 查询数据
		// 判断是否查询成功
		if err != nil {
			fmt.Print(err.Error())
		}
		// 循环遍历查询结果
		for rows.Next() {
			err = rows.Scan(&person.Id, &person.First_Name, &person.Last_Name)
			persons = append(persons, person)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close() // 关闭数据库
		// 返回JSON
		c.JSON(http.StatusOK, gin.H{
			"result": persons,
			"count":  len(persons),
		})
	})
	// POST new person details
	// 添加一个路由，并且使用POST方法
	router.POST("/person", func(c *gin.Context) {
		// 获取表单数据
		var buffer bytes.Buffer                                                            // 创建一个buffer
		first_name := c.PostForm("first_name")                                             // 获取表单数据
		last_name := c.PostForm("last_name")                                               // 获取表单数据
		stmt, err := db.Prepare("insert into person (first_name, last_name) values(?,?);") // 准备执行语句
		// 判断是否准备成功
		if err != nil {
			fmt.Print(err.Error())
		}
		// 执行语句
		_, err = stmt.Exec(first_name, last_name)
		// 判断是否执行成功
		if err != nil {
			fmt.Print(err.Error())
		}
		// Fastest way to append strings
		// 将数据添加到buffer中
		buffer.WriteString(first_name)
		buffer.WriteString(" ")
		buffer.WriteString(last_name)
		defer stmt.Close()      //  关闭语句
		name := buffer.String() //  获取buffer中的数据
		// 返回JSON
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", name),
		})
	})
	// PUT - update a person details
	// 添加一个路由，并且使用PUT方法
	router.PUT("/person", func(c *gin.Context) {
		// 获取表单数据
		var buffer bytes.Buffer
		id := c.Query("id")                                                                   // 获取URL参数
		first_name := c.PostForm("first_name")                                                // 获取表单数据
		last_name := c.PostForm("last_name")                                                  // 获取表单数据
		stmt, err := db.Prepare("update person set first_name= ?, last_name= ? where id= ?;") //  准备执行语句
		// 判断是否准备成功
		if err != nil {
			fmt.Print(err.Error())
		}
		// 执行语句
		_, err = stmt.Exec(first_name, last_name, id)
		// 判断是否执行成功
		if err != nil {
			fmt.Print(err.Error())
		}
		// Fastest way to append strings
		// 将数据添加到buffer中
		buffer.WriteString(first_name) //   将数据添加到buffer中
		buffer.WriteString(" ")        //  将数据添加到buffer中
		buffer.WriteString(last_name)  // 获取buffer中的数据
		defer stmt.Close()             // 关闭语句
		name := buffer.String()        // 获取buffer中的数据
		// 返回JSON
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", name),
		})
	})
	// Delete resources
	// 添加一个路由，并且使用DELETE方法
	router.DELETE("/person", func(c *gin.Context) {
		// 获取URL参数
		id := c.Query("id")                                        // 获取URL参数
		stmt, err := db.Prepare("delete from person where id= ?;") // 准备执行语句
		// 判断是否准备成功
		if err != nil {
			fmt.Print(err.Error())
		}
		// 执行语句
		_, err = stmt.Exec(id)
		// 判断是否执行成功
		if err != nil {
			fmt.Print(err.Error())
		}
		// 返回JSON
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted user: %s", id),
		})
	})

	router.Run(":3000") //  启动服务
}
