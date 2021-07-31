package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/%EC%B7%A8%EC%97%85?q=python&limit=100"

func main() {

	pages := getPages(baseURL)
	fmt.Println(pages)
	for i := 0; i < pages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageUrl := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("requesting", pageUrl)
	resp, err := http.Get(pageUrl)
	checkErr(err)
	checkResp(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")

	//searchCards := doc.Find(".mosaic-provider-jobcards")
	searchCards.Each(func(i int, card *goquery.Selection) {
		link, _ := card.Attr("href")
		fmt.Println("link= ", link)
		title := card.Find(".jobTitle>span").Text()
		fmt.Println("title= ", title)

		location := card.Find(".companyLocation").Text()
		fmt.Println("location= ", location)

	})
}

//페이지 가져오기
func getPages(url string) int {
	pages := 0
	resp, err := http.Get(url)
	checkErr(err)
	checkResp(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages

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
