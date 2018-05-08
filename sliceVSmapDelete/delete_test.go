package del

import "testing"

func benchmarkOfDeleteFromSlice(n int, b *testing.B) {
	idx := n / 2

	for i := 0; i < b.N; i++ {
		s := make([]int, n)
		copy(s[idx:], s[idx+1:])
		s[len(s)-1] = 0
		s = s[:len(s)-1]
	}
}
func BenchmarkOfDeleteFromSlice10(b *testing.B)   { benchmarkOfDeleteFromSlice(10, b) }
func BenchmarkOfDeleteFromSlice30(b *testing.B)   { benchmarkOfDeleteFromSlice(30, b) }
func BenchmarkOfDeleteFromSlice70(b *testing.B)   { benchmarkOfDeleteFromSlice(70, b) }
func BenchmarkOfDeleteFromSlice100(b *testing.B)  { benchmarkOfDeleteFromSlice(100, b) }
func BenchmarkOfDeleteFromSlice300(b *testing.B)  { benchmarkOfDeleteFromSlice(300, b) }
func BenchmarkOfDeleteFromSlice700(b *testing.B)  { benchmarkOfDeleteFromSlice(700, b) }
func BenchmarkOfDeleteFromSlice1000(b *testing.B) { benchmarkOfDeleteFromSlice(1000, b) }
func BenchmarkOfDeleteFromSlice3000(b *testing.B) { benchmarkOfDeleteFromSlice(3000, b) }

func benchmarkOfMakeSlice(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]int, n)
	}
}
func BenchmarkOfMakeSlice10(b *testing.B)   { benchmarkOfMakeSlice(10, b) }
func BenchmarkOfMakeSlice30(b *testing.B)   { benchmarkOfMakeSlice(30, b) }
func BenchmarkOfMakeSlice70(b *testing.B)   { benchmarkOfMakeSlice(70, b) }
func BenchmarkOfMakeSlice100(b *testing.B)  { benchmarkOfMakeSlice(100, b) }
func BenchmarkOfMakeSlice300(b *testing.B)  { benchmarkOfMakeSlice(300, b) }
func BenchmarkOfMakeSlice700(b *testing.B)  { benchmarkOfMakeSlice(700, b) }
func BenchmarkOfMakeSlice1000(b *testing.B) { benchmarkOfMakeSlice(1000, b) }
func BenchmarkOfMakeSlice3000(b *testing.B) { benchmarkOfMakeSlice(3000, b) }

func benchmarkOfDeleteFromMap(n int, b *testing.B) {
	idx := n / 2

	for i := 0; i < b.N; i++ {
		m := make(map[int]int, n)
		delete(m, idx)
	}
}
func BenchmarkOfDeleteFromMap10(b *testing.B)   { benchmarkOfDeleteFromMap(10, b) }
func BenchmarkOfDeleteFromMap30(b *testing.B)   { benchmarkOfDeleteFromMap(30, b) }
func BenchmarkOfDeleteFromMap70(b *testing.B)   { benchmarkOfDeleteFromMap(70, b) }
func BenchmarkOfDeleteFromMap100(b *testing.B)  { benchmarkOfDeleteFromMap(100, b) }
func BenchmarkOfDeleteFromMap300(b *testing.B)  { benchmarkOfDeleteFromMap(300, b) }
func BenchmarkOfDeleteFromMap700(b *testing.B)  { benchmarkOfDeleteFromMap(700, b) }
func BenchmarkOfDeleteFromMap1000(b *testing.B) { benchmarkOfDeleteFromMap(1000, b) }
func BenchmarkOfDeleteFromMap3000(b *testing.B) { benchmarkOfDeleteFromMap(3000, b) }

func benchmarkOfMakeMap(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(map[int]int, n)
	}
}
func BenchmarkOfMakeMap10(b *testing.B)   { benchmarkOfMakeMap(10, b) }
func BenchmarkOfMakeMap30(b *testing.B)   { benchmarkOfMakeMap(30, b) }
func BenchmarkOfMakeMap70(b *testing.B)   { benchmarkOfMakeMap(70, b) }
func BenchmarkOfMakeMap100(b *testing.B)  { benchmarkOfMakeMap(100, b) }
func BenchmarkOfMakeMap300(b *testing.B)  { benchmarkOfMakeMap(300, b) }
func BenchmarkOfMakeMap700(b *testing.B)  { benchmarkOfMakeMap(700, b) }
func BenchmarkOfMakeMap1000(b *testing.B) { benchmarkOfMakeMap(1000, b) }
func BenchmarkOfMakeMap3000(b *testing.B) { benchmarkOfMakeMap(3000, b) }
