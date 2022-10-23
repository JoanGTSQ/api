package main

import "github.com/gin-gonic/gin"

type Response struct {
	ResponseCode int
	Context      *gin.Context
	Response     map[string]interface{}
}

func (r *Response) SendAnswer() {

	if len(r.Response) == 0 {
		r.Context.AbortWithStatus(r.ResponseCode)
		return
	}
	if r.Response["message"] == nil {
		r.Response["message"] = "Success"
	}
	r.Context.JSON(r.ResponseCode, gin.H{
		"message": r.Response["message"],
		"links":   r.Response["links"],
		"data":    r.Response["data"],
		"meta":    r.Response["meta"],
	})
}
