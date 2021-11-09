package datamodel

//appCredentialsStoreItem is cheese
type Credentials struct {
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

const (
	//Action     string
	Credentials_Title       = "User Credentials"
	Credentials_QueryString = "CredentialsStore"
	Credentials_SQLTable    = "credentialsStore"
	Credentials_ID          = "id"
	Credentials_Username    = "username"
	Credentials_Password    = "password"
	Credentials_Firstname   = "firstname"
	Credentials_Lastname    = "lastname"
	Credentials_Knownas     = "knownas"
	Credentials_Email       = "email"
	Credentials_Issued      = "issued"
	Credentials_Expiry      = "expiry"
	Credentials_Role        = "role"
	Credentials_Brand       = "brand"
	Credentials_SYSCreated  = "_created"
	Credentials_SYSWho      = "_who"
	Credentials_SYSHost     = "_host"
	Credentials_SYSUpdated  = "_updated"
)
