#Linkedin Scraper
This is a linkedin job scraper written in golang.

## How to run
Clone this repo:

```$ git clone https://github.com/eduardoeq/linkedin-scraper.git```

Go into the project's root folder.

```$ cd path/to/linkedin-scraper```

Run the project:.

```$ go run cmd/main.go```

Use the API:

GET endpoint `scrape`:

Available parameters:
- keywords
- excluded
- location
- remote

Example:
```http://localhost:1234/scrape?keywords=golang%20developer&excluded=junior&location=UK&remote=true```

Results:
```
    "status": 200,
    "message": "50 jobs found!",
    "jobs": [
        {
            "title": "Golang Developer (Remote)",
            "src": "https://uk.linkedin.com/jobs/view/golang-developer-remote-at-niufitel-s-l-3532219384?refId=XGkZ9%2FK%2FXDlV5kGbVx90pQ%3D%3D\u0026trackingId=HzwGHQg9VEAkCIyD5EKf3w%3D%3D\u0026position=1\u0026pageNum=0\u0026trk=public_jobs_jserp-result_search-card",
            "company": "Niufitel, S.L.",
            "companySrc": "https://es.linkedin.com/company/niufitel-s.l.?trk=public_jobs_jserp-result_job-search-card-subtitle",
            "location": "London, England, United Kingdom",
            "postedAt": "2023-03-20"
        },
        {
            "title": "Golang Developer",
            "src": "https://uk.linkedin.com/jobs/view/golang-developer-at-getground-3484082912?refId=XGkZ9%2FK%2FXDlV5kGbVx90pQ%3D%3D\u0026trackingId=48LcGCXfvhFRHK76IEupNA%3D%3D\u0026position=2\u0026pageNum=0\u0026trk=public_jobs_jserp-result_search-card",
            "company": "GetGround",
            "companySrc": "https://uk.linkedin.com/company/getground?trk=public_jobs_jserp-result_job-search-card-subtitle",
            "location": "London, England, United Kingdom",
            "postedAt": "2023-01-23"
        },
        ...
    ]
  ```