package engine

import (
	"FirstParaseProject/fetcher"
	"log"
)

func Run(seeds ...Request)  {

	var requestes  []Request

	for _,r := range seeds{
		requestes = append(requestes,r)
	}
	for len(requestes) > 0 {

		r := requestes[0]
		requestes = requestes[1:]
		body ,err := fetcher.Fetch(r.Url)
		if err != nil {
           log.Printf("fetcher error :"+"fetcher url %s : %v",r.Url,err)
           continue
		}
		  parserResult := r.ParserFunc(body)
          requestes = append(requestes,parserResult.Requests...)

		for _,item := range parserResult.Items{
			log.Printf("Got item : %V",item)
		}


	}

}