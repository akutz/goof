package goof

import (
	"encoding/json"
	"fmt"
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	e := WithField("hello", "world", "introduction error")
	assert.EqualValues(t, "introduction error", fmt.Sprint(e))
	assert.EqualValues(t, "introduction error", fmt.Sprintf("%s", e))
	assert.EqualValues(t, `"introduction error"`, fmt.Sprintf("%q", e))
	assert.EqualValues(t, "`introduction error`", fmt.Sprintf("%#q", e))
	assert.EqualValues(t, "       introduction error", fmt.Sprintf("%25s", e))
	assert.EqualValues(t, "introduction error       ", fmt.Sprintf("%-25s", e))
	assert.EqualValues(t, `     "introduction error"`, fmt.Sprintf("%25q", e))
	assert.EqualValues(t, `"introduction error"     `, fmt.Sprintf("%-25q", e))
	assert.EqualValues(t, "     `introduction error`", fmt.Sprintf("%#25q", e))
	assert.EqualValues(t, "`introduction error`     ", fmt.Sprintf("%-#25q", e))
	assert.EqualValues(t, "`introduction error`     ", fmt.Sprintf("%#-25q", e))

	assertMsgAndString(t, e, false, false, false)
	assertMsgAndString(t, e, true, false, false)
}

func TestError(t *testing.T) {
	e := WithField("hello", "world", "introduction error")
	assertMsgAndString(t, e, false, false, false)
	assertMsgAndString(t, e, false, true, false)
}

func TestString(t *testing.T) {
	e := WithField("hello", "world", "introduction error")
	assertMsgAndString(t, e, false, false, false)
	assertMsgAndString(t, e, false, false, true)
}

func assertMsgAndString(t *testing.T, e Goof, incErr, incFmt, incStr bool) {
	e.IncludeFieldsInError(incErr)
	e.IncludeFieldsInFormat(incFmt)
	e.IncludeFieldsInString(incStr)
	assertMsgAndStringActual(t, e.Error(), incErr)
	assertMsgAndStringActual(t, e.String(), incStr)
	assertMsgAndStringActual(t, fmt.Sprintf("%s", e), incFmt)
}

func assertMsgAndStringActual(t *testing.T, actual string, inc bool) {
	if inc {
		assert.EqualValues(t, `msg="introduction error" hello=world`, actual)
	} else {
		assert.EqualValues(t, "introduction error", actual)
	}
}

func TestMarshalToJSONSansMessage(t *testing.T) {
	e := WithFields(map[string]interface{}{
		"resourceID": 123,
	}, "invalid resource ID")
	buf, err := json.Marshal(e)
	assert.NoError(t, err)
	t.Log(string(buf))
}

func TestMarshalIndentToJSONSansMessage(t *testing.T) {
	e := WithFields(map[string]interface{}{
		"resourceID": 123,
	}, "invalid resource ID")
	buf, err := json.MarshalIndent(e, "", "  ")
	assert.NoError(t, err)
	t.Log(string(buf))
}

func TestMarshalToJSONWithMessage(t *testing.T) {
	e := WithFields(map[string]interface{}{
		"resourceID": 123,
	}, "invalid resource ID")
	e.IncludeMessageInJSON(true)
	buf, err := json.Marshal(e)
	assert.NoError(t, err)
	t.Log(string(buf))
}

func TestMarshalIndentToJSONWithMessage(t *testing.T) {
	e := WithFields(map[string]interface{}{
		"resourceID": 123,
	}, "invalid resource ID")
	e.IncludeMessageInJSON(true)
	buf, err := json.MarshalIndent(e, "", "  ")
	assert.NoError(t, err)
	t.Log(string(buf))
}

func TestDecode(t *testing.T) {
	decErr, err := UnmarshalJSON(jsonError)
	assert.NoError(t, err)
	assert.Error(t, decErr)

	log.Error(WithError("this failed", New("because of this")))
	log.Error(decErr)
	t.Logf("%v", decErr)
	decErr.IncludeMessageInJSON(true)
	buf, err := decErr.MarshalJSON()
	assert.NoError(t, err)
	t.Log(string(buf))
}

var jsonError = []byte(`{
    "error": {
        "inner": {
            "inner": "bzzzzT! BROKEN",
            "size": 10240
        },
        "iops": 1000
    },
    "message": "IOPS required",
    "status": 500
}`)
