# go
go習得

## 概念
シングルバイナリ
クロスコンパイル
型がしかりしたスクリプト言語
型階層がない...IFを用いてぽくはできるがフラットになる

## 環境構築
```sh
# GOPATHの設定
export GOPATH=
# PATHにgo/binの追加
export PATH=
# extensionを追加＋cmd,shift,p -> go.installでツールを全部入れる
```

##  goコマンド
```go
go mod init hoge //hogeがモジュールのダウンロードもとURLになる
go run hoge.go
go build hoge.go // 実行バイナリの作成
GOOS=windows GOARCH=386 go build // クロスコンパイルの例
go env GOPATH // GOPATHを取得


```

## link
* [プロジェクト構成](https://github.com/golang-standards/project-layout)
* [grpc](https://grpc.io/docs/languages/go/quickstart/ )
* [go document](https://learn.go.dev/)
* [言語仕様](https://golang.org/ref/spec)
* [play ground](https://play.golang.org/p/NMRXBXIOrs6)
* [DDD for go](https://qiita.com/rema424/items/9ffbdf584b705cae6a19)
* []()