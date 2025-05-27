package scanner

import (
	"reflect"
	"testing"

	"github.com/connor-ashton-dev/crafting_interpreters/token"
)

func TestScanTokens_PrintString(t *testing.T) {
	source := "print \"hello\";"
	s := New(source, nil)
	tokens := s.ScanTokens()

	expected := []token.Token{
		token.New(token.PRINT, "print", nil, 1),
		token.New(token.STRING, "\"hello\"", "hello", 1),
		token.New(token.SEMICOLON, ";", nil, 1),
		token.New(token.EOF, "", nil, 1),
	}

	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("ScanTokens() = %v, want %v", tokens, expected)
	}
}

func TestScanTokens_SingleCharacterTokens(t *testing.T) {
	source := "(){},.-+;/*"
	s := New(source, nil)
	tokens := s.ScanTokens()

	expected := []token.Token{
		token.New(token.LEFT_PAREN, "(", nil, 1),
		token.New(token.RIGHT_PAREN, ")", nil, 1),
		token.New(token.LEFT_BRACE, "{", nil, 1),
		token.New(token.RIGHT_BRACE, "}", nil, 1),
		token.New(token.COMMA, ",", nil, 1),
		token.New(token.DOT, ".", nil, 1),
		token.New(token.MINUS, "-", nil, 1),
		token.New(token.PLUS, "+", nil, 1),
		token.New(token.SEMICOLON, ";", nil, 1),
		token.New(token.SLASH, "/", nil, 1),
		token.New(token.STAR, "*", nil, 1),
		token.New(token.EOF, "", nil, 1),
	}

	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("SingleCharacterTokens = %v, want %v", tokens, expected)
	}
}

func TestScanTokens_Numbers(t *testing.T) {
	source := "123 45.67"
	s := New(source, nil)
	tokens := s.ScanTokens()

	expected := []token.Token{
		token.New(token.NUMBER, "123", 123.0, 1),
		token.New(token.NUMBER, "45.67", 45.67, 1),
		token.New(token.EOF, "", nil, 1),
	}

	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("Numbers = %v, want %v", tokens, expected)
	}
}

func TestScanTokens_IdentifiersAndKeywords(t *testing.T) {
	source := "and class var foo bar123"
	s := New(source, nil)
	tokens := s.ScanTokens()

	expected := []token.Token{
		token.New(token.AND, "and", nil, 1),
		token.New(token.CLASS, "class", nil, 1),
		token.New(token.VAR, "var", nil, 1),
		token.New(token.IDENTIFIER, "foo", nil, 1),
		token.New(token.IDENTIFIER, "bar123", nil, 1),
		token.New(token.EOF, "", nil, 1),
	}

	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("IdentifiersAndKeywords = %v, want %v", tokens, expected)
	}
}

func TestScanTokens_CommentsAndWhitespace(t *testing.T) {
	source := "// this is a comment\nvar x = 42; // another comment\n"
	s := New(source, nil)
	tokens := s.ScanTokens()

	expected := []token.Token{
		token.New(token.VAR, "var", nil, 2),
		token.New(token.IDENTIFIER, "x", nil, 2),
		token.New(token.EQUAL, "=", nil, 2),
		token.New(token.NUMBER, "42", 42.0, 2),
		token.New(token.SEMICOLON, ";", nil, 2),
		token.New(token.EOF, "", nil, 3),
	}

	if !reflect.DeepEqual(tokens, expected) {
		t.Errorf("CommentsAndWhitespace = %v, want %v", tokens, expected)
	}
}

func TestScanTokens_ErrorHandling(t *testing.T) {
	source := "@"
	errors := []struct {
		line    int
		message string
	}{}
	s := New(source, func(line int, message string) {
		errors = append(errors, struct {
			line    int
			message string
		}{line, message})
	})
	s.ScanTokens()

	if len(errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errors))
	} else {
		exp := struct {
			line    int
			message string
		}{1, "Unexpected character."}
		if errors[0] != exp {
			t.Errorf("Error = %+v, want %+v", errors[0], exp)
		}
	}
}
