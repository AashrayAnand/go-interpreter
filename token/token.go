package token

type TokenType string

// every token consists of type, and literal value
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers + literals
	IDENT = "IDENT" // add, x, y
	INT   = "INT"   // 0..9

	// operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	BANG   = "!"
	SLASH  = "/"
	TIMES  = "*"
	LT     = "<"
	GT     = ">"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
