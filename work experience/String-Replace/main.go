package main

import (
	"fmt"
	"strings"
)

func main() {
	str := `[4 144 0 183 200 30 163 182 196 81 54 107 241 211 62 57 146 191 6 32 1 5 116 105 109 25 158 67 108 97 115 115 61 50 49 53 53 53 57 49 55 50 57 124 57 50 124 53 48 50 98 46 55 51 50 54 46 100 48 52 97 124 53 50 48 49 55 54 52 52 48 124 48 124 48 124 51 52 52 54 56 48 53 53 50 124 91 80 58 53 53 53 57 49 55 50 57 124 83 58 65 68 83 76 124 67 58 100 80 76 97 97 82 56 112 110 71 74 74 82 104 86 105 50 77 67 114 51 75 99 111 102 117 117 67 116 52 77 79 50 51 52 116 109 100 111 117 112 103 81 95 48 112 98 97 86 81 104 121 87 45 66 83 69 65 87 101 77 102 112 88 72 87 110 124 50 51 101 59 50 52 119 93]`
	mySlice := strings.Replace(str, " ", ",", -1)
	fmt.Println("{", mySlice[1:len(mySlice)-1], "}")
}
