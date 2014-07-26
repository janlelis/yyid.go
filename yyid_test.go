package yyid

import(
	"testing"
	"regexp"
)

func TestYYID(t *testing.T) {
	res := New()
	match, _ := regexp.MatchString("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-f]{12}$", res)
	if !match {
		t.Errorf("'%s' does not look like a UUID", res)
	}
}
