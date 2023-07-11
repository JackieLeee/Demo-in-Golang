package option

import (
	"fmt"
	"testing"
)

func TestOptions(t *testing.T) {
	opt := InitOptions(
		WithSortOption("id", "desc"),
		WithPageOption(10, 1),
	)
	fmt.Printf("options:%#v\n", opt)
}
