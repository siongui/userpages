package locale

import "testing"

func TestParseAcceptLanguage(t *testing.T) {
	t.Log("da, en-gb;q=0.8, en;q=0.7")
	lqs := ParseAcceptLanguage("da, en-gb;q=0.8, en;q=0.7")
	if lqs[0].Lang != "da"    { t.Error(lqs[0].Lang) }
	if lqs[0].Q    != 1       { t.Error(lqs[0].Q) }
	if lqs[1].Lang != "en-gb" { t.Error(lqs[1].Lang) }
	if lqs[1].Q    != 0.8     { t.Error(lqs[1].Q) }
	if lqs[2].Lang != "en"    { t.Error(lqs[2].Lang) }
	if lqs[2].Q    != 0.7     { t.Error(lqs[2].Q) }
	t.Log(lqs)

	t.Log("zh, en-us; q=0.8, en; q=0.6")
	lqs2 := ParseAcceptLanguage("zh, en-us; q=0.8, en; q=0.6")
	t.Log(lqs2)

	t.Log("es-mx, es, en")
	lqs3 := ParseAcceptLanguage("es-mx, es, en")
	t.Log(lqs3)

	t.Log("de; q=1.0, en; q=0.5")
	lqs4 := ParseAcceptLanguage("de; q=1.0, en; q=0.5")
	t.Log(lqs4)
}
