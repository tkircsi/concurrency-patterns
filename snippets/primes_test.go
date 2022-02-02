package snippets

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsPrime(t *testing.T) {
	ret := IsPrime(1)
	require.False(t, ret)
	ret = IsPrime(100)
	require.False(t, ret)
	ret = IsPrime(198)
	require.False(t, ret)

	ret = IsPrime(2)
	require.True(t, ret)
	ret = IsPrime(3)
	require.True(t, ret)
	ret = IsPrime(5)
	require.True(t, ret)
	ret = IsPrime(97)
	require.True(t, ret)

}
