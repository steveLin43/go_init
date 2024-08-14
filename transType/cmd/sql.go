// 初始化子指令
var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql 轉換和處理",
	Long:  "sql 轉換和處理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql 轉換",
	Long:  "sql 轉換",
	Run:   func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:	  dbType,
			Host:	  host,
			UserName: username,
			Password: password,
			Charset:  charset, 
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}

		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "請輸入資料庫的帳號")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "請輸入資料庫的密碼")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "", "請輸入資料庫的HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "", "請輸入資料庫的編碼")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "", "請輸入資料庫的實例類型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "請輸入資料庫的名稱")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "請輸入表名稱")
}
