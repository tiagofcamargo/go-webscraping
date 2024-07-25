package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	// URL alvo
	url := "url do site alvo"

	// Criar um coletor com suporte a cookies
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
		colly.AllowURLRevisit(),
		colly.CacheDir("./cache"),
	)

	var links []string
	var images []string
	var pageTitle string
	var metaDescription string

	// Adicionar cabeçalhos adicionais
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visitando:", r.URL.String())
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.5")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
	})

	// Capturar o título da página
	c.OnHTML("title", func(e *colly.HTMLElement) {
		pageTitle = e.Text
	})

	// Capturar a meta descrição
	c.OnHTML(`meta[name="description"]`, func(e *colly.HTMLElement) {
		metaDescription = e.Attr("content")
	})

	// Capturar todos os links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		links = append(links, link)
	})

	// Capturar todas as imagens
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		img := e.Attr("src")
		images = append(images, img)
	})

	// Tratar erros
	c.OnError(func(_ *colly.Response, err error) {
		log.Fatalf("Erro ao fazer a requisição: %v", err)
	})

	// Iniciar a requisição
	c.Visit(url)

	// Gerar relatório em HTML
	generateReport(pageTitle, metaDescription, links, images)
}

func generateReport(title, description string, links, images []string) {
	report := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Relatório de Coleta</title>
    </head>
    <body>
        <h1>Relatório de Coleta</h1>
        <h2>Título da Página</h2>
        <p>` + title + `</p>
        <h2>Meta Descrição</h2>
        <p>` + description + `</p>
        <h2>Links</h2>
        <ul>
    `

	for _, link := range links {
		report += `<li><a href="` + link + `">` + link + `</a></li>`
	}

	report += `
        </ul>
        <h2>Imagens</h2>
    `

	for _, img := range images {
		report += `<p>` + img + `</p>`
	}

	report += `
    </body>
    </html>
    `

	// Salvar relatório em um arquivo HTML
	file, err := os.Create("relatorio.html")
	if err != nil {
		log.Fatalf("Erro ao criar o arquivo de relatório: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(report)
	if err != nil {
		log.Fatalf("Erro ao escrever no arquivo de relatório: %v", err)
	}

	fmt.Println("Relatório gerado com sucesso: relatorio.html")
}
