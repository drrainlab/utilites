package main

/*
	Parentheses validator based on reversed linked list
*/

type ListNode struct {
	value rune
	next  *ListNode
}

func push(stack **ListNode, value rune) {
	node := &ListNode{
		value: value,
		next:  *stack,
	}
	*stack = node
}

func pop(stack **ListNode) rune {
	if *stack == nil {
		return 0
	}
	value := (*stack).value
	*stack = (*stack).next
	return value
}

func isValid(s string) bool {
	var stack *ListNode

	bracketPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			push(&stack, char)
		case ')', '}', ']':
			if pop(&stack) != bracketPairs[char] {
				return false
			}
		}
	}

	return stack == nil
}
