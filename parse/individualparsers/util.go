package individualparsers

import (
	"io"
	"math"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func entropy(r io.Reader) (float64, error) {
	var d [256]uint64
	var err error

	p := make([]byte, 1024)
	for {
		var n int
		n, err = r.Read(p)
		for i := 0; i < n; i++ {
			d[p[i]] += 1
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return 0.0, err
		}
	}

	var sum uint64
	for _, count := range d {
		sum += count
	}

	sumf := float64(sum)
	if sumf == 0.0 {
		return 0.0, nil
	}

	var shannonEntropy float64
	for _, count := range d {
		pct := float64(count) / sumf
		if pct != 0.0 {
			shannonEntropy += -pct * math.Log2(pct)
		}
	}

	if shannonEntropy > 8.0 {
		shannonEntropy = 8.0
	} else if shannonEntropy < 0.0 {
		shannonEntropy = 0.0
	}

	return shannonEntropy, nil
}
