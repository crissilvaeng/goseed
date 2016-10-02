package universe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	expected := 42
	actual := Answer()

	assert.Equal(t, expected, actual)
}
