package 期末

//go 基础
// （a)type People struct {
//    name string
//}
//
//func main() {
//    js := `{
//         "name":"坤坤"
//     }`
//    var p People
//    err := json.Unmarshal([]byte(js), &p)
//    if err != nil {
//        fmt.Println("err: ", err)
//        return
//    }
//    fmt.Println(p)
//}
//在json.Unmarshal中对js进行了强制类型转换为[]byte,就没能改变到p.name的值，而默认为"\0"，从而无输出

//(b)r中没有接受到值来储存
/*func main() {
	a := 2
	b := 0
	defer fmt.Println("2/0=", a/b)
	defer fmt.Println("3/2 =", 3/2)

	mytest(1)
	fmt.Println("3/2 =", 3/2)
}
func mytest(num int) {
	defer func() {

		if r := recover(); r != nil {

			fmt.Println("panic信息", r)
		}
	}()

	if num == 1 {
		panic("出现异常")
	}
	defer fmt.Println("3/2 =", 3/2)
}*/

// 并发相关
// (a)
/*func main() {
	var mu sync.Mutex

	go func() {
		fmt.Println("有点强人锁男")
		mu.Lock()
	}()

	mu.Unlock()
}*/
//在go func 的子程序中对其上锁,当没有解锁时会一直进行该程序 ，就会一直卡住，从而编译可能报错

//(b)
/*func main() {
	var Channel1 chan int
	Channel1 = make(chan int)//不带缓冲区，造成死锁

	go func() {
		fmt.Println("下山的路又堵起了")
		x := <-Channel1
		fmt.Println(x)
	}()
	Channel1 <- 10
	time.Sleep(100)//等待子线程
}*/

//(c)
/*func main() {
	var Channel1 chan int
	Channel1 = make(chan int)
	var count = 0

	go func() {
		fmt.Println("1下山的路又堵起了")
		count++
	}()
	go func() {
		fmt.Println("2下山的路又堵起了")
		count++
	}()
	go func() {
		fmt.Println("3下山的路又堵起了")
		count++
	}()
	go func() {
		fmt.Println("4下山的路又堵起了")
		count++
	}()
	go func() {
		fmt.Println("5下山的路又堵起了")
		count++
	}()
	go func() {
		fmt.Println("6下山的路又堵起了")
		count++
	}()
	go func() {
		fmt.Println("7下山的路又堵起了")
		count++
	}()
	go func() {
		fmt.Println("8下山的路又堵起了")
		count++
	}()
	go func() {
		fmt.Println("9下山的路又堵起了")
		count += 1
	}()
	go func() {
		fmt.Println("10下山的路又堵起了")
		count += 1
	}()
	for count = 0; count != 10; {
		time.Sleep(200)
	}
	if count == 10 {
		x := <-Channel1
		println(x)
	}
	Channel1 <- 10

}*/

//(d)func main() {
//	ch := make(chan int) // 创造通道
//	go func(ch chan int) {
//
//		for i := 2; ; i++ {
//			ch <- i // 传值
//
//		}
//
//	}(ch)
//	for {
//		p := <-ch
//		fmt.Print(p, " ")
//		ch1 := make(chan int)//每次增加协程来提高效率
//		go func(in, out chan int, p int) {
//			for {
//				i := <-in
//
//				if i%p != 0 {
//					out <- i    //通过通道进行筛选
//				}
//			}
//		}(ch, ch1, p)
//		ch = ch1
//	}
//}//思路源于博客

//数据库
//(a)CREATE TABLE [IF NOT EXISTS] 薪资(id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
//						员工 VARCHAR(20) NOT NULL PRIMARY KEY,
//						性别 CHAR(8) NOT NULL COMMENT '男，女'
//						工作 VARCHAR(20) NOT NULL,
//						业绩 FLOAT(5,2) NULL //						工资 VARCHAR(20))
//
//(b)主键是唯一确定的，可以是一种标识。而外键，是用于与另张表的连接，一张表的主键，就可以是另一张表的外键。主键只能有一个，
//而外键可以多个例如创建学生表，一个表存放学生信息，另个表存放学生借书，学号可以是第一个表的主键，而姓名，班级等就可以作为外键联系另一个表

//编程能力
//
//可以创建两二维数组，二维数组地址表示楼层数和房间号，一个计分变量count并初始化为0，一个临时计上楼f变量并初始化为0，一个进入数n记录进入房间;
//一个二维数组存放0.1来表示某楼某房间是否可以上楼，并统计可以上楼数，计入临时上楼数f;
//另个二维数组存放指示牌数;
//将count初始化为n 房间的指示牌数;
//将指示牌数%f可以得到第一次上楼进入房间(通过是否为1判断）——1
//将count+该房间的指示牌数——2
//重复1，2到楼层到顶楼，输出count
