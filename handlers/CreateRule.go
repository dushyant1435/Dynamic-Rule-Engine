package handlers

import (
	"regexp"
	"strings"
	"y/models"
)

func parseRuleString(ruleString string) []string {
	re := regexp.MustCompile(`[\w']+|[()=><!]`)
	tokens := re.FindAllString(ruleString, -1)
	return tokens
}

// CreateRule converts a rule string into an AST Node
func CreateRule(ruleString string) *models.Node {
	tokens := parseRuleString(ruleString)
	ast, _ := convertTokensToAST(tokens, 0)
	return ast
}

// convertTokensToAST recursively parses tokens and builds the AST
func convertTokensToAST(tokens []string, index int) (*models.Node, int) {
	token := tokens[index]

	if token == "(" {
		left, idx := convertTokensToAST(tokens, index+1)
		operator := tokens[idx]
		right, idx := convertTokensToAST(tokens, idx+1)
		return &models.Node{
			Type:  models.Operator,
			Left:  left,
			Right: right,
			Value: operator,
		}, idx + 1
	} else if token == ")" {
		return nil, index
	} else {
		// Handle condition or logical operators
		if strings.ToUpper(token) == "AND" || strings.ToUpper(token) == "OR" {
			return &models.Node{
				Type:  models.Operator,
				Value: strings.ToUpper(token),
			}, index + 1
		}

		// Handle condition
		condition := token
		if index+2 < len(tokens) && (tokens[index+1] == ">" || tokens[index+1] == "<" || tokens[index+1] == "=" || tokens[index+1] == "!=") {
			condition += " " + tokens[index+1] + " " + tokens[index+2]
			index += 2
		}
		return &models.Node{
			Type:  models.Operand,
			Value: condition,
		}, index + 1
	}
}
