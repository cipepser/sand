# Golangでmapとsliceどちらが速いのか

[GoのパフォーマンスTipsメモ](http://golang.rdy.jp/2016/07/27/performance-tips/)にインデックスアクセスについて、以下のように述べられている。

> mapのインデックスアクセスはsliceの数十倍遅い。 100件以下の場合バイナリサーチでsliceから目的の値を探すほうが早い。 100要素超えくらいからmapのアクセス速度一定の恩恵が発揮される。

実際にベンチマークを取ってみる。

## 測定したいこと

1. 要素数の増加に従って、sliceとmapでインデックスアクセス速度がどのように変わるか
1. sliceから特定の要素を探すとき、mapと性能分岐点になる要素数はいくつか

### 条件

* sliceは`[]int`型とする
* sliceの中身はsort済みとする(indexをそのまま要素の値とする)
* mapは`map[int]int`型とする

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

## 固定indexへアクセス

要素数`n`に対して、中間の`n/2`のインデックスにアクセスする時間を測定する。

### 実装

```go
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
```

### 結果

```sh
BenchmarkOfIndexAccessToMap10-4         	1000000000	         2.81 ns/op
BenchmarkOfIndexAccessToMap30-4         	1000000000	         2.81 ns/op
BenchmarkOfIndexAccessToMap70-4         	1000000000	         2.81 ns/op
BenchmarkOfIndexAccessToMap100-4        	1000000000	         2.83 ns/op
BenchmarkOfIndexAccessToMap300-4        	1000000000	         2.81 ns/op
BenchmarkOfIndexAccessToMap700-4        	1000000000	         2.81 ns/op
BenchmarkOfIndexAccessToMap1000-4       	1000000000	         2.81 ns/op
BenchmarkOfIndexAccessToMap3000-4       	1000000000	         2.81 ns/op

BenchmarkOfIndexAccessToSlice10-4       	2000000000	         0.33 ns/op
BenchmarkOfIndexAccessToSlice30-4       	2000000000	         0.33 ns/op
BenchmarkOfIndexAccessToSlice70-4       	2000000000	         0.33 ns/op
BenchmarkOfIndexAccessToSlice100-4      	2000000000	         0.33 ns/op
BenchmarkOfIndexAccessToSlice300-4      	2000000000	         0.33 ns/op
BenchmarkOfIndexAccessToSlice700-4      	2000000000	         0.33 ns/op
BenchmarkOfIndexAccessToSlice1000-4     	2000000000	         0.33 ns/op
BenchmarkOfIndexAccessToSlice3000-4     	2000000000	         0.33 ns/op
```

表とグラフにまとめると以下の通り。

|  要素数 | map | slice |
|  ------ | ------ | ------ |
|  10 | 2.81 | 0.33 |
|  30 | 2.81 | 0.33 |
|  70 | 2.81 | 0.33 |
|  100 | 2.83 | 0.33 |
|  300 | 2.81 | 0.33 |
|  700 | 2.81 | 0.33 |
|  1000 | 2.81 | 0.33 |
|  3000 | 2.81 | 0.33 |

![Fig1](https://github.com/cipepser/sand/blob/master/sliceVSmap/img/fig1.png)

単純なインデックスアクセスであれば、sliceとmapどちらも固定の時間しか掛からない（sliceでも探索するわけではないので当然といえば当然）。アクセス速度はsliceのほうが約8.5倍速い。
アクセスしたい要素のインデックスが予めわかっているようなケースや、各要素をloopで順番に処理していくような操作では、sliceを使ったほうがいいと考えられる。

## 異なるindexへアクセス

アクセスしたい要素のインデックスがわからないケースも想定してベンチマークを取ってみる。乱数でインデックスを決定することも考えたが、`n`に対して、`b.N`が十分大きいことが前測定でわかったので、`i%n`がアクセスしたい要素とする。sliceの場合は、要素をBinary Searchで探索する必要があるので、実装し、同時に測定した。Binary Search単体でどれくらい時間を要するのかも確認したかったので、sliceへのインデックスアクセスをしない測定も同時に行っている。

### 実装

```go
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
```

Binary Searchは以下で実装。

```go
package slicevsmap

// search returns an index of element `e` in slice `s`.
// s must be sorted slice to execute binary search.
func search(e int, s []int) int {
	idx := binarySearch(e, 0, len(s)-1, s)
	if idx < 0 {
		panic("not found")
	}
	return idx
}

func binarySearch(e, l, r int, s []int) int {
	i := (l + r) / 2

	if i == l {
		switch {
		case s[l] == e:
			return l
		case s[r] == e:
			return r
		default:
			return -1
		}
	}

	switch {
	case s[i] < e:
		return binarySearch(e, i+1, r, s)
	case s[i] > e:
		return binarySearch(e, l, i-1, s)
	case s[i] == i:
		return i
	}

	return -1
}
```

### 結果

```sh
BenchmarkOfMap10-4                      	100000000	        14.9 ns/op
BenchmarkOfMap30-4                      	100000000	        14.8 ns/op
BenchmarkOfMap70-4                      	100000000	        14.7 ns/op
BenchmarkOfMap100-4                     	100000000	        14.8 ns/op
BenchmarkOfMap300-4                     	100000000	        14.6 ns/op
BenchmarkOfMap700-4                     	100000000	        14.7 ns/op
BenchmarkOfMap1000-4                    	100000000	        14.6 ns/op
BenchmarkOfMap3000-4                    	100000000	        14.6 ns/op

BenchmarkOfSearchElementOfSlice10-4     	100000000	        23.8 ns/op
BenchmarkOfSearchElementOfSlice30-4     	50000000	        29.2 ns/op
BenchmarkOfSearchElementOfSlice70-4     	50000000	        33.4 ns/op
BenchmarkOfSearchElementOfSlice100-4    	50000000	        34.9 ns/op
BenchmarkOfSearchElementOfSlice300-4    	30000000	        48.7 ns/op
BenchmarkOfSearchElementOfSlice700-4    	20000000	        70.6 ns/op
BenchmarkOfSearchElementOfSlice1000-4   	20000000	        77.6 ns/op
BenchmarkOfSearchElementOfSlice3000-4   	20000000	        89.5 ns/op

BenchmarkOfSearch10-4                   	100000000	        23.1 ns/op
BenchmarkOfSearch30-4                   	50000000	        29.0 ns/op
BenchmarkOfSearch70-4                   	50000000	        32.9 ns/op
BenchmarkOfSearch100-4                  	50000000	        34.9 ns/op
BenchmarkOfSearch300-4                  	30000000	        48.2 ns/op
BenchmarkOfSearch700-4                  	20000000	        69.7 ns/op
BenchmarkOfSearch1000-4                 	20000000	        77.1 ns/op
BenchmarkOfSearch3000-4                 	20000000	        88.8 ns/op
```

こちらも表とグラフにまとめると以下の通り。

|  要素数 | slice + search | searchのみ | map |
|  ------ | ------ | ------ | ------ |
|  10 | 23.8 | 23.1 | 14.9 |
|  30 | 29.2 | 29 | 14.8 |
|  70 | 33.4 | 32.9 | 14.7 |
|  100 | 34.9 | 34.9 | 14.8 |
|  300 | 48.7 | 48.2 | 14.6 |
|  700 | 70.6 | 69.7 | 14.7 |
|  1000 | 77.6 | 77.1 | 14.6 |
|  3000 | 89.5 | 88.8 | 14.6 |

![Fig2](https://github.com/cipepser/sand/blob/master/sliceVSmap/img/fig2.png)

mapは要素数に依存しない定数時間になっている。sliceのほうは、Binary Searchが`O(logN)`で、探索に大部分の時間を費やしてしまっていることがわかる。sort済みのsliceですらこの結果になってしまうので、`Exists`メソッドを生やしたいような場合では、mapを使ったほうがよいと思われる。

### 結論

1. 要素数の増加に従って、sliceとmapでインデックスアクセス速度がどのように変わるか
→ 単純なインデックスアクセスであれば、sliceのほうが8.5倍程度速い

1. sliceから特定の要素を探すとき、mapと性能分岐点になる要素数はいくつか
→ searchも含めて行う場合は、mapのほうが速い

## References
* [GoのパフォーマンスTipsメモ](http://golang.rdy.jp/2016/07/27/performance-tips/)