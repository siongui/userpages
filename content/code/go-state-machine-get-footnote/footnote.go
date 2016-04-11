package main

import (
	"github.com/PuerkitoBio/goquery"
	"os"
	"strings"
)

const (
	InFootnote = iota
	NotInFootnote
)

type StateMachine struct {
	State        int
	FootnoteBody string
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		State: NotInFootnote,
	}
}

func (s *StateMachine) ProcessLine(line string) {
	if strings.HasPrefix(line, "Reference") {
		s.State = InFootnote
	}

	if strings.HasPrefix(line, "Update") && s.State == InFootnote {
		s.State = NotInFootnote
	}

	if s.State == InFootnote {
		s.FootnoteBody += line
	}
}

func extractFootnote(f *os.File) string {
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		panic(err)
	}

	sm := NewStateMachine()
	doc.Find("body").Contents().Each(func(_ int, s *goquery.Selection) {
		sm.ProcessLine(s.Text())
	})

	return sm.FootnoteBody
}
