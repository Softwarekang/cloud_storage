package jwt

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token1, _ := GenerateToken("安康", "rz@ak131499")
	token2, _ := GenerateToken("安康", "rz@ak131499")
	fmt.Println(token2)
	assert.Equal(t, token1, token2)
}
