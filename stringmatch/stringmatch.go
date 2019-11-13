package stringmatch

const MaxChar = 256
const Size = 256

func Naive(master, temp string) (result bool, pos int) {
	masterLength := len(master)
	tempLength := len(temp)
	for i := 0; i < masterLength; i++ {
		pos := 0
		j := i
		for ; temp[pos] == master[j] && pos < tempLength; pos++ {
			j++
			if pos == tempLength-1 {
				return true, i
			}
		}
	}
	return false, -1
}

func getNext(temp string) []int {
	var k int
	res := []int{0, 0}
	tempLength := len(temp)
	for i := 2; i <= tempLength; i++ {
		k = 0
		for ; k > 0 && temp[k+1] != temp[i-1]; k = res[k] {
		}
		if temp[k+1] == temp[i-1] {
			k++
		}
		res = append(res, k)
	}
	return res
}

func Kmp(master, temp string) (result bool, pos int) {
	next := getNext(temp)
	tempLength := len(temp)
	masterLength := len(master)
	q := 0
	for i := 1; i <= masterLength; i++ {
		for ; q > 0 && temp[q] != master[i-1]; q = next[q] {
		}
		if temp[q] == master[i-1] {
			q++
		}
		if q == tempLength {
			return true, i - tempLength
			q = 0
		}
	}
	return false, -1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func preBmBc(temp string) []int {
	var res []int
	tempLength := len(temp)
	for i := 0; i < MaxChar; i++ {
		res = append(res, tempLength)
	}
	for i := 0; i < tempLength-1; i++ {
		res[temp[i]] = tempLength - 1 - i
	}
	return res
}

func getSuffix(temp string) [Size]int {
	var suff [Size]int
	var f int
	tempLength := len(temp)
	suff[tempLength-1] = tempLength
	g := tempLength - 1
	for i := tempLength - 2; i >= 0; i-- {
		if i > g && suff[i+tempLength-1-f] < i-g {
			suff[i] = suff[i+tempLength-1-f]
		} else {
			if i < g {
				g = i
			}
			f = i
			for g >= 0 && temp[g] == temp[g+tempLength-1-f] {
				g--
			}
			suff[i] = f - g
		}
	}
	return suff
}

func preBmGs(temp string) [Size]int {
	var res [Size]int
	tempLength := len(temp)
	suffix := getSuffix(temp)
	for i := 0; i < tempLength; i++ {
		res[i] = tempLength
	}
	j := 0
	for i := tempLength - 1; i >= 0; i-- {
		if suffix[i] == i+1 {
			for ; j < tempLength-1-i; j++ {
				if res[j] == tempLength {
					res[j] = tempLength - 1 - i
				}
			}
		}
	}
	for i := 0; i <= tempLength-2; i++ {
		res[tempLength-1-suffix[i]] = tempLength - 1 - i
	}
	return res
}

func BoyerMoore(master, temp string) (result bool, pos int) {
	masterLength := len(master)
	tempLength := len(temp)
	bmBc := preBmBc(temp)
	bmGs := preBmGs(temp)
	for j := 0; j <= masterLength-tempLength; {
		var i int
		for i = tempLength - 1; i >= 0 && temp[i] == master[i+j]; i-- {
		}
		if i < 0 {
			return true, j
		} else {
			j += max(bmBc[master[i+j]]-tempLength+1+i, bmGs[i])
		}
	}
	return false, -1
}
