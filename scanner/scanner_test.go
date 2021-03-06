package scanner

import (
	"testing"
)

func TestScan(t *testing.T) {
	var s Scanner
	type expect struct {
		tok Token
		lit string
	}
	testcases := []struct {
		src     []byte
		expects []expect
	}{
		{
			src: []byte(" 123 "),
			expects: []expect{
				{tok: NUMBER, lit: "123"},
				{tok: EOF, lit: ""},
			},
		},
		{
			src: []byte("set!"),
			expects: []expect{
				{tok: IDENT, lit: "set!"},
				{tok: EOF, lit: ""},
			},
		},
		{
			src: []byte("(+ 1 2)"),
			expects: []expect{
				{tok: LPAREN, lit: ""},
				{tok: IDENT, lit: "+"},
				{tok: NUMBER, lit: "1"},
				{tok: NUMBER, lit: "2"},
				{tok: RPAREN, lit: ""},
				{tok: EOF, lit: ""},
			},
		},
		{
			src: []byte("(+ +1 -2 1.11 -1.11)"),
			expects: []expect{
				{tok: LPAREN, lit: ""},
				{tok: IDENT, lit: "+"},
				{tok: NUMBER, lit: "1"},
				{tok: NUMBER, lit: "-2"},
				{tok: NUMBER, lit: "1.11"},
				{tok: NUMBER, lit: "-1.11"},
				{tok: RPAREN, lit: ""},
				{tok: EOF, lit: ""},
			},
		},
		{
			src: []byte("\"test\""),
			expects: []expect{
				{tok: STRING, lit: "test"},
				{tok: EOF, lit: ""},
			},
		},
	}
	for i, tc := range testcases {
		s.Init(tc.src)
		for j, expect := range tc.expects {
			tok, lit, err := s.Scan()
			if err != nil {
				t.Fatalf("case %d-%d: unexpected error %v", i, j, err)
			}
			if tok != expect.tok {
				t.Fatalf("case %d-%d: expected %d, but got %d", i, j, expect.tok, tok)
			}
			if lit != expect.lit {
				t.Fatalf("case %d-%d: expected %s, but got %s", i, j, expect.lit, lit)
			}
		}
	}
}
