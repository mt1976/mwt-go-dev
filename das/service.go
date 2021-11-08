package das

import (
	"database/sql"
	"log"

	"github.com/davecgh/go-spew/spew"
)

func Query(db sql.DB, query string) ([]map[string]interface{}, error) {

	log.Println("Query:", query)
	rows, _ := db.Query(query) // Note: Ignoring errors for brevity
	cols, _ := rows.Columns()

	recs := []map[string]interface{}{}

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		m := make(map[string]interface{})
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
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
		//fmt.Println(m)
		log.Println("Append", m)
		recs = append(recs, m)
	}
	//log.Println("Recs:", recs)
	spew.Dump(recs)
	//log.Println("Query:", m)
	return recs, nil
}

func Poke(DB *sql.DB) error {
	errordb := DB.Ping()
	if errordb != nil {
		log.Println(errordb.Error())
	}
	return errordb
}
