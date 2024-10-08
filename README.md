# go_init

參考資料：《用 Go 語言完成 6 個大型專案》第一章節 + https://github.com/go-programming-tour-book/tour

### https://github.com/golang/go
[Go 基本介紹](https://sz9751210.github.io/posts/go-variable-and-const/)
[nil](https://stackoverflow.com/questions/35983118/what-does-nil-mean-in-golang)


建立相關套件管理 & 初始化模塊 (main.go 跟 go.mod 同一路徑才吃的到設定)( package 名稱盡量與資料夾名稱相同)

```
go mod init 專案名稱
go mod tidy
```

安裝 Cobra 基礎函數庫

```
go get -u github.com/spf13/cobra@v1.0.0
```

驗證單字轉換功能

```
go run main.go help word
go run main.go word -s=eddcycj -m=1
```

驗證時間工具

```
go run main.go time now
go run main.go time calc -c="2029-09-0412:02:33" -d=5m
go run main.go time calc -c="2029-09-0412:02:33" -d=-2h
```

驗證資料庫轉換

```
go run main.go sql struct --username 帳號 --password 密碼 --db=名稱 --table "表名"
```
