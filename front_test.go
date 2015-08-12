package front

import (
	"bytes"
	"io/ioutil"
	"testing"
)

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
	front, body, err := m.Parse(bytes.NewReader(b))
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

func TestYAMLHandler(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/sample.yml")
	if err != nil {
		t.Fatal(err)
	}
	f, err := YAMLHandler(string(data))
	if err != nil {
		t.Errorf("handling yaml %v", err)
	}
	if _, ok := f["language"]; !ok {
		t.Errorf("expected language got nil instead")
	}
}
