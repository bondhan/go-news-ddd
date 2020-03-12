
#!/bin/bash

echo Create News no ID
curl --request POST \
  --url "localhost:8080/api/v1/news" \
  --header 'content-type: application/json' \
  --data '{
	"title": "Shalat berjamaah",
	"slug": "cara-shalat",
	"content": "Cara shalat khusyu sesuai sunnah",
	"status": "publish",
	"topic_slugs": ["religion","national"],
	"tag_slugs":["other-tag","national-tag"]
}'


echo -e "\n"
echo Create News with ID
curl --request POST \
  --url "localhost:8080/api/v1/news" \
  --header 'content-type: application/json' \
  --data '{
  	"id":10,
	"title": "Makan murah dan kenyang",
	"slug": "cara-makan-kenyang",
	"content": "Cara makan kenyang dan murah",
	"status": "draft",
	"topic_slugs": ["other-topic","national"],
	"tag_slugs":["international-tag","national-tag"]
}'

echo -e "\n"
echo List all news
curl localhost:8080/api/v1/news

echo -e "\n"
echo List all news with pagination
curl localhost:8080/api/v1/news?page_number=1&page_size=10

echo -e "\n"
echo List published News
curl localhost:8080/api/v1/news?status=publish

echo -e "\n"
echo List deleted News
curl localhost:8080/api/v1/news?status=deleted

echo -e "\n"
echo List draft News
curl localhost:8080/api/v1/news?status=draft

echo -e "\n"
echo List all news with pagination with status
curl localhost:8080/api/v1/news?page_number=1&page_size=10?status=draft

echo -e "\n"
echo Get news by its slug
curl --request GET --url "localhost:8080/api/v1/news/cara-makan-kenyang" 

echo -e "\n"
echo Get news by its id
curl --request GET --url "localhost:8080/api/v1/news/10"

echo -e "\n"
echo Update news by its id
curl --request PUT --url "localhost:8080/api/v1/news/10" --header 'content-type: application/json' --data \
'{
	"title": "Dialog antar warga",
	"slug": "dialog-warga",
	"content": "Dialog antar warg membuktikan bahwa jiwa masyarakat Indonesia masih mementingkan musyawarah.",
	"status": "publish",
	"topic_slugs": ["politics", "national"],
	"tag_slugs": ["other-tag", "national-tag"]
}'

echo -e "\n"
echo Get news by topic slug
curl --request GET --url "localhost:8080/api/v1/news/topic/politics" 

echo -e "\n"
echo Get news by tag slug
curl --request GET --url "localhost:8080/api/v1/news/tag/other-tag" 

echo -e "\n"
echo #### Topic ####

echo -e "\n"
echo Create a topic
curl --request POST \
  --url "localhost:8080/api/v1/topic" \
  --header 'content-type: application/json' \
  --data '{
    "id": 10,
	"name": "Hobi Berenang",
	"slug": "hobi-berenang"}'

echo -e "\n"
echo Get a topic by slug
curl --request GET --url "localhost:8080/api/v1/topic/hobi-berenang"

echo -e "\n"
echo Update topic by slug
curl --request PUT \
  --url "localhost:8080/api/v1/topic/hobi-berenang" \
  --header 'content-type: application/json' \
  --data '{
	"name": "Hobi Berenang Aja",
	"slug": "hobi-berenang-aja"}'

echo -e "\n"
echo Delete topic by slug
curl --request DELETE --url "localhost:8080/api/v1/topic/hobi-berenang-aja" 

echo -e "\n"
echo Get topic by id
curl --request GET --url "localhost:8080/api/v1/topic/1" 

echo -e "\n"
echo Update topic by id
curl --request PUT \
  --url "localhost:8080/api/v1/topic/3" \
  --header 'content-type: application/json' \
  --data '{
	"name": "Hobi Membaca aja",
	"slug": "hobi-membaca"}'

echo -e "\n"
echo Delete topic by id
curl --request DELETE --url "localhost:8080/api/v1/topic/3" 

echo -e "\n"
echo echo echo  Tag echo echo echo 

echo -e "\n"
echo Create a tag
curl --request POST \
  --url "localhost:8080/api/v1/tag" \
  --header 'content-type: application/json' \
  --data '{
    "id": 10,
	"name": "Hobi Berenang",
	"slug": "hobi-berenang"}'

echo -e "\n"
echo Get a tag by slug
curl --request GET --url "localhost:8080/api/v1/tag/hobi-berenang"

echo -e "\n"
echo Update tag by slug
curl --request PUT \
  --url "localhost:8080/api/v1/tag/hobi-berenang" \
  --header 'content-type: application/json' \
  --data '{
	"name": "Hobi Berenang Aja",
	"slug": "hobi-berenang-aja"}'

echo -e "\n"
echo Delete tag by slug
curl --request DELETE --url "localhost:8080/api/v1/tag/hobi-berenang-aja" 

echo -e "\n"
echo Get tag by id
curl --request GET --url "localhost:8080/api/v1/tag/1" 

echo -e "\n"
echo Update tag by id
curl --request PUT \
  --url "localhost:8080/api/v1/tag/3" \
  --header 'content-type: application/json' \
  --data '{
	"name": "Hobi Membaca aja",
	"slug": "hobi-membaca"}'

echo -e "\n"
echo Delete tag by id
curl --request DELETE --url "localhost:8080/api/v1/tag/3" 

echo
echo -e "\n"