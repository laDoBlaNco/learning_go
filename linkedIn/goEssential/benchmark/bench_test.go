package nlp

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var benchText = "Don't communicate by sharing memory, share memory by communicating."

func BenchmarkTokenize(b *testing.B){ // *testing.B instead of *testing.T
	for i:=0;i<b.N;i++{
		tokens:=Tokenize(benchText)
		require.Equal(b,10,len(tokens))
	}
}

