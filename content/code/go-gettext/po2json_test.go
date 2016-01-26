package po2json

import "testing"

func TestAll(t *testing.T) {
	locales := []string{"zh_TW", "vi_VN"}
	const domain = "messages"
	const localedir = "locale"

	t.Log(getPOPath(locales[0], domain, localedir))
	t.Log(getPOPath(locales[1], domain, localedir))

	t.Log(PO2JSON(locales, domain, localedir))
}
