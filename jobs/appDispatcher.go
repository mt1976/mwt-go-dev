package jobs

import (
	"log"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

func RunJobDispatch(actionType string) {
	logStart(actionType)
	//log.Println(actionType, "*** START ***")
	//_, dispatchList, err := application.GetDispatchStoreList()
	//log.Println("Dispatchers", len(dispatchList), "Errors", err)
	//if err != nil {
	//	log.Panicln(err)
	//}
	//fa := []string{"Foo", "Bar"}
	//	for _, s := range dispatchList {
	//fmt.Printf("Dispatch [%d/%d] ", i+1, noDispatchItems)
	//	Dispatch(s)
	//}
	application.UpdateSchedule(actionType, globals.Dispatcher, "")
	//log.Println(actionType, "*** DONE ***")
	logEnd(actionType)
}

func RunJobDispatchType(actionType string) {
	//log.Println(actionType, "*** START ***")
	logStart(actionType)

	DispatchByType("RV" + actionType)
	application.UpdateSchedule(actionType, globals.Dispatcher, "")
	logEnd(actionType)
	//log.Println(actionType, "*** DONE ***")
}

func DispatchByType(dispatchType string) {
	log.Printf("Dispatch      : Issue %q -> %q \n", dispatchType, globals.SienaProperties["rates"])

}

//Dispatch poo
func Dispatch(s application.DispatchStoreItem) {
	log.Printf("Dispatch      : Issue %q -> %q \n", s.Type, globals.SienaProperties["rates"])
}
