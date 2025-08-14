package main

import (
	"strconv"
	"strings"

	"github.com/BWbwchen/MapReduce/worker"
)

type MapReduce struct {
	mapped map[string]int
	index  int
}

var s *MapReduce

// creating a function to the class
func generator(key string) int {
	var keyindex int
	if s.mapped == nil {
		s.mapped = make(map[string]int)
		s.index = 1
	}
	keyindex, ok := s.mapped[key]
	if !ok {
		s.mapped[key] = s.index
		s.index++
	}

	return keyindex
}

func Map(key string, value string, ctx worker.MrContext) {
	split := strings.Fields(value)
	for _, val := range split {
		ctx.EmitIntermediate(val, strconv.Itoa(generator(val)))
	}
}

func Reduce(key string, values []string, ctx worker.MrContext) {
	ctx.Emit(key, strconv.Itoa(s.mapped[key]))
}
