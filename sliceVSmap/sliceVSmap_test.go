package slicevsmap

import "testing"

func benchmarkOfIndexAccessToMap(n int, b *testing.B) {
	m := make(map[int]int, n)
	idx := n / 2

	for i := 0; i < b.N; i++ {
		_ = m[idx]
	}
}
func BenchmarkOfIndexAccessToMap10(b *testing.B)   { benchmarkOfIndexAccessToMap(10, b) }
func BenchmarkOfIndexAccessToMap30(b *testing.B)   { benchmarkOfIndexAccessToMap(30, b) }
func BenchmarkOfIndexAccessToMap70(b *testing.B)   { benchmarkOfIndexAccessToMap(70, b) }
func BenchmarkOfIndexAccessToMap100(b *testing.B)  { benchmarkOfIndexAccessToMap(100, b) }
func BenchmarkOfIndexAccessToMap300(b *testing.B)  { benchmarkOfIndexAccessToMap(300, b) }
func BenchmarkOfIndexAccessToMap700(b *testing.B)  { benchmarkOfIndexAccessToMap(700, b) }
func BenchmarkOfIndexAccessToMap1000(b *testing.B) { benchmarkOfIndexAccessToMap(1000, b) }
func BenchmarkOfIndexAccessToMap3000(b *testing.B) { benchmarkOfIndexAccessToMap(3000, b) }

func benchmarkOfIndexAccessToSlice(n int, b *testing.B) {
	s := make([]int, n)
	idx := n / 2

	for i := 0; i < b.N; i++ {
		_ = s[idx]
	}
}
func BenchmarkOfIndexAccessToSlice10(b *testing.B)   { benchmarkOfIndexAccessToSlice(10, b) }
func BenchmarkOfIndexAccessToSlice30(b *testing.B)   { benchmarkOfIndexAccessToSlice(30, b) }
func BenchmarkOfIndexAccessToSlice70(b *testing.B)   { benchmarkOfIndexAccessToSlice(70, b) }
func BenchmarkOfIndexAccessToSlice100(b *testing.B)  { benchmarkOfIndexAccessToSlice(100, b) }
func BenchmarkOfIndexAccessToSlice300(b *testing.B)  { benchmarkOfIndexAccessToSlice(300, b) }
func BenchmarkOfIndexAccessToSlice700(b *testing.B)  { benchmarkOfIndexAccessToSlice(700, b) }
func BenchmarkOfIndexAccessToSlice1000(b *testing.B) { benchmarkOfIndexAccessToSlice(1000, b) }
func BenchmarkOfIndexAccessToSlice3000(b *testing.B) { benchmarkOfIndexAccessToSlice(3000, b) }

func benchmarkOfSearchElementOfMap(n int, b *testing.B) {
	m := make(map[int]int, n)
	s := make([]int, n)
	for i := range s {
		s[i] = i
	}

	for i := 0; i < b.N; i++ {
		_ = m[search(i%n, s)]
	}
}
func BenchmarkOfSearchElementOfMap10(b *testing.B)   { benchmarkOfSearchElementOfMap(10, b) }
func BenchmarkOfSearchElementOfMap30(b *testing.B)   { benchmarkOfSearchElementOfMap(30, b) }
func BenchmarkOfSearchElementOfMap70(b *testing.B)   { benchmarkOfSearchElementOfMap(70, b) }
func BenchmarkOfSearchElementOfMap100(b *testing.B)  { benchmarkOfSearchElementOfMap(100, b) }
func BenchmarkOfSearchElementOfMap300(b *testing.B)  { benchmarkOfSearchElementOfMap(300, b) }
func BenchmarkOfSearchElementOfMap700(b *testing.B)  { benchmarkOfSearchElementOfMap(700, b) }
func BenchmarkOfSearchElementOfMap1000(b *testing.B) { benchmarkOfSearchElementOfMap(1000, b) }
func BenchmarkOfSearchElementOfMap3000(b *testing.B) { benchmarkOfSearchElementOfMap(3000, b) }

func benchmarkOfSearchElementOfSlice(n int, b *testing.B) {
	s := make([]int, n)
	for i := range s {
		s[i] = i
	}

	for i := 0; i < b.N; i++ {
		_ = s[search(i%n, s)]
	}
}
func BenchmarkOfSearchElementOfSlice10(b *testing.B)   { benchmarkOfSearchElementOfSlice(10, b) }
func BenchmarkOfSearchElementOfSlice30(b *testing.B)   { benchmarkOfSearchElementOfSlice(30, b) }
func BenchmarkOfSearchElementOfSlice70(b *testing.B)   { benchmarkOfSearchElementOfSlice(70, b) }
func BenchmarkOfSearchElementOfSlice100(b *testing.B)  { benchmarkOfSearchElementOfSlice(100, b) }
func BenchmarkOfSearchElementOfSlice300(b *testing.B)  { benchmarkOfSearchElementOfSlice(300, b) }
func BenchmarkOfSearchElementOfSlice700(b *testing.B)  { benchmarkOfSearchElementOfSlice(700, b) }
func BenchmarkOfSearchElementOfSlice1000(b *testing.B) { benchmarkOfSearchElementOfSlice(1000, b) }
func BenchmarkOfSearchElementOfSlice3000(b *testing.B) { benchmarkOfSearchElementOfSlice(3000, b) }

func benchmarkOfSearch(n int, b *testing.B) {
	s := make([]int, n)
	for i := range s {
		s[i] = i
	}

	for i := 0; i < b.N; i++ {
		_ = search(i%n, s)
	}
}
func BenchmarkOfSearch10(b *testing.B)   { benchmarkOfSearch(10, b) }
func BenchmarkOfSearch30(b *testing.B)   { benchmarkOfSearch(30, b) }
func BenchmarkOfSearch70(b *testing.B)   { benchmarkOfSearch(70, b) }
func BenchmarkOfSearch100(b *testing.B)  { benchmarkOfSearch(100, b) }
func BenchmarkOfSearch300(b *testing.B)  { benchmarkOfSearch(300, b) }
func BenchmarkOfSearch700(b *testing.B)  { benchmarkOfSearch(700, b) }
func BenchmarkOfSearch1000(b *testing.B) { benchmarkOfSearch(1000, b) }
func BenchmarkOfSearch3000(b *testing.B) { benchmarkOfSearch(3000, b) }

func benchmarkOfMap(n int, b *testing.B) {
	m := make(map[int]int, n)

	for i := 0; i < b.N; i++ {
		_ = m[i%n]
	}
}
func BenchmarkOfMap10(b *testing.B)   { benchmarkOfMap(10, b) }
func BenchmarkOfMap30(b *testing.B)   { benchmarkOfMap(30, b) }
func BenchmarkOfMap70(b *testing.B)   { benchmarkOfMap(70, b) }
func BenchmarkOfMap100(b *testing.B)  { benchmarkOfMap(100, b) }
func BenchmarkOfMap300(b *testing.B)  { benchmarkOfMap(300, b) }
func BenchmarkOfMap700(b *testing.B)  { benchmarkOfMap(700, b) }
func BenchmarkOfMap1000(b *testing.B) { benchmarkOfMap(1000, b) }
func BenchmarkOfMap3000(b *testing.B) { benchmarkOfMap(3000, b) }
