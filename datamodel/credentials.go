package datamodel

//appCredentialsStoreItem is cheese
type AppCredentialsStoreItem struct {
	//Action     string
	Id         string `sql:"id"`
	Username   string
	Password   string
	Firstname  string
	Lastname   string
	Knownas    string
	Email      string
	Issued     string
	Expiry     string
	Role       string
	Brand      string
	SYSCreated string `sql:"_created"`
	SYSWho     string `sql:"_who"`
	SYSHost    string `sql:"_host"`
	SYSUpdated string `sql:"_updated"`
}
