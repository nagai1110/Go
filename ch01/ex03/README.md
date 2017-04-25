### 実行時間の差
サンプルの`echo`プログラムと`strings.Join`を使った方式とでは約50nsの差がある。

```
BenchmarkEcho-8                  5000000               263 ns/op
BenchmarkEchoEfficiency-8       10000000               212 ns/op
```
