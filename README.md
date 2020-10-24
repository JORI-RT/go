# go
go習得


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

```

## [プロジェクト構成](https://github.com/golang-standards/project-layout)