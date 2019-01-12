package scanner

import (
	"bytes"
	"log"
)

type Scanner struct {
	src     []byte
	ch      byte
	forward int
	offset  int
}

var eofCh byte = 255

func (s *Scanner) next() {
	if s.forward < len(s.src) {
		s.offset = s.forward
		s.ch = s.src[s.forward]
		s.forward += 1
	} else {
		s.offset = len(s.src)
		s.ch = eofCh
	}
}

func (s *Scanner) Init(src []byte) {
	s.src = src
	s.ch = ' '
	s.forward = 0
	s.offset = 0
	s.next()
}

func (s *Scanner) skipSpaces() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' || s.ch == '\r' {
		s.next()
	}
}

func (s *Scanner) scanNumber() (Token, string) {
	off := s.offset
	for s.ch >= '0' && s.ch <= '9' {
		s.next()
	}
	return INT, string(s.src[off:s.offset])
}

var (
	specialInits      = []byte{'!', '$', '%', '&', '*', '+', '-', '.', '/', ':', '<', '=', '>', '?', '@', '^', '_', '~'}
	specialSubseqents = []byte{'+', '-', '.', '@'}
)

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isInitial(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || bytes.LastIndexByte(specialInits, ch) >= 0
}

func (s *Scanner) scanIdentifier() string {
	offs := s.offset
	for isInitial(s.ch) || isDigit(s.ch) || bytes.LastIndexByte(specialSubseqents, s.ch) >= 0 {
		s.next()
	}
	return string(s.src[offs:s.offset])
}

func (s *Scanner) Scan() (tok Token, lit string) {
	s.skipSpaces()
	ch := s.ch
	if isDigit(ch) {
		return s.scanNumber()
	}
	if isInitial(ch) {
		lit = s.scanIdentifier()
		tok = IDENT
		return
	}
	s.next()
	switch ch {
	case eofCh:
		tok = EOF
	case '(':
		tok = LPAREN
	case ')':
		tok = RPAREN
	default:
		log.Fatalf("Unexpected token %c", ch)
	}
	return
}