package resources

import (
	"io/ioutil"
	"path/filepath"
)

func goldenFile(name string) string {
	data, _ := ioutil.ReadFile(filepath.Join(name + ".golden"))
	return string(data)
}
