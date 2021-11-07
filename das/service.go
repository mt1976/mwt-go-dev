package das

import (
	"database/sql"
	"fmt"
	"log"
)

func Query(db sql.DB, query string) (map[string]interface{}, error) {

	log.Println("Query:", query)
	rows, _ := db.Query(query) // Note: Ignoring errors for brevity
	cols, _ := rows.Columns()

	m := make(map[string]interface{})

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		//m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		// Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...]
		fmt.Print(m)
	}
	return m, nil
}
