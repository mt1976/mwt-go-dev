package monitors

// Simulator_SienaDealImporter_Job defines the job properties, name, period etc..
func Start() error {
	// ----------------------------------------------------------------
	// Job Definition
	// ----------------------------------------------------------------
	go Simulator_SienaFundsChecker_Watch()
	go Simulator_SienaDealImporter_Watch()
	go Simulator_SienaStaticDataImporter_Watch()

	return nil
}
