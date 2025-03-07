package macro

import (
	"testing"
	"time"
)

/******************************************* Scenario Test *************************************************/
func TestKintex(t *testing.T) {
	Kintexcamping()
}

func TestCatchMacroMac(t *testing.T) {

	time.Sleep(2 * time.Second)

	CatchMacroMac(false)
}
