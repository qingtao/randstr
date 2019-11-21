// Package randstr 随即字符串
package randstr

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	// 随即字符串来源
	encodeStr = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
)

// New 生成一个新的长度为n的随即字符串
func New(n int) string {
	return defaultSource.rand(n)
}

var defaultSource *source

// Source 随即字符串生成器
type source struct {
	r *rand.Rand
	p sync.Pool
}

// RandString 随即字符串
func (s *source) rand(n int) string {
	buf := s.p.Get().(*strings.Builder)
	buf.Reset()
	for i := 0; i < n; i++ {
		buf.WriteByte(encodeStr[rand.Intn(len(encodeStr))])
	}
	str := buf.String()
	buf.Reset()
	s.p.Put(buf)
	return str
}

func init() {
	defaultSource = &source{
		p: sync.Pool{
			New: func() interface{} {
				return new(strings.Builder)
			},
		},
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
