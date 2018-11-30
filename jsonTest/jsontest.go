package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

type Citylist struct {
	LinkContent string
	LinkURL string

}
type Cityslice struct {
	 Citylists []Citylist
}

func main() {
	var s Serverslice
	str := `{
              "servers":
               [
                {"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},
                {"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}
               ]
             }`

	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	fmt.Println(s.Servers[0].ServerIP)

	strss := `{
              "citylists":
               [
                {linkContent:"aba",linkURL:"http://m.zhenai.com/zhenghun/aba"},
                {linkContent:"akesu",linkURL:"http://m.zhenai.com/zhenghun/akesu"}
               ]
             }`

	var css Cityslice
	err := json.Unmarshal([]byte(strss),&css)
	if err != nil {
		println(err)
	}
	fmt.Println("wode:",css)


	}