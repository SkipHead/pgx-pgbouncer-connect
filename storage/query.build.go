package storage

import (
	"fmt"
	"strconv"
	"strings"
)

// SelectAllColumns - Return string "SELECT * FROM <table>".
func (q *Query) SelectAllColumns() string {

	return fmt.Sprintf("SELECT %s FROM %s", strings.Join(q.Columns, ","), q.Table)
}

// SelectWhereParam - Return string "SELECT * FROM <table> WHERE <column=param>".
func (q *Query) SelectWhereParam(param string) string {

	return fmt.Sprintf("SELECT %s FROM %s WHERE %s=$1", strings.Join(q.Columns, ","), q.Table, param)
}

// Insert - Return string "INSERT INTO <table> (column1, column2 ...) VALUES ($1, $2 ...)"
func (q *Query) Insert() string {
	var values []string
	if len(q.Columns) > 0 {
		for i := 0; i < len(q.Columns); i++ {
			values = append(values, "$"+strconv.Itoa(i+1))
		}
	}

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		q.Table, strings.Join(q.Columns, ","), strings.Join(values, ","))

	return sql
}

func (q *Query) setUpdate() string {
	var values []string
	if len(q.Columns) > 0 {
		for index, column := range q.Columns {
			values = append(values, fmt.Sprintf("%s=$%s", column, strconv.Itoa(index+2)))
		}
	}

	return strings.Join(values, ",")
}

// Update - Return string "ON CONFLICT <id key> DO UPDATE SET <column=$1>, <column=$2> ..."
func (q *Query) Update() string {

	return fmt.Sprintf("UPDATE %s SET %s WHERE %s", q.Table, q.setUpdate(), q.KeyField)
}

// OnConflictDoUpdate - Return string "ON CONFLICT <id key> DO UPDATE SET <column=$1>, <column=$2> ..."
func (q *Query) OnConflictDoUpdate() string {

	return fmt.Sprintf("ON CONFLICT (%s) DO UPDATE SET %s", q.KeyField, q.setUpdate())
}

// Delete - Return string DELETE FROM <table> WHERE <id>
func (q *Query) Delete() string {

	return fmt.Sprintf("DELETE FROM %s WHERE %s", q.Table, q.KeyField)
}
