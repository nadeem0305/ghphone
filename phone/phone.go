// Package phone validates Ghanaian mobile numbers and identifies the network.
package phone

import (
	"errors"
	"strings"
)

// shortcut: prefix table as of 2026; ported numbers keep their original prefix, so the network is a best guess.
var networks = map[string]string{
	"24": "MTN", "25": "MTN", "53": "MTN", "54": "MTN", "55": "MTN", "59": "MTN",
	"20": "Telecel", "50": "Telecel",
	"26": "AT", "27": "AT", "56": "AT", "57": "AT",
}

var ErrInvalid = errors.New("not a valid Ghanaian mobile number")

type Number struct {
	E164    string // +233XXXXXXXXX
	Local   string // 0XXXXXXXXX
	Network string
}

// Parse accepts 0XXXXXXXXX, 233XXXXXXXXX or +233XXXXXXXXX, with spaces or dashes.
func Parse(s string) (Number, error) {
	s = strings.NewReplacer(" ", "", "-", "", "(", "", ")", "").Replace(strings.TrimSpace(s))
	var rest string
	switch {
	case strings.HasPrefix(s, "+233"):
		rest = s[4:]
	case strings.HasPrefix(s, "233"):
		rest = s[3:]
	case strings.HasPrefix(s, "0"):
		rest = s[1:]
	default:
		return Number{}, ErrInvalid
	}
	if len(rest) != 9 || strings.Trim(rest, "0123456789") != "" {
		return Number{}, ErrInvalid
	}
	network, ok := networks[rest[:2]]
	if !ok {
		return Number{}, ErrInvalid
	}
	return Number{E164: "+233" + rest, Local: "0" + rest, Network: network}, nil
}
