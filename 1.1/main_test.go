package main

import "testing"

var args = []string{"arg1", "arg2", "arg3", "arg4", "arg5"}

func Benchmark_echo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(args)
	}
}

func Benchmark_echo3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(args)
	}
}
