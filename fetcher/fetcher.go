package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)
func determinEncoding(r io.Reader )encoding.Encoding  {

	bytes ,err := bufio.NewReader(r).Peek(1024)
	if err !=nil {
		log.Printf("fetcher error:%v",err)
		return unicode.UTF8
	}

	e,_,_  :=	charset.DetermineEncoding(bytes,"")

	return  e

}

func Fetch(url string) ([]byte, error)  {

	respose , err  :=	http.Get(url)
	if err != nil {
		return nil,err
	}
	defer  respose.Body.Close()

	if respose.StatusCode != http.StatusOK {
		//生成err
	   //	errors.New()
	  return nil,fmt.Errorf("wrog error code:%d",respose.StatusCode)

	}
	//如果网页版本是GBK编码 则抓取下来的中文就是乱码
	e := determinEncoding(respose.Body)
	utf8Reader := transform.NewReader(respose.Body,e.NewDecoder())
	all , err := ioutil.ReadAll(utf8Reader)
	return  all,err


}