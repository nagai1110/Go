### Webサイトが応答しない場合の振る舞い
上手く調べられなかったのでテストコード上でモックサーバーを作って100秒スリープさせてみた。
(練習問題で期待されていることを再現できているのかが分からない)

応答のあるサーバは順に処理されていき、main()関数はテストサーバーの100秒間を待って処理を完了した。
応答しないWebサーバにサクセスした場合、応答を待ち続けてmain()関数が終了しない状態になる？

```
０.18s   11578  http://google.com
0.27s       81  http://baidu.com
0.36s    12499  http://google.co.in
0.66s   253187  http://qq.com
1.31s   107067  http://facebook.com
1.54s   534983  http://youtube.com
1.67s   421362  http://yahoo.com
1.80s   226911  http://tmall.com
1.95s    81795  http://wikipedia.org
2.09s    51781  http://taobao.com
100.00s      0  http://127.0.0.1:56291
100.00s elapsed
```
