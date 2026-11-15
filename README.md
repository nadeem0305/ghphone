# ghphone

**Live:** https://ghphone.vercel.app

Validate Ghanaian mobile numbers and tell which network they belong to (MTN, Telecel, AT).
Useful before sending a Mobile Money prompt or an SMS, so typos fail early instead of at the provider.

```go
import "github.com/nadeem0305/ghphone/phone"

n, err := phone.Parse("+233 24 412 3456")
// n.E164 == "+233244123456", n.Local == "0244123456", n.Network == "MTN"
```

Accepts `0XXXXXXXXX`, `233XXXXXXXXX` and `+233XXXXXXXXX`, with spaces, dashes or brackets.

## CLI

```
go install github.com/nadeem0305/ghphone/cmd/ghphone@latest
ghphone 0244123456 0501234567
cat numbers.txt | ghphone
```

Exits with status 1 if any number is invalid.

## Caveat

Numbers ported between networks keep their original prefix, so the network is a best guess from the prefix table.
