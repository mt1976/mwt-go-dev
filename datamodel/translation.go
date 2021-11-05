package datamodel

//Translation defines a instance of the translationStore
type Translation struct {
	Id          string
	Class       string
	Message     string
	Translation string
	SYSCreated  string
	SYSWho      string
	SYSHost     string
	SYSUpdated  string
}

const (
	Translation_Title          = "Translation"
	Translation_URI            = "TranslationStore"
	Translation_TemplateList   = "TranslationStoreList"
	Translation_TemplateView   = "TranslationStoreView"
	Translation_TemplateEdit   = "TranslationStoreEdit"
	Translation_TemplateSave   = "TranslationStoreSave"
	Translation_TemplateDelete = "TranslationStoreDelete"
	Translation_TemplateNew    = "TranslationStoreNew"
)
