package api

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestCreateDfmHome(t *testing.T) {
	dfmFoldername := ".dfm"
	tmp, _ := ioutil.TempDir("", "test")
	dfmHome := filepath.Join(tmp, dfmFoldername)

	CreateDfmHome(dfmHome)
	if !Exists(dfmHome) {
		t.Fail()
	}
}
