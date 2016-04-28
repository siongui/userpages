package wrappali

import "regexp"

var paliWord = regexp.MustCompile(`[AaBbCcDdEeGgHhIiJjKkLlMmNnOoPpRrSsTtUuVvYyĀāĪīŪūṀṁṂṃŊŋṆṇṄṅÑñṬṭḌḍḶḷ]+`)

func WrapPaliWordsInSpan(html string) string {
	return paliWord.ReplaceAllStringFunc(html, func(match string) string {
		return "<span>" + match + "</span>"
	})
}
