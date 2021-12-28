package adaptor

import "github.com/mt1976/mwt-go-dev/logs"

func StaticImportProcessResponse(filename string) error {
	//	logs.Success("StaticImport_Dispatch")
	logs.Success("StaticImport_ProcessResponse:" + filename)
	var err error
	return err
}
