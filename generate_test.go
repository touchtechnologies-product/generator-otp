package generate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateDigit(t *testing.T) {
	otp := "1234567"
	prefix := "ABCD"
	idGenerate := Generate(4, "otp", "TOUCH")
	require.Equal(t, len(prefix+"-"+otp), len(idGenerate))
}
