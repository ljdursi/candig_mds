package database

import (
	"fmt"

	strings "strings"

	"github.com/Bio-core/jtree/models"
)

//BuildQuery takes a Query object and returns a string of the query
func BuildQuery(query models.Query) string {
	if len(query.SelectedFields) == 1 && query.SelectedFields[0] == "*" {
		query.SelectedFields = GetColumns(query.SelectedTables)
	}
	fields := printFields(query.SelectedFields)
	tables := printTables(query.SelectedTables)
	queryString := "SELECT " + fields + " FROM " + tables
	if len(query.SelectedCondition) != 0 {
		if len(query.SelectedCondition[0]) != 0 {
			conditions := printConditions(query.SelectedCondition)
			queryString += " WHERE (" + conditions + ")"
		}
	}
	return queryString
}

// Print comma separated selected fields
func printFields(selectedFields []string) string {
	var str = ""
	for i := 0; i < len(selectedFields); i++ {
		str += selectedFields[i] + " AS '" + selectedFields[i] + "', "
	}
	str = str[0 : len(str)-2]
	return str
}

func printTables(selectedTables []string) string {
	var str = ""
	for i := 0; i < len(selectedTables); i++ {
		if i == 0 {
			str += selectedTables[i]

			// } else {
			// 	str += " JOIN " + selectedTables[i] + " ON " + "patients.sample_id = samples.sample_id"
		} else {
			str += " JOIN " + selectedTables[i] + " ON " + "samples.sample_id = experiments.sample_id"
		}

	}
	return str
}

func printConditions(SelectedCondition [][]string) string {
	var str = ""
	for i := 0; i < len(SelectedCondition); i++ {
		str += SelectedCondition[i][0] + " " + SelectedCondition[i][1] + SelectedCondition[i][2] + "\"" + SelectedCondition[i][3] + "\" "
	}

	str = str[4 : len(str)-1]
	return str
}

//GetColumns returns colums based off of table names
func GetColumns(tables []string) []string {
	var columns []string
	for _, tableName := range tables {
		rows, err := DB.Query("Select * from " + tableName + " where 0=1")
		defer rows.Close()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		columnsSet, err := rows.Columns()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		for _, j := range columnsSet {
			columns = append(columns, tableName+"."+j)
		}
	}
	return columns
}

//GetTables gets all of the tables in the db
func GetTables() []string {
	var tables []string
	rows, err := DB.Query("Show Tables")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var tname string
		rows.Scan(&tname)
		tables = append(tables, strings.ToLower(tname))
	}
	return tables
}