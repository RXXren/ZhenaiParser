package parser

import (
	"FirstParaseProject/engine"
	"regexp"
)

func ParseCityList(contents []byte) engine.ParserResult {

	//rex := regexp.MustCompile(`{cityList:.*,order:"[A-Z]"}`)
	rex := regexp.MustCompile(`<a href="(http://m.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := rex.FindAllSubmatch(contents,-1)
	result := engine.ParserResult{}
	for _, m := range matches{
		result.Items = append(result.Items,m[2])
				result.Requests = append(result.Requests,engine.Request{
					Url:string(m[1]),
		           ParserFunc:engine.NilParser,
				})

	}
	return result
	//for _,cityes := range matches{
	//	text := string(cityes[:])
	//	rexcity := regexp.MustCompile(`http://m.zhenai.com/zhenghun/([a-zA-Z0-9]+)`)
	//	matchCity := rexcity.FindAllStringSubmatch(text,-1)
	//	for _,city := range matchCity{
	//		result.Items = append(result.Items,city[1])
	//		//fmt.Printf("cityName:%s\n",city[0])
	//		result.Requests = append(result.Requests,engine.Request{
	//			Url:string(city[0]),
    //            ParserFunc:engine.NilParser,
	//		})
	//	}
	////	fmt.Println("count:",len(matchCity))
	//
	//}

}