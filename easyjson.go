package elastic

import (
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// EasyJSONMarshaler is an easyjson-compatible marshaler interface.
//
// https://github.com/mailru/easyjson/blob/master/helpers.go#L14
type EasyJSONMarshaler interface {
	MarshalEasyJSON(w *jwriter.Writer)
}

// EasyJSONUnmarshaler is an easyjson-compatible unmarshaler interface.
//
// https://github.com/mailru/easyjson/blob/master/helpers.go#L19
type EasyJSONUnmarshaler interface {
	UnmarshalEasyJSON(w *jlexer.Lexer)
}
