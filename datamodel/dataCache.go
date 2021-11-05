package datamodel

//appCacheStoreItem is cheese
type DataCache struct {
	//Action     string
	Id         string `sql:"id"`
	Object     string `sql:"object"`
	Field      string `sql:"field"`
	Value      string `sql:"value"`
	Expiry     string `sql:"expiry"`
	SYSCreated string `sql:"_created"`
	SYSWho     string `sql:"_who"`
	SYSHost    string `sql:"_host"`
	SYSUpdated string `sql:"_updated"`
	Source     string `sql:"source"`
}
