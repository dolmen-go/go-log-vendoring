package log

import "fmt"

func Println(v ...interface{}) {
	fmt.Println("Custom logger:", fmt.Sprint(v...))
}
