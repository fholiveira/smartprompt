package main

import "testing"

func BenchmarkPrompt(b *testing.B) {
	for i := 0; i < 1000; i++ {
		parsePrompt("{GREEN:bold}{user}@{host} {BLUE:bold}{location:vimstyle} {sourcecontrol} {PURPLE:bold}{symbol} {TEXT:reset}", false)
	}
}
