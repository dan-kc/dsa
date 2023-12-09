// Time: O(mn)
func GroupAnagram(strs []string) [][]string {
	strsTable := make(map[[26]int][]string)
	for _, s := range strs {
		var count [26]int
		for _, rune := range s {
			count[rune-'a']++
		}

		strsTable[count] = append(strsTable[count], s)
	}

	var result [][]string
	for _, group := range strsTable {
		result = append(result, group)
	}

	return result
}

// Time: O(mnlogn)
// func GroupAnagrams2(strs []string) [][]string {
// 	result := make([][]string, 0)
// 	groups := make(map[string][]string, len(strs))
// 	for _, s := range strs {
// 		sorted := sortedString(s)
// 		groups[sorted] = append(groups[sorted], s)
// 	}
// 	for _, group := range groups {
// 		result = append(result, group)
// 	}
// 	return result
// }
//
// func sortedString(s string) string {
// 	runes := []rune(s)
// 	sort.Slice(runes, func(i, j int) bool {
// 		return runes[i] < runes[j]
// 	})
// 	return string(runes)
// }
