# go-news-ddd
(a try) Implementation of Domain Design Driven (DDD) on a news website (News, Topics and Tags)

## Requirements

Mandatory:
- A News could have many topics and tags
- When list (a) news will show the related topics and tags
- News can have status: draft, deleted and publish
- List all topics
- List all tags
- Filter news based on topics
- Filter news based on tags

Optional:
- Support pagination when getting all news
- TBD

## Build and deployment

```bash
cp .env_example .env
docker system prune --force
docker volume rm godddnews_postgresql godddnews_postgresql_data --force
docker-compose build
docker-compose up -d
```

## Product Items Backlog

- [X] **Mandatory:** Create REST API News & Tag CRUD
  - [X] News
    - [X] Get all
      - [X] Get news by its status
      - [X] Get news by its topic 
      - [X] Get news by its tag
      - [X] Get news by pagination
    - [X] Get by slug
    - [X] Get by id
    - [X] Get by topic
    - [X] Create
    - [X] Update
      - [X] Update by id
      - [X] Update by slug
    - [X] Delete
      - [X] Delete by id
      - [X] Delete by slug
  - [X] Topic
    - [X] Get all 
    - [X] Get by slug
    - [X] Get by id
    - [X] Create
    - [X] Update
      - [X] Update by id
      - [X] Update by slug
    - [X] Delete
      - [X] Delete by id
      - [X] Delete by slug
  - [X] Tag
    - [X] Get all
    - [X] Get by slug
    - [X] Get by id
    - [X] Create
    - [X] Update
      - [X] Update by id
      - [X] Update by slug
    - [X] Delete
      - [X] Delete by id
      - [X] Delete by slug
     
- [ ] **Mandatory:** API Functional Test (see interfaces/handlers/handler_test)
  - [ ] News
    - [ ] Get all
      - [ ] Get news by its status
      - [ ] Get news by its topic 
      - [ ] Get news by its tag
      - [ ] Get news by pagination
    - [ ] Get by slug
    - [ ] Get by id
    - [ ] Get by topic
    - [ ] Create
    - [ ] Update
      - [ ] Update by id
      - [ ] Update by slug
    - [ ] Delete
      - [ ] Delete by id
      - [ ] Delete by slug
  - [ ] Topic
    - [ ] Get all 
    - [ ] Get by slug
    - [ ] Get by id
    - [ ] Create
    - [ ] Update
      - [ ] Update by id
      - [ ] Update by slug
    - [ ] Delete
      - [ ] Delete by id
      - [ ] Delete by slug
  - [ ] Tag
    - [ ] Get all
    - [ ] Get by slug
    - [ ] Get by id
    - [ ] Create
    - [ ] Update
      - [ ] Update by id
      - [ ] Update by slug
    - [ ] Delete
      - [ ] Delete by id
      - [ ] Delete by slug

## Methodologies
- Create User Stories
- Design the database as per requirement:  
  - Relationship (Primary & Foreign Key)
  - Indexing
  - News Versioning (for optimistic lock versioning)
- Create model/struct as domain
- Create repository layer 
- Create layer/service layer
  - Json validation
  - Error management
  - Api version management
  - Activity Logging (ToDo)
  - Handling duplicate requests (ToDo)
- Development Stacks: 
  - Golang: GoChi, logrus
  - Gorm: hard delete, many2many
  - DB: Postgre
  - Migration: Auto Migration
- Deployment:
  - Docker & Docker compose

### Database Design ###
![Image Database Design](docs/images/database.png)


## URL ENDPOINT

#### /api/v1/news

- `GET` : Get all news
- `POST` : Create a news

#### /api/v1/news?status={status}

- `GET` : Get all news by its status

#### /api/v1/news?page_number=1&page_size=10 

- `GET` : Get all news with pagination

#### /api/v1/news/{id}

- `GET` : Get a news by id
- `PUT` : Update a news by id
- `DELETE` : Delete a news by id

#### /api/v1/news/{slug-name}

- `GET` : Get a news by slug
- `PUT` : Update a news by slug
- `DELETE` : Delete a news by slug

#### /api/v1/news/topic/{topic-slug}

- `GET` : Get all news by topic slug

#### /api/v1/news/tag/{tag-slug}

- `GET` : Get all news by tag slug

#### /api/v1/topic

- `GET` : Get all topic
- `POST` : Create a topic

#### /api/v1/topic/{id}

- `GET` : Get a topic by id
- `PUT` : Update a topic by id
- `DELETE` : Delete a topic by id

#### /api/v1/topic/{slug-name}

- `GET` : Get a topic by slug-name
- `PUT` : Update a topic by slug-name
- `DELETE` : Delete a topic by slug-name

#### /api/v1/tag

- `GET` : Get all tag
- `POST` : Create a tag

#### /api/v1/tag/{id}

- `GET` : Get a tag by id
- `PUT` : Update a tag by id
- `DELETE` : Delete a tag by id

#### /api/v1/tag/{slug-name}

- `GET` : Get a tag by slug-name
- `PUT` : Update a tag by slug-name
- `DELETE` : Delete a tag by slug-name

### Usage Examples

### News ###

Create News without ID

```bash
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
```

Create News with ID

