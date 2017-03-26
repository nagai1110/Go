### ベンチマーク結果
最下位ビット検査方式より、テーブル参照方式の方が性能が良い

```
BenchmarkPopCount-8                     2000000000              0.28 ns/op
BenchmarkPopCountByBitShift-8           20000000                82.3 ns/op
```