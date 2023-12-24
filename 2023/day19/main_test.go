package main

import (
	_ "embed"
	"testing"
)

var example1 = `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  19114,
		},
		{
			name:  "actual",
			input: input,
			want:  420739,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  167409079868000,
		},
		{
			name:  "actual",
			input: input,
			want:  130251901420382,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateRangeAccepted(t *testing.T) {
	tests := []struct {
		name string
		r    Range
		c    Comparison
		want Range
	}{
		{
			name: "greater than x - condition updates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{X, '>', 2000},
			want: Range{2001, 4000, 1, 4000, 1, 4000, 1, 4000},
		},
		{
			name: "greater than x - condition does not update range",
			r:    Range{2000, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{X, '>', 1},
			want: Range{2000, 4000, 1, 4000, 1, 4000, 1, 4000},
		},
		{
			name: "greater than x - condition invalidates range",
			r:    Range{1, 2000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{X, '>', 3000},
			want: Range{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "greater than m - condition updates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{M, '>', 2000},
			want: Range{1, 4000, 2001, 4000, 1, 4000, 1, 4000},
		},
		{
			name: "greater than m - condition does not update range",
			r:    Range{1, 4000, 2000, 4000, 1, 4000, 1, 4000},
			c:    Comparison{M, '>', 1},
			want: Range{1, 4000, 2000, 4000, 1, 4000, 1, 4000},
		},
		{
			name: "greater than m - condition invalidates range",
			r:    Range{1, 4000, 1, 2000, 1, 4000, 1, 4000},
			c:    Comparison{M, '>', 3000},
			want: Range{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "greater than a - condition updates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{A, '>', 2000},
			want: Range{1, 4000, 1, 4000, 2001, 4000, 1, 4000},
		},
		{
			name: "greater than a - condition does not update range",
			r:    Range{1, 4000, 1, 4000, 2000, 4000, 1, 4000},
			c:    Comparison{A, '>', 1},
			want: Range{1, 4000, 1, 4000, 2000, 4000, 1, 4000},
		},
		{
			name: "greater than a - condition invalidates range",
			r:    Range{1, 4000, 1, 4000, 1, 2000, 1, 4000},
			c:    Comparison{A, '>', 3000},
			want: Range{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "greater than s - condition updates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{S, '>', 2000},
			want: Range{1, 4000, 1, 4000, 1, 4000, 2001, 4000},
		},
		{
			name: "greater than s - condition does not update range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 2000, 4000},
			c:    Comparison{S, '>', 1},
			want: Range{1, 4000, 1, 4000, 1, 4000, 2000, 4000},
		},
		{
			name: "greater than s - condition invalidates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 1, 2000},
			c:    Comparison{S, '>', 3000},
			want: Range{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "less than than x - condition updates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{X, '<', 2000},
			want: Range{1, 1999, 1, 4000, 1, 4000, 1, 4000},
		},
		{
			name: "less than than x - condition does not update range",
			r:    Range{2000, 3999, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{X, '<', 4000},
			want: Range{2000, 3999, 1, 4000, 1, 4000, 1, 4000},
		},
		{
			name: "less than than x - condition invalidates range",
			r:    Range{2000, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{X, '<', 1000},
			want: Range{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "less than than m - condition updates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{M, '<', 2000},
			want: Range{1, 4000, 1, 1999, 1, 4000, 1, 4000},
		},
		{
			name: "less than than a - condition does not update range",
			r:    Range{1, 4000, 1, 4000, 2000, 3999, 1, 4000},
			c:    Comparison{A, '<', 4000},
			want: Range{1, 4000, 1, 4000, 2000, 3999, 1, 4000},
		},
		{
			name: "less than than s - condition invalidates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 2000, 4000},
			c:    Comparison{S, '<', 1000},
			want: Range{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateRangeAccepted(tt.r, tt.c); got != tt.want {
				t.Errorf("updateRangeAccepted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateRangeRejected(t *testing.T) {
	tests := []struct {
		name string
		r    Range
		c    Comparison
		want Range
	}{
		{
			name: "x - condition updates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{X, '>', 2000},
			want: Range{1, 2000, 1, 4000, 1, 4000, 1, 4000},
		},
		{
			name: "x - condition does not update range",
			r:    Range{1, 2000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{X, '>', 2000},
			want: Range{1, 2000, 1, 4000, 1, 4000, 1, 4000},
		},
		{
			name: "x - condition invalidates range",
			r:    Range{3001, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{X, '>', 3000},
			want: Range{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "m - condition updates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{M, '>', 2000},
			want: Range{1, 4000, 1, 2000, 1, 4000, 1, 4000},
		},
		{
			name: "m - condition does not update range",
			r:    Range{1, 4000, 1, 2000, 1, 4000, 1, 4000},
			c:    Comparison{M, '>', 2000},
			want: Range{1, 4000, 1, 2000, 1, 4000, 1, 4000},
		},
		{
			name: "m - condition invalidates range",
			r:    Range{1, 4000, 3001, 4000, 1, 4000, 1, 4000},
			c:    Comparison{M, '>', 3000},
			want: Range{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "a - condition updates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{A, '>', 2000},
			want: Range{1, 4000, 1, 4000, 1, 2000, 1, 4000},
		},
		{
			name: "a - condition does not update range",
			r:    Range{1, 4000, 1, 4000, 1, 2000, 1, 4000},
			c:    Comparison{A, '>', 2000},
			want: Range{1, 4000, 1, 4000, 1, 2000, 1, 4000},
		},
		{
			name: "a - condition invalidates range",
			r:    Range{1, 4000, 1, 4000, 3001, 4000, 1, 4000},
			c:    Comparison{A, '>', 3000},
			want: Range{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "s - condition updates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{S, '>', 2000},
			want: Range{1, 4000, 1, 4000, 1, 4000, 1, 2000},
		},
		{
			name: "s - condition does not update range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 1, 2000},
			c:    Comparison{S, '>', 2000},
			want: Range{1, 4000, 1, 4000, 1, 4000, 1, 2000},
		},
		{
			name: "s - condition invalidates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 3001, 4000},
			c:    Comparison{S, '>', 3000},
			want: Range{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "less than than x - condition updates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{X, '<', 2000},
			want: Range{2000, 4000, 1, 4000, 1, 4000, 1, 4000},
		},
		{
			name: "less than than x - condition does not update range",
			r:    Range{2000, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{X, '<', 2000},
			want: Range{2000, 4000, 1, 4000, 1, 4000, 1, 4000},
		},
		{
			name: "less than than x - condition invalidates range",
			r:    Range{1, 3999, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{X, '<', 4000},
			want: Range{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "less than than m - condition updates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 1, 4000},
			c:    Comparison{M, '<', 2000},
			want: Range{1, 4000, 2000, 4000, 1, 4000, 1, 4000},
		},
		{
			name: "less than than a - condition does not update range",
			r:    Range{1, 4000, 1, 4000, 2000, 4000, 1, 4000},
			c:    Comparison{A, '<', 2000},
			want: Range{1, 4000, 1, 4000, 2000, 4000, 1, 4000},
		},
		{
			name: "less than than s - condition invalidates range",
			r:    Range{1, 4000, 1, 4000, 1, 4000, 1, 3999},
			c:    Comparison{S, '<', 4000},
			want: Range{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateRangeRejected(tt.r, tt.c); got != tt.want {
				t.Errorf("updateRangeRejected() = %v, want %v", got, tt.want)
			}
		})
	}
}
