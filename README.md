# go_init
### https://github.com/golang/go
[Go 基本介紹](https://sz9751210.github.io/posts/go-variable-and-const/)
[nil](https://stackoverflow.com/questions/35983118/what-does-nil-mean-in-golang)


建立相關套件管理 & 初始化模塊

`go mod init 專案名稱`

安裝 Cobra 基礎函數庫

`go get -u github.com/spf13/cobra@v1.0.0`

驗證單字轉換功能

`go run main.go help word`
`go run main.go word -s=eddcycj -m=1`

驗證時間工具

`go run main.go time now`
`go run main.go time calc -c="2029-09-0412:02:33" -d=5m`
`go run main.go time calc -c="2029-09-0412:02:33" -d=-2h`
