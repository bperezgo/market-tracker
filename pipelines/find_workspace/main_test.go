package main

import (
	"reflect"
	"testing"
)

func Test_Compiling_Go_Work_File(t *testing.T) {
	workspaces := `go 1.18

	use (
		./pipelines/tests
		./pkg
		./replicator
		./tracker
	)
	`
	results := findWorkspaces([]byte(workspaces))
	expected := []string{"./pipelines/tests", "./pkg", "./replicator", "./tracker"}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("results: %v and expected: %v are not equal", results, expected)
		return
	}
}
