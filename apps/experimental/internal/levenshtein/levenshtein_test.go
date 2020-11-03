package levenshtein

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"testing"
	"time"
	"unicode"
)

func Test_min(t *testing.T) {
	tests := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("permutation #%d", i+1), func(t *testing.T) {
			if got := min(tt.a, tt.b, tt.c); got != 1 {
				t.Errorf("min() = %v, want %v", got, 1)
			}
		})
	}
}

func TestLevenshtein(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "size difference #1",
			args: args{
				a: "abcdef",
				b: "abcdefg",
			},
			want: 1,
		},
		{
			name: "size difference #2",
			args: args{
				a: "abcdefgh",
				b: "abcdef",
			},
			want: 2,
		},
		{
			name: "s is empty",
			args: args{
				a: "",
				b: "But in the end it’s only a passing thing, this shadow; ev" +
					"en darkness must pass.",
			},
			want: 79,
		},
		{
			name: "t is empty",
			args: args{
				a: "All we have to decide is what to do with the time that is" +
					" given us.",
				b: "",
			},
			want: 67,
		},
		{
			name: "strings are equal",
			args: args{
				a: "Nechť již hříšné saxofony ďáblů rozezvučí síň úděsnými tó" +
					"ny waltzu, tanga a quickstepu.",
				b: "Nechť již hříšné saxofony ďáblů rozezvučí síň úděsnými tó" +
					"ny waltzu, tanga a quickstepu.",
			},
			want: 0,
		},
		{
			name: "simple 1",
			args: args{
				a: "kitten",
				b: "sitting",
			},
			want: 3,
		},
		{
			name: "simple 2",
			args: args{
				a: "Saturday",
				b: "Sunday",
			},
			want: 3,
		},
		{
			name: "Japanese",
			args: args{
				a: "持ち上げて",
				b: "解き放して",
			},
			want: 4,
		},
	}
	t.Run("LinearSpace", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := LinearSpace(tt.args.a, tt.args.b); got != tt.want {
					t.Errorf("LinearSpace() = %v, want %v", got, tt.want)
				}
			})
		}
	})
	t.Run("WagnerFischer", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := WagnerFischer(tt.args.a, tt.args.b); got != tt.want {
					t.Errorf("WagnerFischer() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomString(n int) string {
	var b strings.Builder

	for i := 0; i < n; i++ {
		b.WriteRune(rand.Int31n(unicode.MaxRune + 1))
	}

	return b.String()
}

const fast = true

func TestFuzzyLevenshtein(t *testing.T) {
	if fast {
		t.Skip()
	}
	for i := 0; i < math.MaxUint16; i++ {
		t.Logf("i = %d", i)
		s1, s2 := randomString(i-rand.Intn(i+1)), randomString(i-rand.Intn(i+1))
		if LinearSpace(s1, s2) != WagnerFischer(s1, s2) {
			t.Errorf("LS(%s, %s) != WF(%s, %s)", s1, s2, s1, s2)
		}
		if LinearSpace(s1, s2) != WagnerFischer(s2, s1) {
			t.Errorf("LS(%s, %s) != WF(%s, %s)", s1, s2, s2, s1)
		}
		if LinearSpace(s2, s1) != WagnerFischer(s1, s2) {
			t.Errorf("LS(%s, %s) != WF(%s, %s)", s2, s1, s1, s2)
		}
		if LinearSpace(s2, s1) != WagnerFischer(s2, s1) {
			t.Errorf("LS(%s, %s) != WF(%s, %s)", s2, s1, s2, s1)
		}
	}
}

var a1e3, b1e3 = randomString(1e3), randomString(1e3)

var result int

func BenchmarkLinearSpace(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = LinearSpace(a1e3, b1e3)
	}
	result = r
}

func BenchmarkWagnerFischer(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = WagnerFischer(a1e3, b1e3)
	}
	result = r
}
