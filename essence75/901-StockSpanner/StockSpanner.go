package _01_StockSpanner

type StockSpanner struct {
	prices []int
	mono   []int
	idx    int
}

func Constructor() StockSpanner {
	mono, prices := make([]int, 0), make([]int, 0)
	return StockSpanner{
		prices: prices,
		mono:   mono,
	}
}

func (this *StockSpanner) Next(price int) int {
	prices, mono := this.prices, this.mono
	if len(mono) == 0 {
		this.mono = append(mono, this.idx)
		this.prices = append(prices, price)
		this.idx += 1
		return this.idx
	} else {
		for len(mono) > 0 && prices[mono[len(mono)-1]] <= price {
			mono = mono[:len(mono)-1]
		}
		var ret int
		if len(mono) > 0 {
			ret = this.idx - mono[len(mono)-1]
		} else {
			ret = this.idx + 1
		}

		this.mono = append(mono, this.idx)
		this.prices = append(prices, price)
		this.idx += 1
		return ret
	}
}
