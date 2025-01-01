// This is the solution of leetcode problem 1768

package main

func mergeAlternately(word1 string, word2 string) string {
	merged := ""
	i := 0

	for ; i < len(word1) && i < len(word2); i++ {
		merged += string(word1[i]) + string(word2[i])
	}

	for ; i < len(word1); i++ {
		merged += string(word1[i])
	}

	for ; i < len(word2); i++ {
		merged += string(word2[i])
	}

	return merged
}
