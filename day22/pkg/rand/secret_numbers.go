package rand

const pRUNE = 16777216

var cache map[int]int = make(map[int]int)

func next(n int) (out int) {
	out = prune(mix(n<<6, n))
	out = prune(mix(out>>5, out))
	out = prune(mix(out<<11, out))
	return
}

func mix(in, secret int) int {
	return in ^ secret
}

func prune(n int) int {
	return n % pRUNE
}

func NextN(seed, simSteps int) (random int) {
	random = seed
	for i := 0; i < simSteps; i++ {
		random = next(random)
	}
	return
}

func allPrices(seed, simSteps int) []price {
	random := seed
	retval := make([]price, simSteps+1)
  retval[0] = price{seed, 0}
	for i := 0; i < simSteps; i++ {
		random = next(random)
		cost := random % 10
		newPrice := price{cost, cost - retval[i].cost}
		retval[i+1] = newPrice
	}

	return retval
}

func MonkeyFromSeed(seeds []int, simSteps int) Monkey {
	retval := make(map[sequence]int)
	for _, seed := range seeds {
		prices := allPrices(seed, simSteps)
    seen := make(map[sequence]struct{})
		for i := 4; i <= simSteps; i++ {
			seq := sequence{prices[i-3].change, prices[i-2].change, prices[i-1].change, prices[i].change}
      if _, ok := seen[seq]; !ok {
        seen[seq] = struct{}{}
        retval[seq] += prices[i].cost
      }
		}
	}
	return retval
}

type sequence struct {
	first, seconds, third, fourth int
}

type price struct {
	cost, change int
}

type Monkey map[sequence]int


