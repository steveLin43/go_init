package sql2struct

import (
	"fmt"
	"go_init/internal/word"
	"os"
	"text/template"
)

// 預先定義結構範本
const structTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}} {{ $length := len .Comment}} {{ if gt $length 0 }}//
	{{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLenn := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}
	{{.Type}}{{.Tag}}{{ else }}{{.Name}}{{ end}}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() sstring {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamel": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}
