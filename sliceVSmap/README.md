# Golangでmapとsliceどちらが速いのか

[GoのパフォーマンスTipsメモ](http://golang.rdy.jp/2016/07/27/performance-tips/)にインデックスアクセスについて、以下のように述べられている。

> mapのインデックスアクセスはsliceの数十倍遅い。 100件以下の場合バイナリサーチでsliceから目的の値を探すほうが早い。 100要素超えくらいからmapのアクセス速度一定の恩恵が発揮される。

実際にベンチマークを取ってみる。

## 問題設定



## 実装



## 結果



## References
* [GoのパフォーマンスTipsメモ](http://golang.rdy.jp/2016/07/27/performance-tips/)