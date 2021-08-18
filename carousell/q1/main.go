package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var max = nums[0]
	var sum = max
	for _, v := range nums[1:] {
		if sum+v > v {
			sum = sum + v
		} else {
			sum = v
		}

		if max < sum {
			max = sum
		}
	}

	if max < 0 {
		return 0
	}

	return max
}

func isValid(text string) (bool, []int) {
	var ret []int
	split := strings.Split(text, ",")
	for _, s := range split {
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return false, nil
		}

		ret = append(ret, int(num))
	}

	return true, ret
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")
	text = strings.ReplaceAll(text, " ", "")
	if ok, nums := isValid(text); ok {
		fmt.Println(maxSum(nums))
	} else {
		fmt.Println(0)
	}
}
