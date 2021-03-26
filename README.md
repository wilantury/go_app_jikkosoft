# Web Backend Application: Golang / Go

### This is an application that exposes two endpoints. CRUD to manage Users and an endpoint to get sorted a given Array with integers numbers
  
## Technologies
>- Golang / Go
## Modules
> - github.com/gorilla/mux
## How to start
> - Clone repo: ```git clone https://github.com/wilantury/go_app_jikkosoft```
> - ```cd go_app_jikkosoft```
> - ```go run main.go```

## Request: POST
### Body:
```
{
	"unsorted":[1,2,3,4,2,4,78,23,12,1]
}
```
> Development host: ```localhost:5000/api/sort``` 