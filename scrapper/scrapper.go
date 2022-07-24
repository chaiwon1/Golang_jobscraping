package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// struct 만들기
type extractedJob struct {
	position string
	company  string
	location string
	status   string
	period   string
}

// main 함수
func Scrape() {
	var jobs []extractedJob

	// getPage 함수 실행해서 결과물(job들 가득 담긴 리스트) 받기
	extractedJob := getPage()

	// jobs에 반복해서(...) extractedJob 넣어주고 jobs라는 이름으로 정리
	jobs = append(jobs, extractedJob...)

	// to csv
	writeJobs(jobs)
	fmt.Println(len(jobs))
}

// url 접속해서 정보들 긁어모으는 함수
func getPage() []extractedJob {
	var jobs []extractedJob
	c := make(chan extractedJob)

	// url 확인
	url := "https://kr.linkedin.com/jobs/jobs-in-%EC%84%9C%EC%9A%B8-%EC%9D%B8%EC%B2%9C-%EC%A7%80%EC%97%AD?trk=homepage-basic_intent-module-jobs&currentJobId=3163953224&position=1&pageNum=0"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// HTML 읽기
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	items := doc.Find("ul.jobs-search__results-list>li")

	// 각 item마다 돌면서 extractJob 함수 실행해서 받은 정보들 jobs라는 빈 리스트에 넣기
	items.Each(func(i int, s *goquery.Selection) {
		go extractJob(s, c)
	})

	for i := 0; i < items.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	return jobs
}

// item들마다 정보 긁어오는 함수
func extractJob(s *goquery.Selection, c chan<- extractedJob) {
	position := cleanString(s.Find(".base-search-card__title").Text())
	company := cleanString(s.Find(".base-search-card__subtitle").Text())
	location := cleanString(s.Find(".job-search-card__location").Text())
	status := cleanString(s.Find(".result-benefits__text").Text())
	period := cleanString(s.Find(".job-search-card__listdate").Text())

	c <- extractedJob{
		position: position,
		company:  company,
		location: location,
		status:   status,
		period:   period,
	}
}

// text cleaning 해주는 함수
func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

// 최종 결과물을 csv로 담는 함수
func writeJobs(jobs []extractedJob) {
	file, _ := os.Create("jobs.csv")

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"position", "company", "location", "status", "period"}

	w.Write(headers)

	for _, job := range jobs {
		jobSlice := []string{job.position, job.company, job.location, job.status, job.period}
		w.Write(jobSlice)
	}
}
