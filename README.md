# CircleCrossGame
○✗(マルバツ)ゲーム core for go langage

## 概要
強化学習の勉強の題材に○✗ゲームを選択したが、○✗ゲームから作ることになってしまった。  
人類が○✗ゲームを作るのに割く時間を削減するために作成。

```
$ go run main.go
 | |
 | |
 | |
input:x y
0 0
you placed (0 0)
o| |
 | |
 | |
cp placed (0 1)
o| |
x| |
 | |
input:x y
1 1
you placed (1 1)
o| |
x|o|
 | |
cp placed (2 0)
o| |x
x|o|
 | |
input:x y
2 2
you placed (2 2)
o| |x
x|o|
 | |o
winner: you
```

## 構成
![uml](https://user-images.githubusercontent.com/26806928/124417329-dba6ba00-dd93-11eb-8c23-b596dbec6424.png)
auto generated by https://github.com/kazukousen/gouml
