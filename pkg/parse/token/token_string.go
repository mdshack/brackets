// Code generated by "stringer --type=Token"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EOF-0]
	_ = x[ILLEGAL-1]
	_ = x[COMMENT-2]
	_ = x[IDENTITY-3]
	_ = x[INTEGER-4]
	_ = x[STRING-5]
	_ = x[BOOLEAN-6]
	_ = x[ADD-7]
	_ = x[SUBTRACT-8]
	_ = x[MULTIPLY-9]
	_ = x[DIVIDE-10]
	_ = x[ASSIGNMENT-11]
	_ = x[LESS_THAN-12]
	_ = x[GREATER_THAN-13]
	_ = x[EQUAL-14]
	_ = x[NOT_EQUAL-15]
	_ = x[LESS_THAN_OR_EQUAL-16]
	_ = x[GREATER_THAN_OR_EQUAL-17]
	_ = x[AND-18]
	_ = x[OR-19]
	_ = x[OPEN-20]
	_ = x[CLOSE-21]
	_ = x[IF-22]
	_ = x[FOR-23]
	_ = x[FUNCTION-24]
	_ = x[RETURN-25]
}

const _Token_name = "EOFILLEGALCOMMENTIDENTITYINTEGERSTRINGBOOLEANADDSUBTRACTMULTIPLYDIVIDEASSIGNMENTLESS_THANGREATER_THANEQUALNOT_EQUALLESS_THAN_OR_EQUALGREATER_THAN_OR_EQUALANDOROPENCLOSEIFFORFUNCTIONRETURN"

var _Token_index = [...]uint8{0, 3, 10, 17, 25, 32, 38, 45, 48, 56, 64, 70, 80, 89, 101, 106, 115, 133, 154, 157, 159, 163, 168, 170, 173, 181, 187}

func (i Token) String() string {
	if i < 0 || i >= Token(len(_Token_index)-1) {
		return "Token(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Token_name[_Token_index[i]:_Token_index[i+1]]
}
