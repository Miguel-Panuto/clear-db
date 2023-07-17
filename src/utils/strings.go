package utils

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func SplitFirst(s string, sep string) []string {
	if len(s) == 0 {
		return []string{}
	}

	return strings.SplitN(s, sep, 2)
}

func Split(s string, sep string) []string {
	if len(s) == 0 {
		return []string{}
	}

	return strings.Split(s, sep)
}

func TrimSplit(s string, sep string) []string {
	arr := Split(s, sep)

	for i := range arr {
		arr[i] = strings.TrimSpace(arr[i])
	}

	return arr
}

func ContainsMany(s string, values ...string) bool {
	contains := true

	for _, el := range values {
		if !strings.Contains(s, el) {
			contains = false
		}
	}

	return contains
}

func IsBetween(s string, from string, until string) bool {
	return strings.HasPrefix(s, from) && strings.HasSuffix(s, until)
}

func SubString(s string, from string, until string) string {
	fromIndex := strings.Index(s, from) + 1
	untilIndex := strings.Index(s, until)
	substring := s[fromIndex:untilIndex]
	return substring
}

func SubSplit(s string, from string, until string, sep string) []string {
	s = SubString(s, from, until)
	splited := TrimSplit(s, sep)
	return splited
}

func VerifyLowerPrefix(s string, prefix string) bool {
	s = strings.ToLower(s)

	s = strings.TrimSpace(s)

	return strings.HasPrefix(s, prefix)
}

func VerifyLower(s string, command string) bool {
	aux := strings.ToLower(s)

	aux = strings.TrimSpace(aux)

	return aux == command
}

func MakeStringArr(i []any) []string {
	arr := []string{}
	for _, value := range i {
		arr = append(arr, fmt.Sprint(value))
	}

	return arr
}

func ContainsInside(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func MultipleSplit(s string, splitters ...string) []string {
	reg := "(?i)"
	for _, splitter := range splitters {
		reg += fmt.Sprintf(`\s*%s\s*|`, regexp.QuoteMeta(splitter))
	}
	reg = strings.TrimSuffix(reg, "|")
	regexSplitter := regexp.MustCompile(reg)
	return regexSplitter.Split(s, -1)
}

func SplitByOperators(s string, operators ...string) []string {
	sort.Slice(operators, func(i, j int) bool {
		return len(operators[i]) > len(operators[j])
	})

	ops := make([]string, len(operators))
	for i, operator := range operators {
		ops[i] = regexp.QuoteMeta(strings.TrimSpace(operator))
	}
	opsRegex := strings.Join(ops, "|")

	opsRegex = fmt.Sprintf(`(?i)(%s)`, opsRegex)

	re := regexp.MustCompile(opsRegex)

	parts := re.Split(s, -1)

	matches := re.FindAllString(s, -1)

	res := make([]string, 0, len(parts)+len(matches))
	for i := range parts {
		trimmedPart := strings.TrimSpace(parts[i])
		res = append(res, trimmedPart)
		if i < len(matches) {
			res = append(res, strings.TrimSpace(matches[i]))
		}
	}
	return res
}
