package jobs

import (
	"testing"

	"github.com/mt1976/mwt-go-dev/core"
)

// TestRunJobHeartBeat is a test
func TestStart(t *testing.T) {
	core.Initialise()
	Start()
}
