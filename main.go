package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/itspage/gofetch/product"
	"fmt"
	"github.com/itspage/gofetch/downloader"
)

func main() {
	app := cli.NewApp()
	app.Name = "gofetch"
	app.Version = "0.0.1"
	app.Usage = "Scrape websites and format results."
	app.Author = "Alexander Page <alex@pagetek.co.uk>"
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "url",
			Usage: "Start scraping from `URL`",
		},
	}
	app.Action = run

	app.Run(os.Args)
}

func run(c *cli.Context) error {
	url := c.String("url")
	if url == "" {
		cli.ShowSubcommandHelp(c)
		return cli.NewExitError("Missing URL flag", 1)
	}
	downloader := new(downloader.HTMLDownloader)
	productList, err := product.NewProductListFromDownloader(downloader, url)
	if err != nil {
		fmt.Println("Download failed", err.Error())
	}
	fmt.Println(productList.JSON())
	return nil
}
