package util

import (
	"testing"
)

func TestNode_Match(t *testing.T) {
	type test struct {
		arg  string
		want bool
	}
	tests := []struct {
		name  string
		node  *Node
		tests []test
	}{{
		"simple",
		NewNodeFromRules(".", "a.wweir.cc", "b.wweir.cc"),
		[]test{
			{"a.wweir.cc", true},
			{"b.wweir.cc", true},
		},
	}, {
		"parent",
		NewNodeFromRules(".", "wweir.cc", "a.wweir.cc"),
		[]test{
			{"wweir.cc", true},
			{"a.wweir.cc", true},
			{"b.wweir.cc", false},
		},
	}, {
		"fuzz1",
		NewNodeFromRules(".", "wweir.cc", "*.wweir.cc"),
		[]test{
			{"wweir.cc", true},
			{"a.wweir.cc", true},
		},
	}, {
		"fuzz2",
		NewNodeFromRules(".", "a.wweir.cc", "*.wweir.cc"),
		[]test{
			{"wweir.cc", true},
			{"a.wweir.cc", true},
			{"b.wweir.cc", true},
		},
	}, {
		"fuzz3",
		NewNodeFromRules(".", "a.*.cc"),
		[]test{
			{"wweir.cc", false},
			{"a.wweir.cc", true},
			{"b.wweir.cc", false},
		},
	}, {
		"fuzz4",
		NewNodeFromRules(".", "*.*.cc", "iamp.*.*"),
		[]test{
			{"wweir.cc", false},
			{"a.wweir.cc", true},
			{"b.wweir.cc", true},
			{"iamp.wweir.cc", true},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, test := range tt.tests {
				if got := tt.node.Match(test.arg); got != test.want {
					t.Errorf("Node.Match(%s) = %v, want %v", test.arg, got, test.want)
				}
			}
		})
	}
}