// Package types @author: Violet-Eva @date  : 2025/6/10 @notes :
package types

type RespBody[t T] struct {
	Data struct{
		Rows []t `json:"rows"`
		Total int `json:"total"`
	} `json:"data"`
	Success bool `json:"success"`
	ErrorCode int `json:"error_code"`
}