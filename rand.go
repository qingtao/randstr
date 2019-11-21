// Package randstr 随机字符串生成器
package randstr

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	defaultSource *source
	// 随机字符串来源
	encodeStr = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
)

// New 生成一个长度为n的随机字符串
func New(n int) string {
	return defaultSource.rand(n)
}

type source struct {
	sync.Pool
}

// rand 随机字符串
func (s *source) rand(n int) string {
	buf := s.Get().(*strings.Builder)
	buf.Reset()
	for i := 0; i < n; i++ {
		buf.WriteByte(encodeStr[rand.Intn(len(encodeStr))])
	}
	str := buf.String()
	s.Put(buf)
	return str
}

func init() {
	defaultSource = &source{
		Pool: sync.Pool{
			New: func() interface{} {
				return new(strings.Builder)
			},
		},
	}
	rand.Seed(time.Now().UnixNano())
}
