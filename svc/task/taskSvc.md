Helps to manage dependency injection, but manually, unlike the built in C# builder.Services: 
```go
// svc > task
package task

func NewHandler(entClient *ent.Client) *handler.Task {
	return handler.New(entClient)
}

```
```go
// main.go
	taskHdl := task.NewHandler(entClient)
```

Example addition:
```go
func NewHandler(entClient *ent.Client) *handler.Task {
	taskSvc := service.New(entClient, redisClient)
	return handler.New(taskSvc)
}
```