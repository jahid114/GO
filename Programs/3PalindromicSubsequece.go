package main

func countPalindromicSubsequence(s string) int {
	charPositions := make(map[rune][]int)
	resultSet := make(map[string]bool)

	for index, char := range s {
		charPositions[char] = append(charPositions[char], index)
	}

	for key, value := range charPositions {
		startIndex := value[0]
		endIndex := value[len(value)-1]
		char := string(key)
		if endIndex-startIndex > 1 {
			for i := startIndex + 1; i < endIndex; i++ {
				palindromeString := char + string(s[i]) + char
				resultSet[palindromeString] = true
			}
		}
	}

	return len(resultSet)
}