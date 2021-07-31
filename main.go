package main

import (
	"PlagDoctor/scrapper"
	"strings"

	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {

	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	//fmt.Println(c.FormValue("term"))
	scrapper.Scrape(term)
	return c.Attachment("jobs.csv", "jobs.csv")
}

func main() {
	//scrapper.Scrape("python")
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
