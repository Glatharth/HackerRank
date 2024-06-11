package main

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {
	var high int32
	var sum int32
	for _, candle := range candles {
		if candle >= high {
			if candle != high {
				sum = 0
			}
			high = candle
			sum++
		}
	}
	return sum
}

func main() {
	birthdayCakeCandles([]int32{3, 2, 1, 3})

}
