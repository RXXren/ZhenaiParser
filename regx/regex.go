package main

import (
	"fmt"
	"regexp"
)

const  text  = `this is a email jakerenDevil@163.com@abc.com 
    email1 is adb@org.com
    email2 is     ucsbjc@qq.com
    email3 is ddd@qq.com.cn
`

func main() {

	   //正则表达式匹配器
	   //.+ 一个或者多个字符 .* 0个或者多个字符
	   //[a-zA-Z0-9] 表示匹配字符和数字，空格，下划线等不能匹配
	   //() 正则表达式提取的功能
     //rex  := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
     rex := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)+(\.[a-zA-Z0-9.]+)`)
	// match := rex.FindString(text)
    // match := rex.FindAllString(text,-1)
     match := rex.FindAllStringSubmatch(text,-1)//获取加括号部分的字符
	for _,m := range match{
		fmt.Println(m)
	}

}
