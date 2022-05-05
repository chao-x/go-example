package _defer

// defer的触发机制:
// 包裹defer的函数返回时
// 包裹defer的函数执行到末尾时
// 所在的goroutine发生panic时

// defer和return的执行顺序：
// 1.先给返回值赋值
// 2.执行defer语句
// 3.return返回

// AnonymousReturn 匿名返回
// 函数返回值值是匿名，defer的闭包不会影响最终返回
// 因为匿名返回先申请了最终返回值变量，然后对返回值做了copy
// defer操作是在给返回值赋值之后
func AnonymousReturn() int { //匿名返回值
	var r int = 6
	defer func() {
		r *= 7
	}()
	return r
}

// NameClosureReturn 实名闭包返回
// 实名返回值一直用的是一块内存
// 闭包操作会影响实名返回的结果
func NameClosureReturn() (r int) { //
	defer func() {
		r *= 7
	}()
	return 6
}

// NameReturn 实名函数返回
// 虽然实名返回一直用的是一块内存
// 但是函数会对参数做copy, 函数内和返回值不是一块内存
func NameReturn() (r int) {
	defer func(r int) {
		r *= 7
	}(r)
	return 6
}
