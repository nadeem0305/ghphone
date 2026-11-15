package ghphone

import "testing"

func TestParse(t *testing.T) {
	ok := map[string]string{
		"0244123456":       "MTN",
		"+233 24 412 3456": "MTN",
		"233-50-123-4567":  "Telecel",
		"(027) 123 4567":   "AT",
		"0591234567":       "MTN",
	}
	for in, want := range ok {
		n, err := Parse(in)
		if err != nil || n.Network != want {
			t.Errorf("Parse(%q) = %+v, %v; want %s", in, n, err, want)
		}
	}
	if n, _ := Parse("+233244123456"); n.Local != "0244123456" || n.E164 != "+233244123456" {
		t.Errorf("formats wrong: %+v", n)
	}
	for _, in := range []string{"", "024412345", "02441234567", "0304123456", "+44 7911 123456", "02441a3456"} {
		if _, err := Parse(in); err == nil {
			t.Errorf("Parse(%q) should fail", in)
		}
	}
}
