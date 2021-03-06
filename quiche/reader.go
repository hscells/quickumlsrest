package quiche

import (
	"bytes"
	"encoding/gob"
	"github.com/hscells/quickumlsrest"
	"io/ioutil"
)

func Load(dir string) (quickumlsrest.Cache, error) {
	gob.Register(quickumlsrest.Cache{})
	b, err := ioutil.ReadFile(dir)
	if err != nil {
		panic(err)
	}
	var c quickumlsrest.Cache
	err = gob.NewDecoder(bytes.NewBuffer(b)).Decode(&c)
	if err != nil {
		panic(err)
	}
	return c, nil
}
