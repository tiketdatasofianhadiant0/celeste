package celeste

import (
	"errors"
	"fmt"
	"strings"
)

type FindToken map[string]any

func (t FindToken) String() string {
	if val, ok := t[ConjunctionAnd]; ok {
		if arr, ok := val.([]any); ok {
			return t.conjunctionAnd(arr)
		}

		return ""
	}

	if val, ok := t[ConjunctionOr]; ok {
		if arr, ok := val.([]any); ok {
			return t.conjunctionOr(arr)
		}

		return ""
	}

	return ""
}

func (t FindToken) conjunctionAnd(arr []any) string {
	var result []string
	for _, val := range arr {
		if token, ok := val.(map[string]any); ok {
			result = append(result, t.processToken(token))
		}
	}

	if len(result) == 0 {
		return ""
	}

	return strings.Join(result, KeywordAnd)
}

func (t FindToken) conjunctionOr(arr []any) string {
	var result []string
	for _, val := range arr {
		if token, ok := val.(map[string]any); ok {
			result = append(result, t.processToken(token))
		}
	}

	if len(result) == 0 {
		return ""
	}

	return strings.Join(result, KeywordOr)
}

func (t FindToken) processToken(token map[string]any, prevToken ...string) string {
	formatAny := func(val any) (string, error) {
		if i, ok := val.(int64); ok {
			return fmt.Sprintf("%v", i), nil
		}

		if f, ok := val.(float64); ok {
			return fmt.Sprintf("%v", f), nil
		}

		if s, ok := val.(string); ok {
			return fmt.Sprintf("%s", s), nil
		}

		if arr, ok := val.([]any); ok {
			var vs []string
			for _, v := range arr {
				vs = append(vs, fmt.Sprintf("%v", v))
			}
			return fmt.Sprintf("(%s)", strings.Join(vs, ", ")), nil
		}

		return "", errors.New("not supported type")
	}

	var result []string
	for key, val := range token {
		if op, exists := operatorMap[key]; exists {
			if v, err := formatAny(val); err == nil {
				result = append(result, fmt.Sprintf("%s %s", op, v))
			} else {
				if token, ok := val.(map[string]any); ok {
					result = append(result, fmt.Sprintf("%s %s", op, t.processToken(token)))
				}
			}

			continue
		}

		if logicalOp, exists := logicalMap[key]; exists {
			if v, err := formatAny(val); err == nil {
				result = append(result, fmt.Sprintf("%s %s", v, logicalOp))
			} else {
				if token, ok := val.(map[string]any); ok {
					result = append(result, fmt.Sprintf("%s %s", logicalOp, t.processToken(token)))
				}
			}

			continue
		}

		if fn, exists := functionMap[key]; exists {
			if v, err := formatAny(val); err == nil {
				result = append(result, fmt.Sprintf("%s(%s)", fn, v))
			} else {
				if token, ok := val.(map[string]any); ok {
					sub := t.processToken(token)
					if idx := strings.IndexAny(sub, ComparisonOperators); idx > 0 {
						sub = "(" + sub[:idx-1] + ")" + sub[idx-1:]
					}
					result = append(result, fmt.Sprintf("%s%s", fn, sub))
				}
			}
			continue
		}

		if v, err := formatAny(val); err == nil {
			result = append(result, fmt.Sprintf("%s = %s", key, v))
			continue
		}

		if token, ok := val.(map[string]any); ok {
			result = append(result, fmt.Sprintf("%s %s", key, t.processToken(token)))
			continue
		}
	}

	if len(result) == 0 {
		return ""
	}

	if len(result) == 1 {
		return result[0]
	}

	return "(" + strings.Join(result, KeywordAnd) + ")"
}
