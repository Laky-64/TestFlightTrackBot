package testflight

import "regexp"

var (
	RegexAppName = regexp.MustCompile(
		`<title>Join the (.+) beta - TestFlight - Apple</title>`,
	)
	RegexStatus = regexp.MustCompile(
		`<div class="beta-status">\n?.*?\n?\s*<span>(.+)</span>\n?\s*</div>`,
	)
	RegexLink = regexp.MustCompile(
		`^https?://testflight\.apple\.com/join/[A-Za-z0-9]+$`,
	)
)
