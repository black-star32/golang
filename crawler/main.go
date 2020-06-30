package main

import (
	"golang/crawler/engine"
	"golang/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		 Url: "http://www.zhenai.com/zhenghun",
		 ParseFunc: parser.ParseCityList,
	})
	//resp, err := http.Get("http://www.zhenai.com/zhenghun")
	//if err != nil{
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("Error: statue cate", resp.StatusCode)
	//	return
	//}
	//
	//e := determineEncoding(resp.Body)
	//
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	//all, err := ioutil.ReadAll(utf8Reader)
	//if err != nil{
	//	panic(err)
	//}
	////fmt.Printf("%s\n", all)
	//printCityList(all)
}

//func printCityList(contents []byte) {
//	 re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`)
//	 //re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/zhuhai" data-v-2cb5b6a2>珠海</a>`)
//
//	 mathces := re.FindAllSubmatch(contents, -1)
//	 for _, m := range mathces{
//	 	fmt.Printf("city: %s, url: %s\n", m[2], m[1])
//	 	//for _, subMatch := range m{
//	 	//	fmt.Printf("%s\n", subMatch)
//		//}
//		//fmt.Println()
//	 }
//	 fmt.Printf("Matches found: %d\n", len(mathces))
//}
//
//
//func determineEncoding(r io.Reader) encoding.Encoding {
//	bytes, err := bufio.NewReader(r).Peek(1024)
//	if err != nil{
//		panic(err)
//	}
//	e, _, _ := charset.DetermineEncoding(bytes,"")
//	return e
//}