[playground](https://play.golang.org/p/p7Bj7NA-3cX)
## error
```go
type error interface {
       Error() string
}

// スコープを短くするために基本はこう
if err := f(); err != nil {
	// エラー処理
}
///// errorのコンストラクト
err = errors.New("hoge")
err = fmt.Errorf("%s is no found", name)

// errorInterfaceを定義した
type PathError struct {
       Op   string
       Path string
       Err  error
}
func (e *PathError) Error() string {
       return e.Op + " " + e.Path + ": " + e.Err.Error()
}


```