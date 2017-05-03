package common

import (
	"fmt"
	"strings"
)

//gorm not support batch insert https://github.com/jinzhu/gorm/issues/255
//so need manual, for example:
//row1 := []interface{}{1, "22", "33"}
//	row2 := []interface{}{2, "df", "re"}
//	rows := [][]interface{}{}
//	rows = append(rows, row1)
//	rows = append(rows, row2)
//	BatchInsert("article", []string{"aa", "bb", "cc"}, rows)
//output:INSERT INTO article(aa,bb,cc) VALUES (?,?,?),(?,?,?)
func BatchInsert(tableName string, columns []string, rows [][]interface{}) (string, []interface{}) {
	valueStrings := make([]string, 0, len(rows))
	valueArgs := make([]interface{}, 0, len(rows)*len(rows[0]))

	var template string
	for range columns {
		template += "?,"
	}
	template = "(" + template[0:len(template)-1] + ")"

	for _, post := range rows {
		valueStrings = append(valueStrings, template)
		for _, v := range post {
			valueArgs = append(valueArgs, v)
		}

	}
	sql := fmt.Sprintf("INSERT INTO %s(%s) VALUES %s", tableName, strings.Join(columns, ","), strings.Join(valueStrings, ","))

	return sql, valueArgs
}
