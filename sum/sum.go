package sum

func Sum(numbers []int )int {
	sum := 0
	for _,value:= range(numbers){
		sum +=value
	}
	return sum
}

func SumAll(sumnumbers ...[]int) []int {
	// lengthOfNumbers := len(sumnumbers)
	// sums := make([]int,lengthOfNumbers)
	// for i,numbers := range sumnumbers {
	// 	sums[i] = Sum(numbers)
	// }
	// return sums

	var sums []int
	for _,number := range sumnumbers{
		sums =append(sums,Sum(number) )
	}

 return  sums
}

func SumAllTails(sumtails ...[]int) []int{
	var sums []int
	for _,numbers :=range sumtails{
		if len(numbers) == 0 {
			sums =append(sums,0)
		}else{

			tails := numbers[1:]
			sums = append(sums,Sum(tails))
		}

	}
	return sums
}