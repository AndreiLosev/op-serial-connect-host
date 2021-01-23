package p2p

import (
	"io/ioutil"
	"op-serial-connect-client/errh"
	"path/filepath"
	"strings"
)

// CreateHostTree ...
func CreateHostTree(path, target string, three []string) []string {
	_, err := ioutil.ReadDir(path)
	pathIsFile := func() bool {
		if err != nil {
			if err.Error() == "readdirent: not a directory" {
				return true
			}
		}
		return false
	}()
	if pathIsFile {
		path = filepath.Dir(path)
	}
	result := make([]string, len(three))
	for i, x := range three {
		result[i] = strings.Replace(x, path, "", -1)
		result[i] = filepath.Join(target, result[i])
	}
	return result
}

// ShowFileTree ...
func ShowFileTree(path string) []string {
	fileInfo, err := ioutil.ReadDir(path)
	err = errh.IsFile(err)
	errh.Panic(err)
	if fileInfo == nil {
		return []string{path}
	}
	result := make([]string, 0)
	for _, x := range fileInfo {
		if x.IsDir() {
			result = append(result, ShowFileTree(filepath.Join(path, x.Name()))...)
		} else {
			result = append(result, filepath.Join(path, x.Name()))
		}
	}
	return result
}
