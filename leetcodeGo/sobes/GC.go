package main

import (
	"fmt"
	"runtime"
	"time"
)

type data struct {
	a []byte
}

func getData() *data {
	a := &data{}
	a.a = make([]byte, 300)
	return a
}

func main() {
	t1 := time.NewTicker(100 * time.Microsecond)

	go func() {
		for range t1.C {
			getData()
		}
	}()

	t2 := time.NewTicker(1 * time.Second)
	var m runtime.MemStats
	now := time.Now()
	for curr := range t2.C {
		runtime.ReadMemStats(&m)
		var j uint32 = 0
		if m.NumGC > 0 {
			j = m.NumGC - uint32(1)
		}
		fmt.Printf("GC enabled %v    GC runs %v    Live now %v     Pause total ms %0.2f    time %5.0f sec     last pause %0.2f\n",
			m.EnableGC, m.NumGC, m.Mallocs-m.Frees, float64(m.PauseTotalNs)/1000/1000, curr.Sub(now).Seconds(), float64(m.PauseNs[j])/1000/1000)
	}
}
