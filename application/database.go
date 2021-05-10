package application

import (
	"database/sql"
	"log"
	"strings"

	globals "github.com/mt1976/mwt-go-dev/globals"
)

func GenerateSourceViews(DB *sql.DB) {
	errdb := DB.Ping()
	if errdb != nil {
		log.Println(errdb.Error())
	}
	//spew.Dump(DB)
	//	spew.Dump(DB.Stats())
	// Get a full list of all views
	_, requiredViews, _ := GetDataList("/config/database/views", "", "")
	createTemplate, err := ReadDataFile("templateCreate.sql", "/config/database/templates")
	if err != nil {
		log.Println(err.Error())
	}
	dropTemplate, err := ReadDataFile("templateDrop.sql", "/config/database/templates")
	if err != nil {
		log.Println(err.Error())
	}

	schemaTemplate, err := ReadDataFile("templateCreateSchema.sql", "/config/database/templates")
	if err != nil {
		log.Println(err.Error())
	}

	schemaName := globals.SienaPropertiesDB["schema"]
	databaseName := globals.SienaPropertiesDB["database"]
	sourceschema := globals.SienaPropertiesDB["sourceschema"]

	schemaTemplate = ReplaceWildcard(schemaTemplate, "!SQL.SCHEMA", schemaName)
	//	fmt.Println("***************************************")
	//	log.Println(schemaTemplate)
	//	fmt.Println("***************************************")
	_, err4 := DB.Exec(schemaTemplate)
	if err4 != nil {
		log.Println(err.Error())
	}
	for _, view := range requiredViews {
		thisDrop := dropTemplate
		thisCreate := createTemplate

		fileID := view
		viewName := strings.ReplaceAll(fileID, ".sql", "")
		//	fmt.Println("view name:", viewName)

		thisDrop = ReplaceWildcard(thisDrop, "!SQL.DB", databaseName)
		thisDrop = ReplaceWildcard(thisDrop, "!SQL.SCHEMA", schemaName)
		thisDrop = ReplaceWildcard(thisDrop, "!SQL.VIEW", viewName)

		sqlBody, err := ReadDataFile(fileID, "/config/database/views")
		if err != nil {
			log.Println(err.Error())
		}

		sqlBody = extractCreate(sqlBody)

		//log.Println("***", sqlBody, "***")

		thisCreate = ReplaceWildcard(thisCreate, "!SQL.DB", databaseName)
		thisCreate = ReplaceWildcard(thisCreate, "!SQL.SCHEMA", schemaName)
		thisCreate = ReplaceWildcard(thisCreate, "!SQL.VIEW", viewName)
		thisCreate = ReplaceWildcard(thisCreate, "!SQL.SOURCE", sourceschema)

		sqlBody = ReplaceWildcard(sqlBody, "!SQL.DB", databaseName)
		sqlBody = ReplaceWildcard(sqlBody, "!SQL.SCHEMA", schemaName)
		sqlBody = ReplaceWildcard(sqlBody, "!SQL.VIEW", viewName)
		sqlBody = ReplaceWildcard(sqlBody, "!SQL.SOURCE", sourceschema)

		thisCreate = ReplaceWildcard(thisCreate, "!SQL.BODY", sqlBody)
		//fmt.Println("index:", i, fileID, thisDrop, thisCreate)
		//	fmt.Println("***************************************")
		//	fmt.Println(fileID, viewName)
		//	fmt.Println("***************************************")
		//	fmt.Println(thisDrop)
		//fmt.Println("***************************************")

		log.Println("Refresh View  : ", viewName)
		//log.Println(">>>", thisDrop, "<<<")

		_, erra := DB.Exec(thisDrop)
		//spew.Dump(info)
		//	spew.Dump(erra)
		if erra != nil {
			log.Panic("ARSE", err.Error())
		}

		//fmt.Println("***************************************")
		//fmt.Println(info)
		//fmt.Println("***************************************")
		//fmt.Println(sqlBody)
		//fmt.Println("***************************************")
		//fmt.Println(thisCreate)
		//fmt.Println("***************************************")

		//spew.Dump(thisCreate)
		//log.Println(">>>", sqlBody, "<<<")
		_, err2 := DB.Exec(sqlBody)
		//log.Println(info, err2)
		if err2 != nil {
			log.Println("BUGGER", err.Error())
		}
		//log.Println("***************************************")
	}
}

func extractCreate(in string) string {
	//result := in
	findCREATE := "CREATE"
	selectPos := strings.Index(in, findCREATE)
	//log.Println(findCREATE, selectPos)
	newString := in[selectPos:len(in)]
	findGO := "GO"
	goPos := strings.Index(newString, findGO)
	outString := newString[0:goPos]
	//log.Println(outString)
	return outString
}

func PokeDatabase(DB *sql.DB) error {
	errordb := DB.Ping()
	if errordb != nil {

	}
	return errordb
}
