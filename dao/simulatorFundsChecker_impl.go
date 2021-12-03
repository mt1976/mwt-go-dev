package dao

import (
	adaptor "github.com/mt1976/mwt-go-dev/adaptor"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

// getFundsCheckList read all employees
func Simulator_FundsChecker_GetList() (int, []dm.Simulator_FundsChecker_Item, error) {

	res, simFundsCheckList, err := adaptor.Simulator_FundsChecker_GetList()
	return res, simFundsCheckList, err
}

// getFundsCheckList read all employees
func Simulator_FundsChecker_GetByID(id string) (int, dm.Simulator_FundsChecker_Item, error) {

	res, simFundsCheckItem, err := adaptor.Simulator_FundsChecker_GetByID(id)
	return res, simFundsCheckItem, err
}

/*
func NewFundsCheck(r FundsCheckItem) {
	if len(r.Id) == 0 {
		newID := uuid.New().String()
		r.Id = newID
	}
	putFundsCheck(r)
}
*/

func Simulator_FundsChecker_Store(thisID string, balance string, resultCode string) {
	adaptor.Simulator_FundsChecker_Store(thisID, balance, resultCode)
}

func PutFundsCheck(r dm.Simulator_FundsChecker_Item) {
	putFundsCheck(r)
}

func putFundsCheck(r dm.Simulator_FundsChecker_Item) {
	//fmt.Println(credentialStore)

	//	fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	//	deletesql := fmt.Sprintf(simFundsCheckSQLDELETE, core.ApplicationPropertiesDB["schema"], r.Id)
	//	inserttsql := fmt.Sprintf(simFundsCheckSQLINSERT, core.ApplicationPropertiesDB["schema"], simFundsCheckSQL, r.Id, r.Name, r.Staticin, r.Staticout, r.Txnin, r.Txnout, r.Fundscheckin, r.Fundscheckout, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated)

	//	log.Println("DELETE:", deletesql, core.ApplicationDB)
	//	log.Println("INSERT:", inserttsql, core.ApplicationDB)
	/*
		_, err2 := core.ApplicationDB.Exec(deletesql)
		if err2 != nil {
			log.Println(err2.Error())
		}
		//	log.Println(fred2, err2)
		_, err := core.ApplicationDB.Exec(inserttsql)
		//	log.Println(fred, err)
		if err != nil {
			log.Println(err.Error())
		}
	*/
}

func DeleteFundsCheck(id string) {
	err := adaptor.Simulator_FundsChecker_DeleteByID(id)
	if err != nil {
		logs.Error("Cannot Delete "+id, err)
	}
}

// fetchFundsCheckData read all employees
func fetchFundsCheckData(tsql string) (int, []dm.Simulator_FundsChecker_Item, dm.Simulator_FundsChecker_Item, error) {
	//log.Println(tsql)
	var simFundsCheck dm.Simulator_FundsChecker_Item
	var simFundsCheckList []dm.Simulator_FundsChecker_Item

	count := 0

	// Populate Arrays etc.
	simFundsCheck.Id = "id"
	simFundsCheck.Name = "name"
	simFundsCheck.Source = "source"
	// no change below
	simFundsCheckList = append(simFundsCheckList, simFundsCheck)
	//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
	count++
	//log.Println(count, simFundsCheckList, simFundsCheck)
	return count, simFundsCheckList, simFundsCheck, nil
}

// getFundsCheckList read all employees
func Simulator_FundsChecker_DeleteByID(id string) error {
	val := adaptor.Simulator_FundsChecker_DeleteByID(id)
	return val
}
