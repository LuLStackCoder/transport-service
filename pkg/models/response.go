package models

type DataStruct struct{ Res bool }

type Response struct {
	Error       bool              `json:"error"`
	ErrorText   string            `json:"errorText"`
	Data        *DataStruct       `json:"data"`
	CustomError map[string]string `json:"customError"`
}
