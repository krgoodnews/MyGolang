package main

import (
	"os"
	"strings"

	scraper "github.com/krgoodnews/JobScraperApp/Scraper"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)

	term := strings.ToLower(scraper.CleanString(c.FormValue("term")))
	scraper.Scrape(term)

	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
