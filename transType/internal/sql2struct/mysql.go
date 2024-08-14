type DBModel struct {
	DBEngine *sql.DB
	DBInfo	 *DBInfo
}

type DBInfo struct {
	DBType	 string
	Host	 string
	UserName string
	Password string
	Charset	 string
}

type TableColumn struct {
	ColumnName	  string
	DataType	  string
	IsNullable	  string
	ColumnKey	  string
	CloumnType	  string
	ColumnComment string
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel(DBInfo: info)
}

// 連接 mysql 資料庫
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// 省略
)

func (m &DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?" + "charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s,
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, " +
			"IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
			"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("沒有資料")
	}
	defer rows.Close()

	var column []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey,
			&column.IsNullable, &column.CloumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	return columns, nil
}

// 表的欄位類型轉換，宣告全域變數
var DBTypeToStructType = map[string]string{
	"int":		 "int32",
	"tinyint":	 "int8",
	"smallint":	 "int",
	"mediumint": "int64",
	"bigint":	 "int64",
	"bit":	 	 "int",
	"bool":		 "bool",
	"enum":	 	 "string",
	"set":		 "string",
	"varchar":	 "string",
}
