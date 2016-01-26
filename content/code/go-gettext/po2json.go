package po2json

import "regexp"
import "io/ioutil"
import "encoding/json"

type msgIdStrPairs map[string]string
type localesMsg map[string]msgIdStrPairs

const pattern = `msgid "(.+)"\nmsgstr "(.+)"`

func getPOPath(locale, domain, localedir string) string {
	return localedir + "/" + locale + "/LC_MESSAGES/" + domain + ".po"
}

func extractFromPOFile(popath string) msgIdStrPairs {
	buf, err := ioutil.ReadFile(popath)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(string(buf), -1)

	pairs := msgIdStrPairs{}
	for _, array := range matches {
		pairs[array[1]] = array[2]
	}
	return pairs
}

func PO2JSON(locales []string, domain, localedir string) string {
	// create PO-like json data for i18n
	obj := localesMsg{}
	for _, locale := range locales {
		// English is default language
		if locale == "en_US" {
			continue
		}

		obj[locale] = extractFromPOFile(getPOPath(locale, domain, localedir))
	}

	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	return string(b)
}
