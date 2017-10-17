package Problem0319

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans int
}{

	{3, 1},

	// 可以有多个 testcase
}

func Test_bulbSwitch(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, bulbSwitch(tc.n), "输入:%v", tc)
	}
}

func Benchmark_bulbSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			bulbSwitch(tc.n)
		}
	}
}
