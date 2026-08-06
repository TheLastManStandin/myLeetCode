package yandex

import "sort"

func MinSumAE(n, k int, a, b []int64) int64 {
	m := n - k
	if m == 0 {
		return 0
	}

	d := make([]int64, n)
	for i := 0; i < n; i++ {
		d[i] = b[i] - a[i]
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i] < d[j]
	})

	pref := make([]int64, n+1)
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i] + d[i]
	}

	const INF int64 = 1 << 60
	ans := INF

	for l := 0; l+m <= n; l++ {
		r := l + m - 1
		mid := (l + r) / 2

		left := d[mid]*int64(mid-l) - (pref[mid] - pref[l])
		right := (pref[r+1] - pref[mid+1]) - d[mid]*int64(r-mid)

		cost := left + right
		if cost < ans {
			ans = cost
		}
	}

	return ans
}
