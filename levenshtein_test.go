package levenshtein

import "testing"

type testCase struct {
	s, t string
	dist int
}

var testCases = []testCase{
	{"", "", 0},
	{"a", "", 1},
	{"", "a", 1},
	{"aa", "a", 1},
	{"a", "aa", 1},
	{"aa", "aa", 0},
	{"", "aa", 2},
	{"aa", "", 2},
	{"shitting", "hitting", 1},
	{"shooting", "shoting", 1},
	{"shooting", "shot", 4},
	{"shooting", "shit", 5},
	{"shooting", "shotng", 2},
	{"bikeshed", "hikeshed", 1},
	{"badman", "bagmen", 2},
	{"balderdash", "dash", 6},
	{"backgammon", "chess", 9},
}

var testCasesCaseInsensitive = []testCase{
	{"", "", 0},
	{"A", "", 1},
	{"", "a", 1},
	{"aA", "a", 1},
	{"a", "Aa", 1},
	{"aA", "aa", 0},
	{"", "aa", 2},
	{"aa", "", 2},
	{"shitting", "hitTing", 1},
	{"shOoting", "shoTing", 1},
	{"shOoting", "shoT", 4},
	{"shooting", "shiT", 5},
	{"shooting", "Shotng", 2},
	{"bikeshed", "hiKeshed", 1},
	{"badman", "bagmen", 2},
	{"balderdash", "dAsh", 6},
	{"backgammon", "Chess", 9},
	{"CHESS", "CHESS", 0},
}

func TestDistance(t *testing.T) {
	for _, tc := range testCases {
		if actual := Distance(tc.s, tc.t); actual != tc.dist {
			t.Errorf("levenshtein.Distance(%s, %s) != %d (actual: %d)", tc.s, tc.t, tc.dist, actual)
		}
	}
}

func TestDistanceCaseInsensitive(t *testing.T) {
	for _, tc := range testCasesCaseInsensitive {
		if actual := DistanceCaseInsensitive(tc.s, tc.t); actual != tc.dist {
			t.Errorf("levenshtein.DistanceCaseInsensitive(%s, %s) != %d (actual: %d)", tc.s, tc.t, tc.dist, actual)
		}
	}
}
