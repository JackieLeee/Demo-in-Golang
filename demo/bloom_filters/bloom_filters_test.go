package bloom_filters

import (
	"testing"

	"github.com/bits-and-blooms/bloom/v3"
)

func TestBloomFilters(t *testing.T) {
	filter := bloom.NewWithEstimates(1000000, 0.01)
	if filter.Test([]byte("A")) {
		t.Log("first test: exists")
	} else {
		t.Log("first test: not exists")
	}
	filter.Add([]byte("A"))
	if filter.Test([]byte("A")) {
		t.Log("second test: exists")
	} else {
		t.Log("second test: not exists")
	}
}
