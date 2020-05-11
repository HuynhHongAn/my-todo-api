package common

import (
	"fmt"
	"my-todo-api/schema"
)

//TODO (anhh): change "err" to interface
func Format(status int, err interface{}, data interface{}) schema.Response{
	var resp schema.Response
	resp.Status = status
	resp.Code = 0 //Todo (anhh): change this later
	if err != nil {
		resp.Message = fmt.Sprintf("%v", err)
	}
	resp.Data = data
	return resp
}
