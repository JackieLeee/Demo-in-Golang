package carbon

import (
	"testing"

	"github.com/golang-module/carbon/v2"
)

func TestCarbon(t *testing.T) {
	t.Log(carbon.Now().ToString())
}
