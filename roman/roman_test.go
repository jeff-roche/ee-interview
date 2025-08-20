package roman_test

import (
	"mywsapp/roman"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIndex(t *testing.T) {
	t.Run("should return the base10 representation of the roman numeral", func(tt *testing.T) {
		r := "XXXVI"

		resp, err := roman.Parse(r)
		assert.Nil(tt, err)
		assert.Equal(tt, resp, 36)
	})

	t.Run("should only accept integer inputs", func(tt *testing.T) {
		r := "I am Greek, not Roman"

		resp, err := roman.Parse(r)
		assert.NotNil(tt, err)
		assert.Equal(tt, resp, "")
	})
}
