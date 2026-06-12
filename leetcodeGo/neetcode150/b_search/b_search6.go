package b_search

import (
	"fmt"
)

type timeStamp struct {
	time int
	val  string
}
type TimeMap struct {
	hashmap map[string][]timeStamp
}

func Constructor() TimeMap {
	return TimeMap{
		hashmap: make(map[string][]timeStamp),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.hashmap[key] = append(this.hashmap[key], timeStamp{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	timeMap, ok := this.hashmap[key]
	if !ok {
		return ""
	}
	return bSearchVal(timeMap, timestamp)
}

func bSearchVal(timeMap []timeStamp, timestamp int) string {
	left := 0
	right := len(timeMap)
	res := ""

	for left < right {
		mid := left + (right-left)/2
		if timeMap[mid].time > timestamp {
			right = mid
		} else {
			res = timeMap[mid].val
			left = mid + 1
		}
	}

	return res
}

func main() {
	tm := Constructor()
	tm.Set("k1", "v1", 10)
	tm.Set("k1", "v2", 20)
	tm.Set("k1", "v3", 30)
	fmt.Println(tm.Get("k1", 15))
}
