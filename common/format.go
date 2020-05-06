package common

import "my-todo-api/schema"

func Format(status int, err error, data interface{}) schema.Response{
	var resp schema.Response
	resp.Status = status
	resp.Code = 0 //Todo (anhh): change this later
	if err != nil {
		resp.Message = err.Error()
	}
	resp.Data = data
	return resp
}
