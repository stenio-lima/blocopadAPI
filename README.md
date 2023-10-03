# BlocopadAPI
Blocopad is a simple and easy to use online notepad.

# Run database: you need run a redis instance with Docker
```bash
docker run -p 6379:6379 --name some-redis -d redis
```
# Run application with:
```bash
cd cmd
go run main.go
```
# POST and GET requests:
```bash
POST: curl -i --header "Content-Type: application/json" --request POST --data '{"data" : "save this", "onetime" : false}' http://localhost:8080/api/note
GET: curl -i  http://localhost:8080/api/note/{id}
```
# OneTime
The onetime function allows such a note to be viewed once and deleted after being read.
```bash
onetime: curl -i --header "Content-Type: application/json" --request POST --data '{"data" : "save this one time", "onetime" : true}' http://localhost:8080/api/note
```
# Read more about the main technologies used in this project:
- <a href="https://github.com/redis/go-redis">Redis</a>
- <a href="github.com/google/uuid">Google uuid</a>
- <a href="github.com/gorilla/mux">Gorilla Mux</a>

