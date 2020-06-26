package strings_test

import (
	"hsecode.com/stdlib/strings"
	"testing"
)

func isWord(w string) bool {
	if w == "abcd" || w == "abc" || w == "aaa" || w == "bc" || w == "a" {
		return true
	} else {
		return false
	}
}

func TestSegmentation(t *testing.T) {
	answer, _ := strings.Segmentation("aabcdaaabcaaabcd", isWord)
	t.Log(answer)
	answer, _ = strings.Segmentation("aaaaaaaabcaaaaaaaaaaaaaaaaaaaaaabcdaaaaaaaaaaaaaaaaaaaaaaaaaaaaabcdaaaaaaaaaaaaaaaaabcbcbcbcbcbcaaaaaaabcd", isWord)
	t.Log(answer)
	answer, _ = strings.Segmentation("abcd", isWord)
	t.Log(answer)
}

func TestLCS(t *testing.T) {
	t.Log(strings.LCS("vintner", "writers"))
	t.Log(strings.LCS("ABCD", "ACBAD"))
	t.Log(len(strings.LCS("H", "ABCDEFG")))
	t.Log(strings.LCS("cwgbbclcjpc", "vtvzpzwwejz"))
	t.Log(strings.LCS("aeoxdejmm;cslvkbrucxl;,'xmeniubsyrfxkdl,zpexuorucjkclzpkoiujkl,;z.p[oinutrkl,zpworiuyjkhugytfgljiolynkluhigyuftolikuyjtredfghbjklizdvfjh[oighkjvnzc;kvjd[foiguahf;zcxlvn'lfijg[aiogu[afodivhdfknvbjhlkvgypiurytqeqghdk;ah;aoighrrrrrai;hlknjgliytperiytperytyqpiuerligh;aih", "dfvghbjkijouiytyrtedxcfvbhjnkijuiyutyrdtesdcvbnjsirxfn;aou/zlougrbt;rofc;zoiufk.zjhg.lwirt[pewripeutocsednug;slkb.xkgjbvh.cit;axout[pqrpeti[reaakml,uytrewesxdcfvbhnjkml,'lkopiojiuhyugyttrr"))
}
