package jobs

import (
	"testing"

	"github.com/mt1976/mwt-go-dev/globals"
)

// TestRunJobHeartBeat is a test
func TestStart(t *testing.T) {
	globals.Initialise()
	Start()
}
