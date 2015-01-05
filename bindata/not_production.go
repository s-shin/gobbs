// +build !production

package bindata

import (
	"github.com/s-shin/gobbs/util"
	"io/ioutil"
	"path"
)

func Asset(name string) ([]byte, error) {
	return ioutil.ReadFile(path.Join(util.ProjectDir(), name))
}
