# go-swagger-ui



#### go  get 


```
go get  github.com/aiganbao/go-swagger-ui

```


### dep ensure


```go
dep ensure -v -add  github.com/aiganbao/go-swagger-ui
```



### design dsl

```go


var _ = Resource("swaggerUI", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})

	Files("/swagger-ui/*filepath", "swagger-ui/")
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Files("/swagger.json", "swagger/swagger.json")
})

```


### Mount swagger  controller

```go
// Mount "swagger" controller
c3 := NewSwaggerUIController(service)
c3.FileSystem = swagger.SwaggerFs
app.MountSwaggerUIController(service, c3)
c4 := NewSwaggerController(service)
app.MountSwaggerController(service, c4)

```


