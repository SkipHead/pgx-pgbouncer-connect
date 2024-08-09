package storage

import (
	"fmt"
	"strconv"
	"strings"
)

// SelectAllColumns will create a raw query sql string "SELECT * FROM <my_table>".
func (o *Orm) SelectAllColumns() string {

	return fmt.Sprintf("SELECT %s FROM %s", strings.Join(o.Columns, ","), o.Table)
}

// SelectWhereParam  will create a raw query sql string "SELECT * FROM <table> WHERE <column=param>".
func (o *Orm) SelectWhereParam(param string) string {

	return fmt.Sprintf("SELECT %s FROM %s WHERE %s=$1", strings.Join(o.Columns, ","), o.Table, param)
}

// Insert  will create a raw query sql string "INSERT INTO <table> (column1, column2 ...) VALUES ($1, $2 ...)".
func (o *Orm) Insert() string {
	var values []string
	if len(o.Columns) > 0 {
		for i := 0; i < len(o.Columns); i++ {
			values = append(values, "$"+strconv.Itoa(i+1))
		}
	}

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		o.Table, strings.Join(o.Columns, ","), strings.Join(values, ","))

	return sql
}

func (o *Orm) setUpdate() string {
	var values []string
	if len(o.Columns) > 0 {
		for index, column := range o.Columns {
			values = append(values, fmt.Sprintf("%s=$%s", column, strconv.Itoa(index+2)))
		}
	}

	return strings.Join(values, ",")
}

// Update will create a raw query sql string "UPDATE SET <column=$1>, <column=$2> ...".
func (o *Orm) Update() string {

	return fmt.Sprintf("UPDATE %s SET %s WHERE %s", o.Table, o.setUpdate(), o.KeyField)
}

// OnConflictDoUpdate will create a raw query sql string "ON CONFLICT <id key> DO UPDATE SET <column=$1>, <column=$2> ..."
func (o *Orm) OnConflictDoUpdate() string {

	return fmt.Sprintf("ON CONFLICT (%s) DO UPDATE SET %s", o.KeyField, o.setUpdate())
}

// Delete will create a raw query sql string DELETE FROM <table> WHERE <id>=KeyField.
func (o *Orm) Delete() string {

	return fmt.Sprintf("DELETE FROM %s WHERE %s", o.Table, o.KeyField)
}
