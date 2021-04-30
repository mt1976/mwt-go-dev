package jobs

import (
	"log"

	application "github.com/mt1976/mwt-go-dev/application"
)

func RunJobDispatch(actionType string) {
	//log.Println(actionType, "*** START ***")
	_, dispatchList, err := application.GetDispatchStoreList()
	//log.Println("Dispatchers", len(dispatchList), "Errors", err)
	if err != nil {
		log.Panicln(err)
	}
	//fa := []string{"Foo", "Bar"}
	for _, s := range dispatchList {
		//fmt.Printf("Dispatch [%d/%d] ", i+1, noDispatchItems)
		Dispatch(s)
	}
	application.UpdateSchedule(actionType, Dispatcher, "")
	//log.Println(actionType, "*** DONE ***")
}

func RunJobDispatchType(actionType string) {
	//log.Println(actionType, "*** START ***")

	DispatchByType("RV" + actionType)
	application.UpdateSchedule(actionType, Dispatcher, "")
	//log.Println(actionType, "*** DONE ***")
}

func DispatchByType(dispatchType string) {
	_, dispatchList, err := application.GetDispatchStoreByTYPE(dispatchType)
	//log.Println("Dispatchers", len(dispatchList), "Errors", err)
	if err != nil {
		log.Panicln(err)
	}
	//fa := []string{"Foo", "Bar"}
	for _, s := range dispatchList {
		//log.Printf("Dispatch [%d/%d] ", i+1, noDispatchItems)
		Dispatch(s)
	}
}

//Dispatch poo
func Dispatch(s application.DispatchStoreItem) {
	log.Printf("Dispatch      : %q %q -> %q \n", s.System, s.Type, s.Path)
}
