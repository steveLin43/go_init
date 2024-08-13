var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `需要計算的時間，有效單位為時間戳記或已格式化後的時間`)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持續時間，有效時間單位`)
}