package el

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	l := New()
	assert.EqualValues(t, DefaultConf(), l.c)
}
