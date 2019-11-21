package randstr_test

import (
	"fmt"
	"testing"

	"github.com/qingtao/randstr"
)

func ExampleNew() {
	s := randstr.New(6)
	fmt.Println(s)
}

func TestNew(t *testing.T) {
	s := randstr.New(6)
	t.Logf("%s", s)
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randstr.New(6)
	}
}
