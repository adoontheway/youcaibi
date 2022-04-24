package main

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	t.SkipNow() //一定要写在第一行
	res := 1
	if res != 1 {
		t.Error("Value not valid")
	}
}

// go 不保证测试会顺序执行
// 要保证需求的话需要用subtests来
func TestSubs(t *testing.T) {
	t.Run("a1", func(t *testing.T) {
		fmt.Println("a1")
	})
	t.Run("a2", func(t *testing.T) {
		fmt.Println("a2")
	})
	t.Run("a3", func(t *testing.T) {
		fmt.Println("a4")
	})
	t.Run("a4", TestPrint) // 子测试一般用小写开头，最佳实践
}

// 这个是测试文件的主体，有这个函数的话会线跑这个，没有的话执行其他测试
func TestMain(m *testing.M) {
	// todo 准备工作：连数据库，启动http服务器之类的
	m.Run() //  不写这个的话就不会执行其他测试
}

// Test以Test开头，Benchmark函数以Benchmark开头，
// benchmark的case一般会跑benchmark.N次，
// 这个次数会根据case的执行时间是否稳定来增加执行次数
// 所以要保证bench case的运行时间的稳定性，不然的话可能会跑不完benchmark
// 不稳定的话会一直增加
// go test -bench=. 只会跑benchmark
func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Sprintf("%04d", n)
	}
}
