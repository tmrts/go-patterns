# Bridge Pattern
Decouples an interface from its implementation so that the two can vary independently

## Implementation

```go

// interface
type Request interface {
	HttpRequest() (*http.Request, error)
}

type Client struct {
	Client *http.Client
}

func (c *Client) Query(req Request) (resp *http.Response, err error) {
	httpreq,_:=req.HttpRequest()
	resp, err = c.Client.Do(httpreq)
	return
}

// implementation
type BaiduRequest struct {
}

func (cdn *BaiduRequest) HttpRequest() (*http.Request,  error) {
	return http.NewRequest("GET", "https://www.baidu.com", nil)
}

// implementation
type Googleequest struct {
}

func (cdn *Googleequest) HttpRequest() (*http.Request, error) {
	return http.NewRequest("GET", "https://www.google.com/?hl=zh-cn&gws_rd=ssl", nil)
}
```

## Usage
```go
client := &Client{http.DefaultClient}

baiduReq := &BaiduRequest{}
fmt.Println(client.Query(baiduReq))

googleReq := &Googleequest{}
fmt.Println(client.Query(googleReq))
```

