package main

import (
	"os"
	"scrapper/scrapper"

	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	scrapper.Scrape()
	return c.Attachment(fileName, "job.csv")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.GET("scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
