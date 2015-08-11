package front

import (
	"bytes"
	"io/ioutil"
	"testing"
)

var files = []struct {
	delim, file string
}{
	{"+++", "testdata/front/toml.md"},
	{"---", "testdata/front/yml.md"},
}

func TestMatter(t *testing.T) {
	bodyData, err := ioutil.ReadFile("testdata/front/body.md")
	if err != nil {
		t.Error(err)
	}
	m := NewMatter()
	m.Handle("+++", JSONHandler)
	b, err := ioutil.ReadFile("testdata/front/json.md")
	if err != nil {
		t.Error(err)
	}
	front, body, err := m.parse(bytes.NewReader(b))
	if err != nil {
		t.Error(err)
	}
	if body != string(bodyData) {
		t.Errorf("expected %s got %s", string(bodyData), body)
	}
	if _, ok := front["title"]; !ok {
		t.Error("expected front matter to contain title got nil instead")
	}
}
