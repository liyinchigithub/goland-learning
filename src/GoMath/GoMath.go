package GoMath

import (
	crypto_rand "crypto/rand" // 加密随机数
	"encoding/base64"
	"fmt"
	"math/big"
	math_rand "math/rand"
	"time"
)

/*
	生成随机整数（固定范围内）
	[官方文档]http://word.topgoer.com/pkg/math_rand.htm
	[参考]https://www.cnblogs.com/niuben/p/13925133.html
*/

func TestMathRandomA1() {
	// 调用rand的方法生成伪随机int值
	fmt.Println(math_rand.Int())// 5577006791947779410
	fmt.Println(math_rand.Int31())// 2019727887
	fmt.Println(math_rand.Intn(5))// 3
}

func TestMathRandomA2() {
	for i := 0; i < 10; i++ {
		fmt.Println(math_rand.Intn(100)) //返回[0,100)的随机整数
	}
	// 当代码运行多次发现时，结果都是一样的。不管怎么运行代码，产生的结果都是这三个数，不会变。为什么？这是因为我们还没有设置随机数种子seed。
}

func TestMathRandomB() {
	math_rand.Seed(time.Now().UnixNano()) // 取纳秒时间戳，可以保证每次的随机数种子都不同
	fmt.Println(math_rand.Intn(100)) //返回[0,100)的随机整数
}

/*
	生成随机指定位数的整数，当位数不够时，通过前边补0达到长度一致
*/
func TestMathRandomC() {
	for i := 0; i < 10; i++ {
		// 控制位数	%.4d
		fmt.Printf("%.4d", math_rand.Int31()%10000) // 8081 7887 1847 4059 2081 1318 4425 2540 0456 3300
	}
}

/*
	生成随机加密串
	[官方文档]http://word.topgoer.com/pkg/crypto_rand.htm
*/
func TestMathRandomD1() {
	n, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(128))
	if err == nil {
		fmt.Println("rand.Int：", n, n.BitLen())
	}
	//2、Prime
	p, err := crypto_rand.Prime(crypto_rand.Reader, 5)
	if err == nil {
		fmt.Println("rand.Prime：", p)
	}
	//3、Read
	b := make([]byte, 32)
	m, err := crypto_rand.Read(b)
	if err == nil {
		fmt.Println("rand.Read：", b[:m])
		fmt.Println("rand.Read：", base64.URLEncoding.EncodeToString(b))
	}
}
// 将随机的byte值填充到b 数组中，以供b使用
func TestMathRandomD2() {
	b := make([]byte, 20)
    fmt.Println(b)

    _, err := math_rand.Read(b)
    if err != nil {
        fmt.Println(err.Error())
    }

    fmt.Println(b)
}

/*
	随机指定位数的，当位数不够是，可以通过前边补0达到长度一致
*/
func TestMathRandomE() {
	math_rand.Seed(time.Now().UnixNano()) // 取纳秒时间戳，可以保证每次的随机数种子都不同
	fmt.Println(math_rand.Int())
	fmt.Println(math_rand.Int31())
	fmt.Println(math_rand.Intn(5))
}

/*
	固定选项中，随机选择一个
*/
func TestMathRandomF() {
	math_rand.Seed(42) // Try changing this number!
	// 数组
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	fmt.Println("Magic 8-Ball says:", answers[math_rand.Intn(len(answers))])
}
