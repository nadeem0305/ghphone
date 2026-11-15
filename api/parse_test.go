package handler

import (
	"encoding/json"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHandler(t *testing.T) {
	rec := httptest.NewRecorder()
	Handler(rec, httptest.NewRequest("GET", "/api/parse?numbers="+url.QueryEscape("0244123456\n\nnope\n"), nil))
	var got []result
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || !got[0].Valid || got[0].Network != "MTN" || got[1].Valid {
		t.Errorf("got %+v", got)
	}
}
