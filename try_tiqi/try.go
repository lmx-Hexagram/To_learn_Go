package main

import (
	//"fmt"
	//"io"
	//"net/http"
	//"os"
	"fmt"
)

const (
	sunday = iota
	monday
	tuseday
	wednesday
	thursday
	friday
	saturday
)

func main() {
	a := 100
	var b int = 200

	fmt.Println(sunday)
	fmt.Println(saturday)

	number := [6]int{1, 2, 3, 4, 5} //没赋值默认是0
	for i, x := range number {
		fmt.Printf("第%d位的值 = %d\n", i, x)
	}

	println(max(a, b))
	//resp, error := http.Get("https://api.caiyunapp.com/v2/TAkhjf8d1nlSlspN/121.6544,25.1552/realtime.json ")
	//if error != nil {
	//deal with error
	//	return
	//}
	//defer resp.Body.Close()
	//io.Copy(os.Stdout, resp.Body)
}

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
