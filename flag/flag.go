/*
// 1. 有設定值重複而覆蓋的問題
// go run main.go -name=eddycjy -n=test
func main() {
    var name string
    flag.StringVar(&name, "name", "Go 語言初見", "說明")
    flag.StringVar(&name, "n", "Go 語言初見", "說明")
    flag.Parse()

    log.Printf("name: %s", name)
}
*/

/*
// go run main.go go -name=eddycjy
// go run main.go php -n=test
var name string
func main() {
    flag.Parse()  //先解析命令列，後分配，避免重複而覆蓋的問題
    goCmd := flag.NewFlagSet("go", flag.ExitOnError)
    goCmd.StringVar(&name, "name", "Go 語言", "說明")
    phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
    phpCmd.StringVar(&name, "n", "PHP 語言", "說明")

    args := flag.Args()
    if len(args) <=0 {
        return
    }

    switch args(0) {
        case "go":
            _= goCmd.Parse(args[1:])
        case "php":
            _= phpCmd.Parse(args[1:])
    }

    log.Printf("name: %s", name)
}
*/

// 擴展成各類型通用
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

type Name string

func (i *Name) String() string {
	return fmt.Sprint(*i)
}

func (i *Name) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("name flag already set")
	}

	*i = Name("eddycjy:" + value)
	return nil
}

func main() {
	var name Name
	flag.Var(&name, "name", "說明資訊")
	flag.Parse()
	log.Printf("name: %s", name)
}
