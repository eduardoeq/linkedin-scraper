package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
	"golang.org/x/exp/slices"
)

const (
	LETTER_BYTES       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SHOW_BODY_ON_ERROR = false

	MAIN_URL     = "https://www.linkedin.com/"
	SEARCH_ROUTE = "jobs/search/"

	JOB_SELECTOR     = ".base-card"
	SRC_SELECTOR     = ".base-card__full-link"
	TITLE_SELECTOR   = "div.base-search-card__info > h3"
	COMPANY_SELECTOR = "div.base-search-card__info > h4 > a"

	LOCATION_SELECTOR  = "div.base-search-card__info > div > span"
	POSTED_AT_SELECTOR = "div.base-search-card__info > div > time"
)

// var searchUrl string = fmt.Sprintf("%s%s?location=%s&keywords=%s&start=", MAIN_URL, SEARCH_ROUTE, LOCATION, KEYWORDS)
var start int = 0

type job struct {
	Title      string `json:"title"`
	Src        string `json:"src"`
	Company    string `json:"company"`
	CompanySrc string `json:"companySrc"`
	Location   string `json:"location"`
	PostedAt   string `json:"postedAt"`
}

type jobResonse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Jobs    []job  `json:"jobs"`
}

type searchParameters struct {
	Keywords         string `json:"keywords"`
	ExcludedKeywords string `json:"excluded"`
	Location         string `json:"location"`
	Remote           string `json:"remote"`
}

var jobs []job

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = LETTER_BYTES[rand.Intn(len(LETTER_BYTES))]
	}
	return string(b)
}

func scrape(context *gin.Context) {
	jobs = nil
	start = 0

	keywords := context.Query("keywords")
	excluded := context.Query("excluded")
	location := context.Query("location")
	remote := context.Query("remote")

	searchUrl := fmt.Sprintf(
		"%s%s?location=%s&pageNum=0&keywords=%s NOT \"%s\"&f_WT=%s&start=",
		MAIN_URL, SEARCH_ROUTE, location, keywords, excluded, remote,
	)

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		if SHOW_BODY_ON_ERROR {
			body := []string{}

			for _, element := range r.Body {
				body = append(body, string(element))
			}
		} else {
			fmt.Println("Request URL: ", r.Request.URL, "\nError: ", err)
		}
	})

	c.OnHTML(JOB_SELECTOR, func(e *colly.HTMLElement) {
		jobTitle := e.ChildText(TITLE_SELECTOR)
		jobSrc := e.ChildAttr(SRC_SELECTOR, "href")
		company := e.ChildText(COMPANY_SELECTOR)
		companySrc := e.ChildAttr(COMPANY_SELECTOR, "href")
		location := e.ChildText(LOCATION_SELECTOR)
		postedAt := e.ChildAttr(POSTED_AT_SELECTOR, "datetime")

		foundJob := job{
			jobTitle,
			jobSrc,
			company,
			companySrc,
			location,
			postedAt,
		}

		isNewJob := !slices.Contains(jobs, foundJob)
		if isNewJob {
			jobs = append(jobs, foundJob)
		}
	})

	for start < 50 {
		c.Visit(fmt.Sprintf("%s%d", searchUrl, start))
		start += 25
		time.Sleep(1 * time.Second)
	}

	response := jobResonse{
		http.StatusOK,
		fmt.Sprintf("%d jobs found!", len(jobs)),
		jobs,
	}

	if len(jobs) > 0 {
		context.IndentedJSON(http.StatusOK, response)
		return
	}

	context.IndentedJSON(http.StatusBadGateway, "Something wrong happened")
}

func main() {
	router := gin.Default()
	router.GET("/scrape", scrape)
	router.Run("localhost:1234")
}
