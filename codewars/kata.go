package codewars

import (
	"fmt"
	"strings"
)

// RoundToNext5
// Given an integer as input, can you round it to the next (meaning, "greater than or equal") multiple of 5?
//
// Examples:
//
// input:    output:
// 0    ->   0
// 2    ->   5
// 3    ->   5
// 12   ->   15
// 21   ->   25
// 30   ->   30
// -2   ->   0
// -5   ->   -5
// etc.
// Input may be any positive or negative integer (including 0).
//
// You can assume that all inputs are valid integers.
func RoundToNext5(n int) int {
	if n < 0 {
		return n / 5 * 5
	} else {
		return ((n + 4) / 5) * 5
	}
}

// Summation
// Write a program that finds the summation of every number from 1 to num (both inclusive).
// The number will always be a positive integer greater than 0. Your function only needs
// to return the result, what is shown between parentheses in the example below is how you
//  reach that result and it's not part of it, see the sample tests.

func Summation(n int) int {
	return n * (n + 1) / 2
}

// Acknowledgments:
// I thank yvonne-liu for the idea and for the example tests :)
//
// Description:
// Encrypt this!
//
// You want to create secret messages which can be deciphered by the Decipher this! kata. Here are the conditions:
//
// Your message is a string containing space separated words.
// You need to encrypt each word in the message using the following rules:
// The first letter must be converted to its ASCII code.
// The second letter must be switched with the last letter
// Keepin' it simple: There are no special characters in the input.
// Examples:
// EncryptThis("Hello") == "72olle"
// EncryptThis("good") == "103doo"
// EncryptThis("hello world") == "104olle 119drlo"
// https://www.codewars.com/kata/5848565e273af816fb000449/train/go
func EncryptThis(text string) string {
	splits := strings.Split(text, " ")
	sb := strings.Builder{}

	for i, v := range splits {
		if i != 0 {
			sb.WriteString(" ")
		}
		fmt.Fprintf(&sb, "%d", v[0])
		sb.WriteByte(v[len(v)-1])
		sb.WriteString(v[2 : len(v)-1])
		sb.WriteByte(v[1])
	}

	return sb.String()
}
