# UniSat Go SDK

## 安装 Install

```go
go get github.com/go-unisat/go-unisat
```

## 示例

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
defer cancel()
resp, err := unisat.GetBlockchainInfo(ctx, unisat.DefaultTest, bear)
if err != nil {
    log.Fatalf("GetBlockchainInfo error: %s", err)
}
```

## 联系作者

你可以在X推特([@alexgiantwhale](https://twitter.com/alexgiantwhale))和电报(alexgiantwhale)上找到我。