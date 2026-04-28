# Gin 学习

## Router

创建一个gin.Default()对象可以开启路由



## GET

此部分与spring boot处理http请求对应说明

```go
c := *gin.Context
```



### @PathVariable 获取路径参数

- :name - 匹配单个路径参数
- *action - 匹配前缀之后的所有内容

通过c.Param(名称)获取对应的GET请求参数，并添加到路径后。

```go
// This handler will match /user/john but will not match /user/ or /user
  router.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)
  })
// However, this one will match /user/john/ and also /user/john/send
// If no other routers match /user/john, it will redirect to /user/john/
  router.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")
    action := c.Param("action")
    message := name + " is " + action
    c.String(http.StatusOK, message)
  })
```

### query 参数

对应/user/search?id=1这种路径，包含query参数id，可以使用以下方法查询query参数：

```go
c.Query(参数名)
c.DefaultQuery(参数名, 默认值)
```

