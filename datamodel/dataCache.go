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

const (
	DataCache_Title          = "Translation"
	DataCache_URI            = "TranslationStore"
	DataCache_TemplateList   = "TranslationStoreList"
	DataCache_TemplateView   = "TranslationStoreView"
	DataCache_TemplateEdit   = "TranslationStoreEdit"
	DataCache_TemplateSave   = "TranslationStoreSave"
	DataCache_TemplateDelete = "TranslationStoreDelete"
	DataCache_TemplateNew    = "TranslationStoreNew"
)
