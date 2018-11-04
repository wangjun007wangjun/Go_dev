package add

import "fmt"

var Name string = "Hellow"
var PreName string = "World"

//init()函数会自动被执行
func init() {
	fmt.Println("init被自动执行了")
}
