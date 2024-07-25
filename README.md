
# Web Scraping em Golang




## Installation

Install Web Scraping with go

```bash
mkdir nome da pasta
cd nome da pasta
go mod init nome do projeto
go get -u github.com/gocolly/colly
```


    
## Explanation of the Code

- Collector Creation: colly.NewCollector() creates a new collector to perform scraping.
- OnHTML callback:
- The .content-wrapper .article-item selector is used to find article items.
- Adjust as needed if the page structure changes. e.ChildText(".article-title") extracts the title of the article
- e.ChildText(".article-summary") extracts the article summary. e.ChildAttr(".article-title a", "href") extracts the link to the article
- OnRequest and OnError callbacks:
- OnRequest prints the URL being visited. OnError handles and logs request errors
- Visit URL: The c.Visit() method starts collecting data from the given URL


## Usage/Examples

```
go run main.go
```

