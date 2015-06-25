package multicache

import "testing"

/**
This file is part of go-multicache, a library for handling caches with multiple
keys and replacement algorithms.

Copyright 2015 Joseph Lewis <joseph@josephlewis.net>
Licensed under the MIT license
**/

var roundRobinTestCases = []ReplacementAlgorithmTestcase{
	// Miss all items because they aren't in cache
	{&RoundRobin{}, 4, []string{"a", "b", "c", "d", "e"}, []bool{false, false, false, false, false}},
	// Overwrite the first element and try it again
	{&RoundRobin{}, 3, []string{"a", "b", "c", "d", "a"}, []bool{false, false, false, false, false}},
	// Overwrite multiple in a row
	{&RoundRobin{}, 2, []string{"a", "b", "c", "a", "b", "c"}, []bool{false, false, false, false, false, false}},
	// Hit things
	{&RoundRobin{}, 2, []string{"a", "b", "a", "b"}, []bool{false, false, true, true}},
	// Out of order access
	{&RoundRobin{}, 2, []string{"a", "b", "b", "a"}, []bool{false, false, true, true}},
	// Make sure we didn't get rid of b
	{&RoundRobin{}, 2, []string{"a", "b", "c", "b"}, []bool{false, false, false, true}}}

func TestRoundRobin(t *testing.T) {

	for _, testcase := range roundRobinTestCases {
		testcase.RunTest(t)
	}
}
