package envconv

import (
	"fmt"
	"strings"
)

func MapFromArr(a []string) map[string]string {
	m := map[string]string{}
	for _, s := range a {
		kv := strings.Split(s, "=")
		if len(kv) >= 2 && kv[0] != "" {
			m[kv[0]] = kv[1]
		}
	}
	return m
}

func ToMap(ss ...string) map[string]string {
	return MapFromArr(ss)
}

func ArrFromMap(m map[string]string) []string {
	a := []string{}
	for k, v := range m {
		if k != "" {
			a = append(a, fmt.Sprintf("%s=%s", k, v))
		}
	}
	return a
}

func ToArr(ss ...string) []string {
	a := []string{}
	for i, s := range ss {
		if i%2 == 1 {
			sm1 := ss[i-1]
			if sm1 != "" {
				a = append(a, fmt.Sprintf("%s=%s", sm1, s))
			}
		}
	}
	return ArrFromMap(ToMap(a...))
}
