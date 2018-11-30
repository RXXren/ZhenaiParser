package main

import (
	"FirstParaseProject/engine"
	"FirstParaseProject/zhenai/parser"
)

func  printCityList(content []byte)  {



}

func main() {


	engine.Run(engine.Request{

		Url:"http://m.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,

	})







}
