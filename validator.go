package main

/*
	Parentheses validator based on linked list
*/

var parenthesisDict = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

var closingMap = map[rune]bool{
	')': true,
	']': true,
	'}': true,
}

type closingRune struct {
	val  rune
	next *closingRune
}

func (crune *closingRune) check(r rune) (result bool) {
	return r == crune.val
}

func (crune *closingRune) push(r rune) *closingRune {
	return &closingRune{
		val:  r,
		next: crune,
	}
}

func isValid(s string) bool {

	var closeWatcher *closingRune

	for _, symb := range s {
		if nextClosingRune, ok := parenthesisDict[symb]; ok {
			if closeWatcher == nil {
				closeWatcher = &closingRune{
					val: nextClosingRune,
				}
			} else {
				closeWatcher = closeWatcher.push(nextClosingRune)
			}
		} else {
			if _, ok := closingMap[symb]; ok {
				if closeWatcher == nil {
					return false
				}
				if closeWatcher.check(symb) {
					if closeWatcher.next != nil {
						*closeWatcher = *(closeWatcher.next)
					} else {
						closeWatcher = nil
					}
					continue
				} else {
					return false
				}
			}
		}
	}

	return closeWatcher == nil
}
