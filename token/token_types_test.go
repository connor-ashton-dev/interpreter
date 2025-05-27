package token

import (
	"testing"
)

func TestTokenTypeString(t *testing.T) {
	tests := []struct {
		typeVal  TokenType
		expected string
	}{
		{LEFT_PAREN, "LEFT_PAREN"},
		{RIGHT_PAREN, "RIGHT_PAREN"},
		{LEFT_BRACE, "LEFT_BRACE"},
		{RIGHT_BRACE, "RIGHT_BRACE"},
		{COMMA, "COMMA"},
		{DOT, "DOT"},
		{MINUS, "MINUS"},
		{PLUS, "PLUS"},
		{SEMICOLON, "SEMICOLON"},
		{SLASH, "SLASH"},
		{STAR, "STAR"},
		{BANG, "BANG"},
		{BANG_EQUAL, "BANG_EQUAL"},
		{EQUAL, "EQUAL"},
		{EQUAL_EQUAL, "EQUAL_EQUAL"},
		{GREATER, "GREATER"},
		{GREATER_EQUAL, "GREATER_EQUAL"},
		{LESS, "LESS"},
		{LESS_EQUAL, "LESS_EQUAL"},
		{IDENTIFIER, "IDENTIFIER"},
		{STRING, "STRING"},
		{NUMBER, "NUMBER"},
		{AND, "AND"},
		{CLASS, "CLASS"},
		{ELSE, "ELSE"},
		{FALSE, "FALSE"},
		{FUN, "FUN"},
		{FOR, "FOR"},
		{IF, "IF"},
		{NIL, "NIL"},
		{OR, "OR"},
		{PRINT, "PRINT"},
		{RETURN, "RETURN"},
		{SUPER, "SUPER"},
		{THIS, "THIS"},
		{TRUE, "TRUE"},
		{VAR, "VAR"},
		{WHILE, "WHILE"},
		{EOF, "EOF"},
	}
	for _, test := range tests {
		if got := test.typeVal.String(); got != test.expected {
			t.Errorf("TokenType(%d).String() = %q, want %q", test.typeVal, got, test.expected)
		}
	}
}

func TestKeywordsMap(t *testing.T) {
	cases := []struct {
		word     string
		expected TokenType
		found    bool
	}{
		{"and", AND, true},
		{"class", CLASS, true},
		{"else", ELSE, true},
		{"false", FALSE, true},
		{"for", FOR, true},
		{"fun", FUN, true},
		{"if", IF, true},
		{"nil", NIL, true},
		{"or", OR, true},
		{"print", PRINT, true},
		{"return", RETURN, true},
		{"super", SUPER, true},
		{"this", THIS, true},
		{"true", TRUE, true},
		{"var", VAR, true},
		{"while", WHILE, true},
		{"notakeyword", 0, false},
	}
	for _, c := range cases {
		val, ok := Keywords[c.word]
		if ok != c.found {
			t.Errorf("Keywords[%q] found = %v, want %v", c.word, ok, c.found)
		}
		if ok && val != c.expected {
			t.Errorf("Keywords[%q] = %v, want %v", c.word, val, c.expected)
		}
	}
}
