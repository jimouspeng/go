/*
 * @Author: jimouspeng
 * @Date: 2022-04-25 10:17:19
 * @Description: 第一个go文件
 * go run xxx 执行代码
 * go build xxx 生成二进制文件
 * @FilePath: \go\helloworld.go
 */
package main // 定义包名 你必须在源文件中非注释的第一行指明这个文件属于哪个包

 import "fmt" // 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数

func main() {
	fmt.Println("HELLO WORLD")
	var jimous string = "jimous is cool";
	var jimous2 int;
	var jimous3, jimous4, jimous5 int
	fmt.Println(jimous, '\n')
	fmt.Println(jimous, '\n')
	fmt.Println(jimous2, '\n')
	fmt.Println(jimous3, '\n')
	fmt.Println(jimous4, '\n')
	fmt.Println(jimous5, '\n')
}
