package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/%EC%B7%A8%EC%97%85?q=python&limit=100"

func main() {

	pages := getPages(baseURL)
	fmt.Println(pages)

}

//페이지 가져오기
func getPages(url string) int {
	resp, err := http.Get(url)
	checkErr(err)
	checkResp(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Find("a"))
	})
	return 0

}

//에러 처리
func checkErr(err error) {

	if err != nil {
		log.Fatalln(err)
	}
}

func checkResp(resp *http.Response) {

	//fmt.Println("resp.StatusCode = ", resp.StatusCode, "resp.Status = ", resp.Status)

	if resp.StatusCode != 200 {
		log.Fatalln("request failed with status : %d %s", resp.StatusCode, resp.Status)
	}
}
