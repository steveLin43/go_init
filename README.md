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

## Go 的效能與除錯套件介紹
1. 效能分析器 PProf: top 指令、trace 指令、list 指令
2. 追蹤剖析 trace
3. GODEBUG: 排程追蹤(schedtrace 指令、scheddetail 指令)、GC 追蹤(gctrace 指令)
4. 處理程序診斷工具 gops: 有大量功能，常用的包含: 檢視指定處理程序資訊(<pid> 指令)、檢視呼叫堆疊(stack 指令)、檢視記憶體使用(memstats 指令)、檢視執行狀況(stats 指令)、檢視 trace(trace 指令)
5. expvar: 公開和發布度量指標
6. Prometheus: 利用技術堆疊來公開和發布度量指標

## 逃逸分析
堆積 Heap: 人為手動管理，一般儲存較大物件。分配相對較慢，有關的指令相對較多。如果是多處參考的 func 也會放在這裡，避免 func 結束後東西就不見了。
堆疊 Stack: 編譯器管理，自動申請、分配、釋放，一般不太大。常見的函數、參數、區域變數等都會放在堆疊上。如果是未確定類型，例如 interface{}，也會放在這裡。

### 逃逸分析判斷
1. 查看逃逸過程，利用編譯器的指令
```
go build -gcflags '-m -l' main.go
```
2. 透過反編譯指令檢視
```
go tool compile -S main.go
```