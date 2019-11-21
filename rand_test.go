package randstr_test

import (
	"fmt"
	"testing"

	"github.com/qingtao/randstr"
)

func ExampleNew() {
	// 生成长度为6的随机字符串
	s := randstr.New(6)
	// output
	fmt.Println(s)
}

func TestNew(t *testing.T) {
	t.Parallel()
	t.Run("4", func(t *testing.T) {
		s := randstr.New(4)
		t.Logf("%s", s)
	})
	t.Run("6", func(t *testing.T) {
		s := randstr.New(6)
		t.Logf("%s", s)
	})
	t.Run("8", func(t *testing.T) {
		s := randstr.New(8)
		t.Logf("%s", s)
	})
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randstr.New(6)
	}
}
