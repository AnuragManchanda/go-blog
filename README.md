# go-blog
**Blog APIs Build in Golang**

To run -> **docker-compose up --build**

server will run on localhost:80

**API Specs**

GET /articles\
curl --location --request GET 'http://localhost:80/articles'

GET /articles/2\
curl --location --request GET 'http://localhost:80/articles/2'

POST /articles
curl --location --request POST 'http://localhost:80/articles' \
--header 'Content-Type: application/json' \
--data-raw '{
	"title": "anurag1",
	"author": "anurag_author",
	"content": "content_anurag"
}'
