package restoreip

import (
	"strconv"
	"strings"
)

func RestoreIPAddresses(s string) []string {

	if len(s) < 4 || len(s) > 12 {
		return nil
	}

	var res []string

	var fn func(s string, octets []string)

	fn = func(s string, octets []string) {

		if len(octets) == 4 && len(s) > 0 {
			return
		}

		if len(s) == 0 && len(octets) == 4 {
			res = append(res, strings.Join(octets, "."))
		}

		if len(s) == 0 {
			return
		}

		for l := 1; l <= 3 && l <= len(s); l++ {
			if validOctet(s[:l]) {
				fn(s[l:], append(octets, s[:l]))
			}
		}
	}

	fn(s, nil)

	return res
}

func validOctet(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	n, _ := strconv.Atoi(s) // s состоит из 1–3 цифр, ошибки не будет
	return n <= 255
}
