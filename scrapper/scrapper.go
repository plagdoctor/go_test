package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
}

var baseURL string = "https://kr.indeed.com/%EC%B7%A8%EC%97%85?q=python&limit=100"

func main() {

	pages := getPages(baseURL)
	fmt.Println(pages)
	for i := 0; i < pages; i++ {
		getPage(i)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
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

	//searchCards := doc.Find(".mosaic-provider-jobcards")
	searchCards := doc.Find(".tapItem")

	var records [][]string
	var oneRecord []string

	searchCards.Each(func(i int, card *goquery.Selection) {
		oneRecord = extractJob(card)
		records = append(records, oneRecord)
	})
	extractAsCsv(records)
}

//추출 하기  .class>하위
func extractJob(card *goquery.Selection) []string {

	link, _ := card.Attr("href")
	//fmt.Println("link= ", cleanString(link))
	title := card.Find(".jobTitle>span").Text()
	//fmt.Println("title= ", cleanString(title))

	location := card.Find(".companyLocation").Text()
	//fmt.Println("location= ", cleanString(location))

	var oneRecord []string

	oneRecord = []string{link, title, location}

	return oneRecord
}

//csv 파일 생성 리턴할땐 () 로 싸주자
func extractAsCsv(records [][]string) {

	var file *os.File
	var err error

	file, err = os.Open("./output.csv")
	if err != nil {
		file, err = os.Create("./output.csv")
	} else if err == nil {
		//fmt.Println("title= ", cleanString(title))
	}
	defer file.Close()

	// csv writer 생성
	wr := csv.NewWriter(bufio.NewWriter(file))
	defer wr.Flush()

	// csv 내용 쓰기
	wr.WriteAll(records)

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
