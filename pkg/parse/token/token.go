package token

type Token int

const (
	EOF Token = iota
	ILLEGAL
	COMMENT
	IDENTITY

	INTEGER // 123
	STRING  // "test"

	// Single character operators
	ADD          // +
	SUBTRACT     // -
	MULTIPLY     // *
	DIVIDE       // /
	ASSIGNMENT   // =
	LESS_THAN    // <
	GREATER_THAN // >

	// Comparison operators
	EQUAL                 // ==
	NOT_EQUAL             // !=
	LESS_THAN_OR_EQUAL    // <=
	GREATER_THAN_OR_EQUAL // >=

	// Logic compounders
	AND // &&
	OR  // ||

	// General open/closers
	OPEN  // {
	CLOSE // }

	// Keywords
	IF       // if
	FOR      // for
	FUNCTION // function
	RETURN   // return
)
