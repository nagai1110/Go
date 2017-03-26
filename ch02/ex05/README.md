### ベンチマーク結果

1が設定されている最下位ビットをクリアする方式より、テーブル参照方式の方が性能が良い

```
BenchmarkPopCount-8                     2000000000              0.28 ns/op
BenchmarkPopCountByLSBClearance-8       100000000               19.8 ns/op
```

練習問題2.4のベンチマーク結果とも比較すると、最下位ビットの検査を64回繰り返す方式が一番が性能が悪い
```
BenchmarkPopCount-8                     2000000000              0.28 ns/op
BenchmarkPopCountByBitShift-8           20000000                82.3 ns/op
```