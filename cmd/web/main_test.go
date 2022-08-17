package main

import "testing"

// test function for the run function's test
func TestRun(t *testing.T){
	_,err := run()
	if err != nil {
		t.Error("failed run()")
	}
}