package validateparentheses

import "fmt"

type stack struct {
	data []string
}

func (s *stack) push(val string) error {
	if s == nil {
		return fmt.Errorf("nil obj")
	}
	s.data = append(s.data, val)
	return nil
}

func (s *stack) pop() (string, error) {
	if len(s.data) == 0 {
		return "", fmt.Errorf("empty stack")
	}
	index := len(s.data) - 1
	output := s.data[index]
	// s.data[index] = ""
	s.data = s.data[:index]
	return output, nil
}

func (s *stack) isEmpty() bool {
	if len(s.data) == 0 {
		return true
	}
	return false
}

func isValid(in string) bool {
	bracketMap := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	stk := &stack{}
	for i := range in {
		if val, ok := bracketMap[string(in[i])]; !ok {
			stk.push(string(in[i]))
		} else {
			if poped, err := stk.pop(); poped != val || err != nil {
				return false
			}
		}

	}
	if !stk.isEmpty() {
		return false
	}
	return true
}

// short version
// func isValid(in string) bool {
// 	bracketMap := map[string]string{
// 		"}": "{",
// 		"]": "[",
// 		")": "(",
// 	}
//     var stk []string
// 	for i := range in {
// 		if val, ok := bracketMap[string(in[i])]; !ok {
//             stk = append(stk,string(in[i]))
// 		} else if len(stk) == 0 || stk[len(stk)-1] != val {
//             return false
//         } else {
//             stk = stk[:len(stk)-1]
//         }
// 	}
//     return len(stk) == 0
// }
