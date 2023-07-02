package main

/*
	Parentheses validator based on reversed linked list with
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

type List struct {
	head *closingRune
}

// inserts new closing rune at the head of the list
func (l *List) Insert(r rune) {
	newHead := &closingRune{val: r, next: nil}
	if l.head == nil {
		l.head = newHead
	} else {
		newHead.next = l.head
		l.head = newHead
	}
}

// checks if rune at the head of the list is equal incoming rune
// if so, pops the head and returns true
// otherwise returns false
func (l *List) Check(r rune) bool {
	if l.head == nil {
		return false
	}
	if l.head.val == r {
		if l.head.next != nil {
			l.head = l.head.next
		} else {
			l.head = nil
		}
		return true
	}
	return false
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	l := &List{}

	for _, symb := range s {
		if nextClosingRune, ok := parenthesisDict[symb]; ok {
			l.Insert(nextClosingRune)
		} else {
			if _, ok := closingMap[symb]; ok {
				if !l.Check(symb) {
					return false
				}
			}
		}
	}

	return l.head == nil
}
