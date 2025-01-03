package main

import (
	_ "embed"
	"testing"
)

var example1 = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

var realInput1 = `bwb, wrrbg, ubwg, rwg, urbbgrr, wugr, rububww, rbrw, ggw, grgwu, wb, gwuuubw, gbr, ugur, rug, ugw, rg, uru, wgrbgug, buwwrwu, bbbuw, urbug, rrbr, uwbr, wurr, wbwrw, bug, rggwguru, brrgub, uuwg, gww, gwb, ubrb, wgrg, rubbgwu, bbb, buu, gbgrwg, rwuwbg, uubwbu, buuu, uwurwgu, gwbugwgw, guru, brww, bwgrugb, rbbgrw, ruubg, wub, wbw, bgww, bgugg, rru, urrwuw, bwwgw, brbww, gwug, bwwwbwb, uubggg, gbgb, rbgub, bwuwgg, wrbb, buwrw, gbbwrrg, gug, wuww, u, ru, gbbbu, ugwrwgb, rgb, rwwuwuwb, ruugwrb, rgruw, wwuwwbbw, urrrbu, bwg, ub, urg, ruwg, bwr, wbwb, brbwrwww, rbwgrgub, gbb, grwwb, bwggw, urwr, rrrug, buw, gbbgrguu, rbgg, brbwur, ububwb, bur, wruru, buurbu, ggb, ggr, rggwrub, uwbbbbgb, rrrb, uwr, bwwrr, brgw, urrwug, gub, rwwb, wgw, ggur, bbrb, uwrbgbu, buubu, urb, bg, gguubbrg, gg, rwr, uubb, wwb, uwgbub, brbgrrb, ubbbrwb, ugr, ggwwug, wbgbu, grw, wrb, brr, rub, gr, guwb, ggrgr, rrb, uuurb, brw, wrg, guw, uggr, uuurgw, wrrgu, wwuu, uruuu, bwrr, wwg, bwguu, rugww, rbw, wur, buuuw, rwb, bbg, gbu, rurw, brwbbb, wbbgguw, rrrr, bwbbb, bgu, uurbb, gwrbwb, wgwrgbwg, wgrw, gugw, urrbb, rww, ruu, uwbrbb, bbbrwb, bugw, wug, wgg, rwru, r, uwg, wggr, brbwgb, uwbuwg, ubr, rbgwrrw, grwbur, grrrbb, wgr, rrbw, uurw, wbgg, buwg, wru, ruwwu, wbbbrg, grgbgbr, uuu, wrr, ubwr, rubbgb, gbrw, gbug, uwrur, ugru, uruggr, ubu, uwgg, wrgb, wbg, ug, guu, bgw, wbu, wuru, rur, wwbwwr, wbugg, wrwgwwb, rrgu, uuugg, bgbb, gb, uwugwgu, bguw, ugwgwugb, brburg, grrbw, brb, bguwb, bwugu, ubwgg, rbr, wwrb, bwgrug, rbww, gwg, rgrgu, bwbgrgb, wwbu, bwugw, uguu, uubuggwb, uwb, gbg, uubur, urwbub, uggrr, wrurb, bbr, uggg, gur, wrwrw, bbbrug, bbuw, b, wbgw, bubgubg, ggg, bru, gbuwr, ruw, wgu, ubb, gru, bb, uwbug, ugbrgu, uruwrrbb, rrug, bgbu, ggwgbrr, rwuuuuu, bgr, rrwu, ruub, burw, bbww, uuuu, gwwb, wrur, grb, br, bwwgg, gburg, ugubw, gbwbwbb, gwrgubrg, bguwwr, grgugbw, rugw, guww, brrrr, uguurb, uwwbbw, guwgu, guguubr, rwu, bbbwb, wggru, bub, rwrg, gw, bw, wrgbrb, ggbru, gwbw, wuwgu, rwgurr, gugr, rbu, wrbguww, uuww, rbuwu, uwu, rburgwu, wgbbuu, rggrur, wbr, buubb, rgwb, bgg, rwurbg, urgrg, rbrrr, rbur, wr, uwuruggr, buug, wubw, uubbbu, gwu, ubgg, rgu, uwwr, brgrgw, uuw, wwr, wu, grgwbbb, rgrbr, bbu, wuwgr, rrg, urw, urgbgr, bwwb, gurgu, rr, rgr, rbg, rgrr, uwgug, gwggbw, rgg, ubgr, ggww, gugggr, rrbgrw, rrr, ubw, uw, gbw, wrww, uwwwrwbu, bgwrugu, guur, grrruuw, ubrgg, ggrb, grg, bgrru, rurrb, ugg, wwgr, uwbrg, wuw, ubbu, rbb, www, wwubgu, ur, bububg, gwr, grgguwg, bubr, wwu, ggru, ugu, rgrru, uburu, wuu, wbugrwu, uur, urwgrwg, bwgwgwg, ubgru, burugw, wgbwrr, urr, wgb, g, ruwggg, wbb, bbrrg, wbuw, uubwugw, rwrr, uubgr, uww, gu, uubu, rrw, ugb, buuwr, uug, rbwbubgb, wgwburur, rrurr, uguw, wbrw, ggu, ruwgwb, rgw, ubrbgu, ururb, rb, ruuu, ubg, rbbbggwu, wg, wbrbbg, gwwwuu, bugg, uu, rbgwrwu, gggbbu, rbggu, wbwwr, uwbbu, ggrwu, brg, buub, bggguu, bu, gbrubgwr, wrwwb, bbw, wubrbb, bbwbbur, bww, uwwuuw

wgurrgbbgwuuwwbwwuwggwrbggrbrwrgbbwwwbbbubuuwbwuuguwr`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  6,
		},
		{
			name:  "line 1 of input",
			input: realInput1,
			want:  1,
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
			want:  16,
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
