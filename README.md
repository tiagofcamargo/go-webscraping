# go-webscraping

# Comandos necessários

mkdir nome da pasta
cd nome da pasta
go mod init nome do projeto
go get -u github.com/gocolly/colly

# Explicação do Código
# Criação do Coletor: O colly.NewCollector() cria um novo coletor para realizar o scraping.

# Callback OnHTML:

# O seletor .content-wrapper .article-item é usado para localizar os itens de artigo. Ajuste conforme necessário se a estrutura da página mudar.
# e.ChildText(".article-title") extrai o título do artigo.
# e.ChildText(".article-summary") extrai o resumo do artigo.
# e.ChildAttr(".article-title a", "href") extrai o link para o artigo.
# Callbacks OnRequest e OnError:

# OnRequest imprime a URL sendo visitada.
# OnError trata e registra erros de requisição.
# Visitar a URL: O método c.Visit() inicia a coleta de dados da URL fornecida.

## Comando para rodar o código

go run main.go