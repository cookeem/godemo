package benchmark

//在根目录运行：go run gin/gin_demo.go
//在当前目录下运行： go test -bench=".*"
import (
	"testing"
	"math/rand"
	"net/http"
)

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		v1 := rand.Float64()
		v2 := rand.Float64()
		if v2 == 0 {
			v2 = rand.Float64()
		}
		Division(v1, v2)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	Benchmark_Division(b)

}

func Benchmark_Http(b *testing.B) {
	b.StopTimer()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:8081/", nil)
		if err != nil {
			panic(err)
		}
		http.DefaultClient.Do(req)
	}
}