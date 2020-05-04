package timeutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDurationString(t *testing.T) {
  result, err := ParseDuration("42s")
  assert.NoError(t, err, "it should not return an error")
  assert.Equal(t, time.Second * 42, result)

  result, err = ParseDuration("42h")
  assert.NoError(t, err, "it should not return an error")
  assert.Equal(t, time.Hour * 42, result)

  result, err = ParseDuration("2d")
  assert.NoError(t, err, "it should not return an error")
  assert.Equal(t, time.Hour * 48, result)

  result, err = ParseDuration("d")
  assert.NotNil(t, err)
}
