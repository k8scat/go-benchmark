# Go 性能测试

## Benchmark 基准测试

### benchmark 相关参数说明

```bash
# -bench 使用正则表达式匹配
# -benchtime 指定运行的次数，即 b.N，也可以指定运行的时间，比如 -benchtime=5s
# -cpu=2,4 改变 GOMAXPROCS，-cpu 支持传入一个列表作为参数，比如 -cpu=2,4，表示使用 2 核和 4 核 分别进行测试
# -count 参数可以用来设置 benchmark 的轮数，默认是 1
# -benchmem 参数可以度量内存分配的次数
```

### maps vs switches

```bash
go test -bench="Switch|Map" -benchtime=100000000x .
```

### make with cap

```bash
go test -bench="Generate" -benchmem .
```

### ResetTimer

如果在 benchmark 开始前，需要一些准备工作，如果准备工作比较耗时，则需要将这部分代码的耗时忽略掉。

```go
func BenchmarkFib(b *testing.B) {
    time.Sleep(time.Second * 3) // 模拟耗时准备任务
    b.ResetTimer() // 重置定时器
    for n := 0; n < b.N; n++ {
        fib(30) // run fib(30) b.N times
    }
}
```

### StopTimer and StartTimer

每次函数调用前后需要一些准备工作和清理工作，我们可以使用 StopTimer 暂停计时以及使用 StartTimer 开始计时。

例如，如果测试一个冒泡函数的性能，每次调用冒泡函数前，需要随机生成一个数字序列，这是非常耗时的操作，这种场景下，就需要使用 StopTimer 和 StartTimer 避免将这部分时间计算在内。

```go
package main

import (
    "math/rand"
    "testing"
    "time"
)

func generateWithCap(n int) []int {
    rand.Seed(time.Now().UnixNano())
    nums := make([]int, 0, n)
    for i := 0; i < n; i++ {
        nums = append(nums, rand.Int())
    }
    return nums
}

func bubbleSort(nums []int) {
    for i := 0; i < len(nums); i++ {
        for j := 1; j < len(nums)-i; j++ {
            if nums[j] < nums[j-1] {
                nums[j], nums[j-1] = nums[j-1], nums[j]
            }
        }
    }
}

func BenchmarkBubbleSort(b *testing.B) {
    for n := 0; n < b.N; n++ {
        b.StopTimer()
        nums := generateWithCap(10000)
        b.StartTimer()
        bubbleSort(nums)
    }
}
```

## 性能测试

### CPU

```bash
go run pprof/cpu.go
```

## Memory

```bash
go run pprof/mem.go
```

## 参考文章

- [benchmark 基准测试](https://geektutu.com/post/hpg-benchmark.html)
- [Go maps vs switches](https://adayinthelifeof.nl/2020/08/12/mac2win.html)
