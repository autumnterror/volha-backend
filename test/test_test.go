package test

import (
	"fmt"
	"testing"

	"github.com/rs/xid"
)

func TestMain(t *testing.T) {
	for _ = range 50 {
		fmt.Println(xid.New().String())
	}
}
