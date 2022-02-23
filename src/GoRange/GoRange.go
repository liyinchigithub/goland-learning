package GoRange

import (
	"fmt"
)

/*
	Go 语言中 range 关键字用于 for 循环中[迭代数组(array)]、[切片(slice)]、[通道(channel)]或[集合(map)]的元素。
	在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
*/

/*
	遍历切片（动态数组）
	第一个参数 匿名变量
	第二个参数 索引值
*/ 
func GoRangeA()  {
	  //这是我们使用range去求一个slice的和。使用数组跟这个很类似
	  nums := []int{2, 3, 4}
	  sum := 0
	  for _, num := range nums {
		  sum += num
	  }
	  //上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	  fmt.Println("sum:", sum)
	  //在数组上使用range将传入index和值两个变量。
}

/*
	遍历数组
	第一个参数 索引下标
	第二个参数 索引值
*/ 
func GoRangeB()  {
	// 声明一个切片（动态数组）
	nums:=[]int{1,2,3,4}
	// 遍历切片
	for i, num := range nums {
            fmt.Printf("num:%d，index：%d\n", num,i)
    }
}

/*
	遍历map集合
	第一个参数 map key
	第二个参数 map value
*/ 
func GoRangeC()  {
	//range也可以用在map的键值对上。
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
		/*
		输出
		b -> banana
		a -> apple
		*/
    }
}

/*
	遍历字符串 
	第一个参数 字符值
	第二个参数 unicode字符串
*/
func GoRangeD()  {
	//range也可以用来枚举Unicode字符串。
    for i, c := range "go" {
        fmt.Println(i, c)//第一个参数是[字符的索引]，第二个是字符（Unicode的值）本身
    }
}

/*
  遍历通道	chanel
*/
func GoRangeE()  {
	
}