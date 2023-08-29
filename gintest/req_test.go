package gintest

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_mockJSONRequest tests the MockJSONRequest helper.
func Test_mockJSONRequest(t *testing.T) {
	ctx := gin.CreateTestContextOnly(httptest.NewRecorder(), gin.New())

	MockJSONRequest(ctx, map[string]any{"foo": "bar"})

	req := ctx.Request
	require.NotNil(t, req)

	body := req.Body
	require.NotNil(t, body)
	assert.NotEmpty(t, body)

	bodyBytes, err := io.ReadAll(body)
	require.NoError(t, err)
	assert.Equal(t, "{\"foo\":\"bar\"}\n", string(bodyBytes))
}
