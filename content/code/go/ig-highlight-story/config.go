package igstory

// Cookies information for accessing stories

// See the following Stack Overflow answer to get the values:
// https://stackoverflow.com/a/44773079
var config = map[string]string{
	"ds_user_id": "",
	"sessionid":  "",
	"csrftoken":  "",
	"User-Agent": "Instagram 10.3.2 (iPhone7,2; iPhone OS 9_3_3; en_US; en-US; scale=2.00; 750x1334) AppleWebKit/420+",
}

// Login Instagram first and then you can obtain *ds_user_id*
// in Chrome Developer Tools.
//
// See
// https://stackoverflow.com/a/44773079
// Or
// https://github.com/siongui/goigstorylink#obtain-cookies
func SetUserId(s string) {
	config["ds_user_id"] = s
}

// Login Instagram first and then you can obtain *sessionid*
// in Chrome Developer Tools.
//
// See
// https://stackoverflow.com/a/44773079
// Or
// https://github.com/siongui/goigstorylink#obtain-cookies
func SetSessionId(s string) {
	config["sessionid"] = s
}

// Login Instagram first and then you can obtain *csrftoken*
// in Chrome Developer Tools.
//
// See
// https://stackoverflow.com/a/44773079
// Or
// https://github.com/siongui/goigstorylink#obtain-cookies
func SetCsrfToken(s string) {
	config["csrftoken"] = s
}
