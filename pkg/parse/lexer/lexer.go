package lexer

import (
	"bufio"
	"io"
	"unicode"

	"github.com/mdshack/brackets/pkg/parse/token"
)

type Position struct {
	line   int
	column int
}

type Lexer struct {
	position Position
	reader   *bufio.Reader
}

func NewFromReader(reader io.Reader) *Lexer {
	return &Lexer{
		position: Position{line: 1, column: 0},
		reader:   bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() (Position, token.Token, string) {
	for {
		r, _, err := l.reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				return l.position, token.EOF, ""
			}

			panic(err)
		}

		l.position.column++

		switch r {
		case '\n':
			l.NextLine()
		case '+':
			return l.position, token.ADD, "+"
		case '-':
			return l.position, token.SUBTRACT, "-"
		case '*':
			return l.position, token.MULTIPLY, "*"
		case '/':
			// Handle COMMENT
			if next := l.Peek(); next == '/' {
				return l.LexComment(r)
			}

			return l.position, token.DIVIDE, "/"
		case '=':
			next := l.Peek()

			// Handle EQUAL
			if next == '=' {
				start := l.position
				l.AcceptPeek()
				return start, token.EQUAL, "=="
			}

			return l.position, token.ASSIGNMENT, "="
		case '<':
			// Handle LESS_THAN_OR_EQUAL
			if next := l.Peek(); next == '/' {
				start := l.position
				l.AcceptPeek()
				return start, token.LESS_THAN_OR_EQUAL, "<="
			}

			return l.position, token.LESS_THAN, "<"
		case '>':
			// Handle GREATER_THAN_OR_EQUAL
			if next := l.Peek(); next == '=' {
				start := l.position
				l.AcceptPeek()
				return start, token.GREATER_THAN_OR_EQUAL, ">="
			}

			return l.position, token.GREATER_THAN, ">"
		case '{':
			return l.position, token.OPEN, "{"
		case '}':
			return l.position, token.CLOSE, "}"
		case '&':
			// Handle AND
			if next := l.Peek(); next == '&' {
				start := l.position
				l.AcceptPeek()
				return start, token.AND, "&&"
			}

			// No clue what this is
			return l.position, token.ILLEGAL, "&"
		case '|':
			// Handle OR
			if next := l.Peek(); next == '|' {
				start := l.position
				l.AcceptPeek()
				return start, token.OR, "||"
			}

			// No clue what this is
			return l.position, token.ILLEGAL, "|"
		case '"':
			return l.LexString(r)
		default:
			if unicode.IsDigit(r) {
				return l.LexInteger(r)
			} else if unicode.IsLetter(r) {
				return l.LexKeywordOrIdentity(r)
			} else if unicode.IsSpace(r) {
				continue
			}
		}
	}
}

// Peek ahead, don't move the reader or lexer position
func (l *Lexer) Peek() rune {
	r, _, _ := l.reader.ReadRune()
	// todo: handle read errors

	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	return r
}

// We peeked, accepted the result
func (l *Lexer) AcceptPeek() {
	l.reader.ReadRune()
	l.position.column++
}

// Move position to next line
func (l *Lexer) NextLine() {
	l.position.line++
	l.position.column = 0
}

func (l *Lexer) LexKeywordOrIdentity(r rune) (Position, token.Token, string) {
	current := string(r)
	start := l.position

	for {
		next := l.Peek()

		if !unicode.IsLetter(next) {
			switch current {
			case "if":
				return start, token.IF, current
			case "for":
				return start, token.FOR, current
			case "function":
				return start, token.FUNCTION, current
			case "return":
				return start, token.RETURN, current
			}

			break
		}

		l.AcceptPeek()
		current += string(next)
	}

	return start, token.IDENTITY, current
}

func (l *Lexer) LexComment(r rune) (Position, token.Token, string) {
	current := string(r)
	start := l.position

	for {
		next := l.Peek()

		if next == '\n' {
			break
		}

		l.AcceptPeek()
		current += string(next)
	}

	return start, token.COMMENT, current
}

func (l *Lexer) LexString(r rune) (Position, token.Token, string) {
	current := string(r)
	start := l.position

	for {
		next := l.Peek()
		l.AcceptPeek()
		current += string(next)

		if next == '"' {
			break
		}
	}

	return start, token.INTEGER, current
}

func (l *Lexer) LexInteger(r rune) (Position, token.Token, string) {
	current := string(r)
	start := l.position

	for {
		next := l.Peek()

		if !unicode.IsDigit(next) {
			break
		}

		l.AcceptPeek()
		current += string(next)
	}

	return start, token.INTEGER, current
}
