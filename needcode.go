package main

func main() {

}

func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sComp := make(map[string]int)
	tComp := make(map[string]int)

	for i := 0; i < len(s); i++ {
		sLetter := s[i : i+1]
		tLetter := t[i : i+1]

		sComp[sLetter] += 1
		tComp[tLetter] += 1
	}

	for letter, count := range sComp {
		if tComp[letter] == 0 || tComp[letter] != count {
			return false
		}
	}
	return true
}
