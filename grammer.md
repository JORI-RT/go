[playground](https://play.golang.org/p/p7Bj7NA-3cX)
## function and type
```go
func main(){
       func(){
              fmt.println("hello")
       }() // この()は無名関数を定義してすぐに呼び出し

// empty interface
var f = interface{}
f = 100
f = "" // javaのオブジェクトの代わり

// 型アサーション
var v interface{}
v = 100
n,ok := v.(int)
fmt.Println(n, ok)
s,ok := v.(string)
fmt.Println(s, ok)

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
// panic/recover
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, "sss")
		}
	}()
	panic("ERROR") // 文字列じゃなくてもいい
}


```