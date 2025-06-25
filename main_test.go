package main

import (
	"testing"
)

func TestCalculateRPN(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		err      bool
	}{
		{"3 4 +", 7, false},
		{"5 2 -", 3, false},
		{"6 7 *", 42, false},
		{"10 2 /", 5, false},
		{"3 4 + 2 *", 14, false},
		{"10 2 3 + /", 2, false}, // 10 / (2 + 3) = 10 / 5 = 2
		{"", 0, true},           // 空文字列
		{"3 +", 0, true},        // 演算子のみ
		{"3 4", 0, true},        // 数値のみ
		{"abc", 0, true},        // 不正な入力
		{"3 4 + -", 0, true},    // スタック不足
		{"3 0 /", 0, true},      // ゼロ除算
	}

	for _, test := range tests {
		result, err := CalculateRPN(test.input)
		if test.err {
			if err == nil {
				t.Errorf("CalculateRPN(%q) expected an error, but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("CalculateRPN(%q) returned an unexpected error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("CalculateRPN(%q) expected %f, but got %f", test.input, test.expected, result)
			}
		}
	}
}
