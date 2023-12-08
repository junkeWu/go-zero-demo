### 1. "获取用户信息"

1. route definition

- Url: /user/info
- Method: POST
- Request: `UserInfoRequest`
- Response: `ResponseBsee`

2. request definition



```golang
type UserInfoRequest struct {
	UserId int64 `json:"userId"`
}
```


3. response definition



```golang
type ResponseBsee struct {
	Code int `json:"code"`
	Error string `json:"error"`
	Data interface{} `json:"data"`
}
```

