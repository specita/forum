package common

import (
	"testing"
	"fmt"
)

func TestBatchInsert(t *testing.T) {
	row1 := []interface{}{1, "22", "33"}
	row2 := []interface{}{2, "df", "re"}
	rows := [][]interface{}{}
	rows = append(rows, row1)
	rows = append(rows, row2)
	sql,vals := BatchInsert("article", []string{"aa", "bb", "cc"}, rows)

	fmt.Println(sql)
	fmt.Println(vals...)
}
