package resources

import (
	"io/ioutil"
	"path/filepath"
)

func goldenFile(name string) string {
	data, _ := ioutil.ReadFile(filepath.Join("testdata", name+".golden"))
	return string(data)
}
