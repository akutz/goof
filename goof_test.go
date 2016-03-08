package goof

import (
	"encoding/json"
	"testing"
)

func TestMarshalToJSONSansMessage(t *testing.T) {
	e := WithFields(map[string]interface{}{
		"resourceID": 123,
	}, "invalid resource ID")
	buf, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(buf))
}

func TestMarshalIndentToJSONSansMessage(t *testing.T) {
	e := WithFields(map[string]interface{}{
		"resourceID": 123,
	}, "invalid resource ID")
	buf, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(buf))
}

func TestMarshalToJSONWithMessage(t *testing.T) {
	e := WithFields(map[string]interface{}{
		"resourceID": 123,
	}, "invalid resource ID")
	e.IncludeMessageInJSON(true)
	buf, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(buf))
}

func TestMarshalIndentToJSONWithMessage(t *testing.T) {
	e := WithFields(map[string]interface{}{
		"resourceID": 123,
	}, "invalid resource ID")
	e.IncludeMessageInJSON(true)
	buf, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(buf))
}
