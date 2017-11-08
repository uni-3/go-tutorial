import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// １回だけよばれる
	f0 := 0
	f1 := 1
	f2 := f0 + f1
	return func() int {
		// for文の回数よばれる。各変数は更新され続ける
		fib := f0
		f2 = f0 + f1
		f0 = f1
		f1 = f2
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

