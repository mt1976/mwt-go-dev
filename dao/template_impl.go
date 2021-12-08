package dao

import dm "github.com/mt1976/mwt-go-dev/datamodel"

func template_ExtraField_Extra(recItem dm.Template) string {
	return "Extra"
}

func template_ExtraField2_Extra(recItem dm.Template) string {
	return "Extra"
}

func template_FIELD2_Override(recItem dm.Template) string {
	return dm.Template_FIELD2 + " Override"
}
