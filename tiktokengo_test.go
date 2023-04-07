package tiktokengo_test

import (
	"testing"

	"github.com/devfullcycle/tiktokengo"
)

func TestCountTokens(t *testing.T) {
    testCases := []struct {
        prompt string
        count  int
    }{
        {prompt: "", count: 0},
        {prompt: "Hello world", count: 2},
        {prompt: "foo bar baz", count: 3},
        // add more test cases as needed
    }

    for _, tc := range testCases {
        got := tiktokengo.CountTokens(tc.prompt)
        if got != tc.count {
            t.Errorf("CountTokens(%q) = %v; want %v", tc.prompt, got, tc.count)
        }
    }
}
