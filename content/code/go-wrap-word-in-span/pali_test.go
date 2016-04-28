package wrappali

import "testing"

const html = `1.
Manopubbaṅgamā dhammā, manoseṭṭhā manomayā;
Manasā ce paduṭṭhena, bhāsati vā karoti vā;
Tato naṃ dukkhamanveti, cakkaṃva vahato padaṃ.`

func TestWrapPaliWordsInSpan(t *testing.T) {
	t.Log(WrapPaliWordsInSpan(html))
}
