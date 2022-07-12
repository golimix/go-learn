!!! 不使用注入 
 
```go
type App struct {
}

# 假设这个方法将会匹配并处理 GET /biu/<id> 这样的请求
func (a *App) GetData(id string) string {
    # todo: write your data query
    return "some data"
}

func NewApp() *App {
    return &App{}
}

app := App()
app.Run()
```