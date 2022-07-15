package configmanager

import (
	"io/ioutil"

	"github.com/EwenLan/silicon-server/slog"
)

const filePermit = 0644

type baseConfigInterface interface {
	load(string)
	save(string)
}

type baseConfigure struct {
	filePath    string
	fileContent []byte
}

func (b *baseConfigure) init(path string) {
	b.filePath = path
}

func (b *baseConfigure) load() {
	f, err := ioutil.ReadFile(b.filePath)
	if err != nil {
		slog.Errorf("fail to load configuration, file = %s, err = %s", b.filePath, err)
		return
	}
	b.fileContent = f
	slog.Debugf("load configuraiton successfully, file = %s", b.filePath)
}

func (b *baseConfigure) save() {
	err := ioutil.WriteFile(b.filePath, b.fileContent, filePermit)
	if err != nil {
		slog.Errorf("fail to save configuraiton, file = %s, err = %s", b.filePath, err)
	}
}
