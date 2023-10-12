package convert

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ToStringWithDefault(t *testing.T) {
	t.Run("ValidString", func(t *testing.T) {
		value := ToStringWithDefault("apple", "default")
		require.Equal(t, "apple", value)
	})

	t.Run("NotString", func(t *testing.T) {
		value := ToStringWithDefault(1, "default")
		require.Equal(t, "default", value)
	})

	t.Run("NotString", func(t *testing.T) {
		value := ToStringWithDefault("", "default")
		require.Equal(t, "", value)
	})
}

func Test_ToValueWithDefault(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		value := ToValueWithDefault(RefOf("apple"), "default")
		require.Equal(t, "apple", value)
	})

	t.Run("Int", func(t *testing.T) {
		value := ToValueWithDefault(RefOf(1), 0)
		require.Equal(t, 1, value)
	})

	t.Run("Nil", func(t *testing.T) {
		value := ToValueWithDefault(nil, "default")
		require.Equal(t, "default", value)
	})
}

func Test_RefOf(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		value := RefOf("apple")
		require.Equal(t, "apple", *value)
	})

	t.Run("Int", func(t *testing.T) {
		value := RefOf(1)
		require.Equal(t, 1, *value)
	})
}
