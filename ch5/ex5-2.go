package main

import "fmt"

func study(name string, math int, english int, science int) {
	fmt.Println(name+" 님 평균점수는", (math+english+science)/3, "입니다")
}