```bash
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
```

List all news

```bash
curl localhost:8080/api/v1/news
```

List all news with pagination
```bash
curl localhost:8080/api/v1/news?page_number=1&page_size=10
```

List published News
```bash
curl localhost:8080/api/v1/news?status=publish
```

List deleted News
```bash
curl localhost:8080/api/v1/news?status=deleted
```

List draft News
```bash
curl localhost:8080/api/v1/news?status=draft
```

List all news with pagination with status
```bash
curl localhost:8080/api/v1/news?page_number=1&page_size=10?status=draft
```

Get news by its slug
```bash
curl --request GET --url "localhost:8080/api/v1/news/cara-makan-kenyang" 
```

Get news by its id
```bash
curl --request GET --url "localhost:8080/api/v1/news/10"
```

Update news by its id

```bash
curl --request PUT --url "localhost:8080/api/v1/news/10" --header 'content-type: application/json' --data \
'{
	"title": "Dialog antar warga",
	"slug": "dialog-warga",
	"content": "Dialog antar warg membuktikan bahwa jiwa masyarakat Indonesia masih mementingkan musyawarah.",
	"status": "publish",
	"topic_slugs": ["politics", "national"],
	"tag_slugs": ["other-tag", "national-tag"]
}'
```

Get news by topic's slug
```bash
curl --request GET --url "localhost:8080/api/v1/news/topic/politics" 
```

Get news by tag's slug
```bash
curl --request GET --url "localhost:8080/api/v1/news/tag/other-tag" 
```

Delete news by its id
```bash
curl --request DELETE --url "localhost:8080/api/v1/news/10" 
```

### Topic ###

Create a topic

```bash
curl --request POST \
  --url "localhost:8080/api/v1/topic" \
  --header 'content-type: application/json' \
  --data '{
    "id": 10,
	"name": "Hobi Berenang",
	"slug": "hobi-berenang"}'
```

Get a topic by slug
```bash
curl --request GET --url "localhost:8080/api/v1/topic/hobi-berenang"
```

Update topic by slug
```bash
curl --request PUT \
  --url "localhost:8080/api/v1/topic/hobi-berenang" \
  --header 'content-type: application/json' \
  --data '{
	"name": "Hobi Berenang Aja",
	"slug": "hobi-berenang-aja"}'
```
Delete topic by slug
```bash
curl --request DELETE --url "localhost:8080/api/v1/topic/hobi-berenang-aja" 
```
Get topic by id
```bash
curl --request GET --url "localhost:8080/api/v1/topic/1" 
```

Update topic by id
```bash
curl --request PUT \
  --url "localhost:8080/api/v1/topic/3" \
  --header 'content-type: application/json' \
  --data '{
	"name": "Hobi Membaca aja",
	"slug": "hobi-membaca"}'
```

Delete topic by id
```bash
curl --request DELETE --url "localhost:8080/api/v1/topic/3" 
```

### Tag ###


Create a tag

```bash
curl --request POST \
  --url "localhost:8080/api/v1/tag" \
  --header 'content-type: application/json' \
  --data '{
    "id": 10,
	"name": "Hobi Berenang",
	"slug": "hobi-berenang"}'
```

Get a tag by slug
```bash
curl --request GET --url "localhost:8080/api/v1/tag/hobi-berenang"
```

Update tag by slug
```bash
curl --request PUT \
  --url "localhost:8080/api/v1/tag/hobi-berenang" \
  --header 'content-type: application/json' \
  --data '{
	"name": "Hobi Berenang Aja",
	"slug": "hobi-berenang-aja"}'
```
Delete tag by slug
```bash
curl --request DELETE --url "localhost:8080/api/v1/tag/hobi-berenang-aja" 
```
Get tag by id
```bash
curl --request GET --url "localhost:8080/api/v1/tag/1" 
```

Update tag by id
```bash
curl --request PUT \
  --url "localhost:8080/api/v1/tag/3" \
  --header 'content-type: application/json' \
  --data '{
	"name": "Hobi Membaca aja",
	"slug": "hobi-membaca"}'
```

Delete tag by id
```bash
curl --request DELETE --url "localhost:8080/api/v1/tag/3" 
```

## Migration

Up <br>
`migrate -database "postgres://root:root@localhost:54322/godddnews_db?sslmode=disable" -path migration/scripts/ up`

Down <br>
`migrate -database "postgres://root:root@localhost:54322/godddnews_db?sslmode=disable" -path migration/scripts/ down`

Version <br>
`migrate -database "postgres://root:root@localhost:54322/godddnews_db?sslmode=disable" -path migration/scripts/ version`

## Notes
- To enable logging on DB query and log, comment the "PRODUCTION_ENV" in .env

## Author

author: bondhan.novandy@gmail.com/Depok/Indonesia

## References
- Gochi router: https://github.com/go-chi/chi/blob/master/_examples/rest/main.go
- DB Migration: https://github.com/golang-migrate/migrate
- DDD Skeleton : https://github.com/takashabe/go-ddd-sample
- DDD Implementation : https://github.com/jojoarianto/go-ddd-api
- GORM Documentation : http://doc.gorm.io
- API test: https://github.com/steinfletcher/apitest
