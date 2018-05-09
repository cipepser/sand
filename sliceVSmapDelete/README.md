# sliceとmapのdeleteはどちらが速いのか

[Golangでmapとsliceどちらが速いのか](https://github.com/cipepser/sand/tree/master/sliceVSmap)に引き続いて、要素を削除する場合のベンチマークを取ってみる。

## やり方

### slice

メモリリークしないように[SliceTricks](https://github.com/golang/go/wiki/SliceTricks)にある以下の方法で要素を削除する。

```go
copy(a[i:], a[i+1:])
a[len(a)-1] = nil // or the zero value of T
a = a[:len(a)-1]
```

### map

組み込みの`delete`で削除する

```go
delete(map[typeA]typeB, typeA)
```

## 条件

* sliceは`[]int`型とする
* mapは`map[int]int`型とする
* sliceから削除したい要素のindexは`n/2`を使う
* mapから削除したい要素を`n/2`とする

## 環境

```sh
macOS High Sierra version 10.13.4（17E202）
MacBook Pro (Retina, 13-inch, Early 2015)
2.7 GHz Intel Core i5
16 GB 1867 MHz DDR3
Intel Iris Graphics 6100 1536 MB
```

```sh
❯ go version
go version go1.10.1 darwin/amd64
```

## 実装

map, sliceどちらも削除前の状態を毎回使いまわす必要があるが、どちらも参照型で実装されるので、毎回`make`で作り直すことにする。
`make`自体がmapとsliceで差がある(Appendix参照)ので、それぞれmakeするだけベンチマークも測定し、差分(`make+delete`-`make`)を結果とする。
makeに掛かる時間は要素数に単調増加するので、それぞれ差分の計算は、同じ要素数でmakeした結果から求める。

ベンチマークの実装は以下の通り。

### slice

```go
func benchmarkOfDeleteFromSlice(n int, b *testing.B) {
	idx := n / 2

	for i := 0; i < b.N; i++ {
		s := make([]int, n)
		copy(s[idx:], s[idx+1:])
		s[len(s)-1] = 0
		s = s[:len(s)-1]
	}
}

func benchmarkOfMakeSlice(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]int, n)
	}
}
```

### map

```go
func benchmarkOfDeleteFromMap(n int, b *testing.B) {
	idx := n / 2

	for i := 0; i < b.N; i++ {
		m := make(map[int]int, n)
		delete(m, idx)
	}
}
func benchmarkOfMakeMap(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make(map[int]int, n)
	}
}
```

## 結果

![Fig1](https://github.com/cipepser/sand/blob/master/sliceVSmapDelete/img/fig1.delete.png)

要素数が少ないと`make`に掛かる時間が測定によってばらつきがあり、差分がマイナスになったが、要素数が1000を超えたあたりからmapよりsliceが有利であることがわかる。またsliceは削除に要する時間が要素数に依らず、ある程度一定の値となっている。

## Appendix

ベンチマーク結果の生データは以下の通り。

```sh
BenchmarkOfDeleteFromSlice10-4     	30000000	        40.0 ns/op
BenchmarkOfDeleteFromSlice30-4     	20000000	        66.5 ns/op
BenchmarkOfDeleteFromSlice70-4     	10000000	       134 ns/op
BenchmarkOfDeleteFromSlice100-4    	10000000	       229 ns/op
BenchmarkOfDeleteFromSlice300-4    	 3000000	       554 ns/op
BenchmarkOfDeleteFromSlice700-4    	 1000000	      1034 ns/op
BenchmarkOfDeleteFromSlice1000-4   	 1000000	      1478 ns/op
BenchmarkOfDeleteFromSlice3000-4   	  300000	      3903 ns/op

BenchmarkOfMakeSlice10-4           	50000000	        33.8 ns/op
BenchmarkOfMakeSlice30-4           	20000000	        59.6 ns/op
BenchmarkOfMakeSlice70-4           	10000000	       119 ns/op
BenchmarkOfMakeSlice100-4          	10000000	       173 ns/op
BenchmarkOfMakeSlice300-4          	 3000000	       494 ns/op
BenchmarkOfMakeSlice700-4          	 2000000	       938 ns/op
BenchmarkOfMakeSlice1000-4         	 1000000	      1465 ns/op
BenchmarkOfMakeSlice3000-4         	  300000	      3785 ns/op

BenchmarkOfDeleteFromMap10-4       	20000000	       102 ns/op
BenchmarkOfDeleteFromMap30-4       	 5000000	       255 ns/op
BenchmarkOfDeleteFromMap70-4       	 3000000	       604 ns/op
BenchmarkOfDeleteFromMap100-4      	 3000000	       583 ns/op
BenchmarkOfDeleteFromMap300-4      	 1000000	      1546 ns/op
BenchmarkOfDeleteFromMap700-4      	  500000	      2992 ns/op
BenchmarkOfDeleteFromMap1000-4     	  200000	      6079 ns/op
BenchmarkOfDeleteFromMap3000-4     	  100000	     12120 ns/op

BenchmarkOfMakeMap10-4             	20000000	        97.8 ns/op
BenchmarkOfMakeMap30-4             	 5000000	       251 ns/op
BenchmarkOfMakeMap70-4             	 3000000	       590 ns/op
BenchmarkOfMakeMap100-4            	 2000000	       580 ns/op
BenchmarkOfMakeMap300-4            	 1000000	      1563 ns/op
BenchmarkOfMakeMap700-4            	  500000	      3100 ns/op
BenchmarkOfMakeMap1000-4           	  200000	      5980 ns/op
BenchmarkOfMakeMap3000-4           	  100000	     11305 ns/op
```

すべてまとめてグラフにすると以下の通り。

![FigA](https://github.com/cipepser/sand/blob/master/sliceVSmapDelete/img/figA.all.png)

`make`のみの所要時間は以下の通り。

![FigB](https://github.com/cipepser/sand/blob/master/sliceVSmapDelete/img/figB.make.png)

一貫して、sliceのほうが`make`に要する時間は少なかった。
sliceとmapの内部実装も以下の通りで、sliceのほうが初期化に掛かる時間が少ないと思われる。

### sliceの内部実装

```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```

### mapの内部実装

```go
// A hash iteration structure.
// If you modify hiter, also change cmd/internal/gc/reflect.go to indicate
// the layout of this structure.
type hiter struct {
	key         unsafe.Pointer // Must be in first position.  Write nil to indicate iteration end (see cmd/internal/gc/range.go).
	value       unsafe.Pointer // Must be in second position (see cmd/internal/gc/range.go).
	t           *maptype
	h           *hmap
	buckets     unsafe.Pointer // bucket ptr at hash_iter initialization time
	bptr        *bmap          // current bucket
	overflow    *[]*bmap       // keeps overflow buckets of hmap.buckets alive
	oldoverflow *[]*bmap       // keeps overflow buckets of hmap.oldbuckets alive
	startBucket uintptr        // bucket iteration started at
	offset      uint8          // intra-bucket offset to start from during iteration (should be big enough to hold bucketCnt-1)
	wrapped     bool           // already wrapped around from end of bucket array to beginning
	B           uint8
	i           uint8
	bucket      uintptr
	checkBucket uintptr
}
```

## References
* [Golangでmapとsliceどちらが速いのか](https://github.com/cipepser/sand/tree/master/sliceVSmap)
* [SliceTricks](https://github.com/golang/go/wiki/SliceTricks)
* [Source file src/runtime/hashmap.go - The Go Programming Language](https://golang.org/src/runtime/hashmap.go)
* [Source file src/runtime/slice.go - The Go Programming Language](https://golang.org/src/runtime/slice.go)