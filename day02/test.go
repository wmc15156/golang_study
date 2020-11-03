package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

type extractJob struct {
	title string
	company string
	location string
	summary string
}

func main () {
	// var jobs []extractJob
	totalLength := getPages()

	for i:=0; i < totalLength; i++ {

		if i == 0 {
			reflect.TypeOf(extractJobs)
			reflect.TypeOf(extractJobs)
			extractJobs := getPage(i)
			fmt.Println(reflect.TypeOf(extractJobs[0]))
		}
		// jobs = append(jobs, extractJobs...)
	}

}

func cleanString (str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)),"")
}

func getPage(number int) []extractJob{
	var jobs []extractJob
	pageURL := baseURL + "&start=" + strconv.Itoa(number * 50)
	fmt.Println(pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".jobsearch-SerpJobCard").Each(func(i int, selection *goquery.Selection) {
		job := extractJobs(selection)
		jobs = append(jobs, job)
	})

	return jobs
}

func extractJobs (selection *goquery.Selection) extractJob{
	title := cleanString(selection.Find(".title > a").Text())
	company := cleanString(selection.Find(".company").Text())
	location := cleanString(selection.Find(".location").Text())
	summary := cleanString(selection.Find(".summary").Text())

	return extractJob{
		title: title,
		company: company,
		location: location,
		summary: summary,
	}
}

func getPages() int {
	res, err := http.Get(baseURL)
	pageLength := 0
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, selection *goquery.Selection) {
		pageLength = selection.Find("a").Length()

	})
	return pageLength
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode (res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}