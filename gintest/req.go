package gintest

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// MockJSONRequest is a helper to quickly mock a JSON request.
func MockJSONRequest(ctx *gin.Context, data any) {
	buf := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", "/", buf)
	if err != nil {
		panic(err)
	}
	ctx.Request = req
}

// FixHttpTestRecorder must be called if you test gin with an httptest.ResponseRecorder and want to
// assert the responses status code.
//
// This is required to to this bug in gin: https://github.com/gin-gonic/gin/issues/1120
// (Wrong status code returned when using httptest.NewRecorder)
func FixHttpTestRecorder(ctx *gin.Context) {
	ctx.Writer.WriteHeaderNow()
}
