package main

import (
    "fmt"
    "strings"
)


func isBalanced(myStr string) string {
    stack := []string{}
    openList := []string{"[", "{", "("}
    closeList := []string{"]", "}", ")"}

    joinedOpenList := strings.Join(openList, "")

    for _, char := range myStr {
        if strings.Contains(joinedOpenList, string(char)) {
            stack = append(stack, string(char))
        } else {
            for idx, closingChar := range closeList { // Iterate through closeList characters
                if string(char) == closingChar {
                    if len(stack) == 0 || openList[idx] != stack[len(stack)-1] {
                        return "Unbalanced"
                    }
                    stack = stack[:len(stack)-1]
                    break // Exit inner loop if match found
                }
            }
        }
    }

    if len(stack) == 0 {
        return "Balanced"
    }
    return "Unbalanced"
}
  

func main() {
	strings := []string{"{[]{()}}", "[{}{})(]", "((()"}
	for _, str := range strings {
		fmt.Println(str, "-", isBalanced(str))
	}
}
