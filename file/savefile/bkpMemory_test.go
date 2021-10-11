package savefile_test

import (
	bkpmemory "github.com/tonnytg/std-golang-wiki/file/savefile"
	"testing"
)

func TestBkpMemory(t *testing.T) {
	file := "test_bkp.txt"
	// in message set \n before close for best experience"
	msg := "Hello\n"
	err := bkpmemory.BkpMemory(file, msg)
	if err != nil {
		t.Errorf("can't bkp memory %v", err)
	}
}
