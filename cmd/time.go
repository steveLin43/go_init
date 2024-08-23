package cmd

import (
	"go_init/internal/timer"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

func init() {
	timeCmd.Flags().StringVarP(&calculateTime, "str", "c", "", "請輸入計算時間")
	timeCmd.Flags().StringVarP(&duration, "mode", "d", "", "請輸入經過時間")
}

// 使用標準函式庫的 RFC3339 標準格式
var t = time.Now().Format(time.RFC3339)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "時間格式處理",
	Long:  "時間格式處理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "取得目前時間",
	Long:  "取得目前時間",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("輸出結果: %s, %d", nowTime.Format("2006-01-0215:04:05"), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "計算所需時間",
	Long:  "計算所需時間",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-0215:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ") //strings.Contains?
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-0215:04"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}

		log.Printf("輸出結果: %s, %d", t.Format(layout), t.Unix())
	},
}
