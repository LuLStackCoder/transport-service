// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD2b7633eDecodeGithubComLuLStackCoderTestServicePkgModels(in *jlexer.Lexer, out *Response) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "error":
			out.Error = bool(in.Bool())
		case "errorText":
			out.ErrorText = string(in.String())
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				if out.Data == nil {
					out.Data = new(DataStruct)
				}
				(*out.Data).UnmarshalEasyJSON(in)
			}
		case "customError":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.CustomError = make(map[string]string)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 string
					v1 = string(in.String())
					(out.CustomError)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComLuLStackCoderTestServicePkgModels(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Error))
	}
	{
		const prefix string = ",\"errorText\":"
		out.RawString(prefix)
		out.String(string(in.ErrorText))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		if in.Data == nil {
			out.RawString("null")
		} else {
			(*in.Data).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"customError\":"
		out.RawString(prefix)
		if in.CustomError == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.CustomError {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				out.String(string(v2Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComLuLStackCoderTestServicePkgModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComLuLStackCoderTestServicePkgModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComLuLStackCoderTestServicePkgModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComLuLStackCoderTestServicePkgModels(l, v)
}
func easyjsonD2b7633eDecodeGithubComLuLStackCoderTestServicePkgModels1(in *jlexer.Lexer, out *Request) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComLuLStackCoderTestServicePkgModels1(out *jwriter.Writer, in Request) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Request) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComLuLStackCoderTestServicePkgModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Request) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComLuLStackCoderTestServicePkgModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Request) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComLuLStackCoderTestServicePkgModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Request) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComLuLStackCoderTestServicePkgModels1(l, v)
}
func easyjsonD2b7633eDecodeGithubComLuLStackCoderTestServicePkgModels2(in *jlexer.Lexer, out *DataStruct) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Res":
			out.Res = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComLuLStackCoderTestServicePkgModels2(out *jwriter.Writer, in DataStruct) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Res\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Res))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DataStruct) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComLuLStackCoderTestServicePkgModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DataStruct) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComLuLStackCoderTestServicePkgModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DataStruct) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComLuLStackCoderTestServicePkgModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DataStruct) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComLuLStackCoderTestServicePkgModels2(l, v)
}
