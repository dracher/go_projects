package helpers

import (
	"testing"
)

func TestGetRhevhBuildRepoPath(t *testing.T) {

	path := getRhevhBuildRepoPath()
	if path != "/tmp" {
		t.Fatal("rhevh_buld_repo not equal to '/tmp'")
	}
}
