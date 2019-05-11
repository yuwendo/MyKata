package main

import (
	"reflect"
	"testing"
)

//Test_sumEvenFibonacci for the function that is used in SumEvenFibonacciNumbers.go
func Test_sumEvenFibonacci(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "default", args: args{limit: 8}, want: 10},
		{name: "case1", args: args{limit: 111111}, want: 60696},
		{name: "case2", args: args{limit: 8675309}, want: 4613732},
		{name: "case3", args: args{limit: 1}, want: 2},
		{name: "case4", args: args{limit: 96746}, want: 60696},
		{name: "case5", args: args{limit: 144100000}, want: 82790070},
		{name: "case6", args: args{limit: 65140000}, want: 82790070},
		{name: "case7", args: args{limit: 7347000000}, want: 6293134512},
		{name: "case8", args: args{limit: 10000000000000}, want: 8583840088782},
		{name: "case9", args: args{limit: 123456789000000}, want: 154030760585064},
		{name: "case10", args: args{limit: 2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumEvenFibonacci(tt.args.limit); got != tt.want {
				t.Errorf("sumEvenFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_fizzBuzzCuckooClock for the function that is used in FizzBuzzCuckooClock.go
func Test_fizzBuzzCuckooClock(t *testing.T) {
	type args struct {
		time string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "default", args: args{time: "13:34"}, want: "tick"},
		{name: "case1", args: args{time: "21:00"}, want: "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"},
		{name: "case2", args: args{time: "11:15"}, want: "Fizz Buzz"},
		{name: "case3", args: args{time: "03:03"}, want: "Fizz"},
		{name: "case4", args: args{time: "14:30"}, want: "Cuckoo"},
		{name: "case5", args: args{time: "08:55"}, want: "Buzz"},
		{name: "case6", args: args{time: "00:00"}, want: "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"},
		{name: "case7", args: args{time: "12:00"}, want: "Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo Cuckoo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzzCuckooClock(tt.args.time); got != tt.want {
				t.Errorf("fizzBuzzCuckooClock() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_bandNameGenerator for the function that is used in BandNameGenerator.go
func Test_bandNameGenerator(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "default", args: args{word: "dolphin"}, want: "The Dolphin"},
		{name: "case1", args: args{word: "alaska"}, want: "Alaskalaska"},
		{name: "case2", args: args{word: "knife"}, want: "The Knife"},
		{name: "case3", args: args{word: "tart"}, want: "Tartart"},
		{name: "case4", args: args{word: "sandles"}, want: "Sandlesandles"},
		{name: "case5", args: args{word: "bed"}, want: "The Bed"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bandNameGenerator(tt.args.word); got != tt.want {
				t.Errorf("bandNameGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_isTriangle for the function that is used in IsThisATriangle.go
func Test_isTriangle(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "default", args: args{a: 1, b: 2, c: 2}, want: true},
		{name: "case1", args: args{a: 7, b: 2, c: 2}, want: false},
		{name: "case2", args: args{a: 1, b: 2, c: 3}, want: false},
		{name: "case3", args: args{a: 3, b: 1, c: 2}, want: false},
		{name: "case4", args: args{a: 5, b: 1, c: 2}, want: false},
		{name: "case5", args: args{a: 1, b: 2, c: 5}, want: false},
		{name: "case6", args: args{a: 2, b: 5, c: 1}, want: false},
		{name: "case7", args: args{a: 4, b: 2, c: 3}, want: true},
		{name: "case8", args: args{a: 5, b: 1, c: 5}, want: true},
		{name: "case9", args: args{a: 2, b: 2, c: 2}, want: true},
		{name: "case10", args: args{a: -1, b: 2, c: 3}, want: false},
		{name: "case11", args: args{a: 1, b: -2, c: 3}, want: false},
		{name: "case12", args: args{a: 1, b: 2, c: -3}, want: false},
		{name: "case13", args: args{a: 0, b: 2, c: 3}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTriangle(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("isTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_accum for the function that is used in Mumbling.go
func Test_accum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "default", args: args{s: "ZpglnRxqenU"}, want: "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu"},
		{name: "case1", args: args{s: "NyffsGeyylB"}, want: "N-Yy-Fff-Ffff-Sssss-Gggggg-Eeeeeee-Yyyyyyyy-Yyyyyyyyy-Llllllllll-Bbbbbbbbbbb"},
		{name: "case2", args: args{s: "MjtkuBovqrU"}, want: "M-Jj-Ttt-Kkkk-Uuuuu-Bbbbbb-Ooooooo-Vvvvvvvv-Qqqqqqqqq-Rrrrrrrrrr-Uuuuuuuuuuu"},
		{name: "case3", args: args{s: "EvidjUnokmM"}, want: "E-Vv-Iii-Dddd-Jjjjj-Uuuuuu-Nnnnnnn-Oooooooo-Kkkkkkkkk-Mmmmmmmmmm-Mmmmmmmmmmm"},
		{name: "case4", args: args{s: "HbideVbxncC"}, want: "H-Bb-Iii-Dddd-Eeeee-Vvvvvv-Bbbbbbb-Xxxxxxxx-Nnnnnnnnn-Cccccccccc-Ccccccccccc"},
		{name: "case5", args: args{s: "VwhvtHtrxfE"}, want: "V-Ww-Hhh-Vvvv-Ttttt-Hhhhhh-Ttttttt-Rrrrrrrr-Xxxxxxxxx-Ffffffffff-Eeeeeeeeeee"},
		{name: "case6", args: args{s: "KurgiKmkphY"}, want: "K-Uu-Rrr-Gggg-Iiiii-Kkkkkk-Mmmmmmm-Kkkkkkkk-Ppppppppp-Hhhhhhhhhh-Yyyyyyyyyyy"},
		{name: "case7", args: args{s: "NctlfBlnmfH"}, want: "N-Cc-Ttt-Llll-Fffff-Bbbbbb-Lllllll-Nnnnnnnn-Mmmmmmmmm-Ffffffffff-Hhhhhhhhhhh"},
		{name: "case8", args: args{s: "WegunHvbdmV"}, want: "W-Ee-Ggg-Uuuu-Nnnnn-Hhhhhh-Vvvvvvv-Bbbbbbbb-Ddddddddd-Mmmmmmmmmm-Vvvvvvvvvvv"},
		{name: "case9", args: args{s: "VoywwSpqidE"}, want: "V-Oo-Yyy-Wwww-Wwwww-Ssssss-Ppppppp-Qqqqqqqq-Iiiiiiiii-Dddddddddd-Eeeeeeeeeee"},
		{name: "case10", args: args{s: "VbaixFpxdcO"}, want: "V-Bb-Aaa-Iiii-Xxxxx-Ffffff-Ppppppp-Xxxxxxxx-Ddddddddd-Cccccccccc-Ooooooooooo"},
		{name: "case11", args: args{s: "OlyqvYwkuzF"}, want: "O-Ll-Yyy-Qqqq-Vvvvv-Yyyyyy-Wwwwwww-Kkkkkkkk-Uuuuuuuuu-Zzzzzzzzzz-Fffffffffff"},
		{name: "case12", args: args{s: "JrhfdMtchiH"}, want: "J-Rr-Hhh-Ffff-Ddddd-Mmmmmm-Ttttttt-Cccccccc-Hhhhhhhhh-Iiiiiiiiii-Hhhhhhhhhhh"},
		{name: "case13", args: args{s: "JiwpcSwslvW"}, want: "J-Ii-Www-Pppp-Ccccc-Ssssss-Wwwwwww-Ssssssss-Lllllllll-Vvvvvvvvvv-Wwwwwwwwwww"},
		{name: "case14", args: args{s: "EagpiEvmabJ"}, want: "E-Aa-Ggg-Pppp-Iiiii-Eeeeee-Vvvvvvv-Mmmmmmmm-Aaaaaaaaa-Bbbbbbbbbb-Jjjjjjjjjjj"},
		{name: "case15", args: args{s: "RznlcEmuxxP"}, want: "R-Zz-Nnn-Llll-Ccccc-Eeeeee-Mmmmmmm-Uuuuuuuu-Xxxxxxxxx-Xxxxxxxxxx-Ppppppppppp"},
		{name: "case16", args: args{s: "OrggaExarzP"}, want: "O-Rr-Ggg-Gggg-Aaaaa-Eeeeee-Xxxxxxx-Aaaaaaaa-Rrrrrrrrr-Zzzzzzzzzz-Ppppppppppp"},
		{name: "case17", args: args{s: "DriraMtedfB"}, want: "D-Rr-Iii-Rrrr-Aaaaa-Mmmmmm-Ttttttt-Eeeeeeee-Ddddddddd-Ffffffffff-Bbbbbbbbbbb"},
		{name: "case18", args: args{s: "BjxseRxgtjT"}, want: "B-Jj-Xxx-Ssss-Eeeee-Rrrrrr-Xxxxxxx-Gggggggg-Ttttttttt-Jjjjjjjjjj-Ttttttttttt"},
		{name: "case19", args: args{s: "EquhxOswchE"}, want: "E-Qq-Uuu-Hhhh-Xxxxx-Oooooo-Sssssss-Wwwwwwww-Ccccccccc-Hhhhhhhhhh-Eeeeeeeeeee"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := accum(tt.args.s); got != tt.want {
				t.Errorf("accum() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_hasUniqueChar for the function that is used in AllUnique.go
func Test_hasUniqueChar(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "default", args: args{str: "  nAa"}, want: false},
		{name: "case1", args: args{str: "abcde"}, want: true},
		{name: "case2", args: args{str: "++-"}, want: false},
		{name: "case3", args: args{str: "AaBbC"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasUniqueChar(tt.args.str); got != tt.want {
				t.Errorf("hasUniqueChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_twoOldestAges for the function that is used in TwoOldestAges.go
func Test_twoOldestAges(t *testing.T) {
	type args struct {
		ages []int
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		// TODO: Add test cases.
		{name: "default", args: args{ages: []int{1, 5, 87, 45, 8, 8}}, want: [2]int{45, 87}},
		{name: "case1", args: args{ages: []int{6, 5, 83, 5, 3, 18}}, want: [2]int{18, 83}},
		{name: "case2", args: args{ages: []int{39, 53, 83, 51, 59, 61, 95, 23, 99, 49}}, want: [2]int{95, 99}},
		{name: "case3", args: args{ages: []int{43, 25, 83, 11, 31, 91, 85, 25, 95, 65}}, want: [2]int{91, 95}},
		{name: "case4", args: args{ages: []int{61, 11, 33, 79, 81, 27, 79, 83, 9, 95}}, want: [2]int{83, 95}},
		{name: "case5", args: args{ages: []int{63, 69, 15, 21, 81, 77, 85, 15, 19, 31}}, want: [2]int{81, 85}},
		{name: "case6", args: args{ages: []int{27, 13, 25, 93, 9, 65, 69, 45, 83, 87}}, want: [2]int{87, 93}},
		{name: "case7", args: args{ages: []int{93, 35, 53, 67, 17, 23, 89, 75, 15, 53}}, want: [2]int{89, 93}},
		{name: "case8", args: args{ages: []int{69, 67, 59, 45, 59, 37, 65, 39, 85, 21}}, want: [2]int{69, 85}},
		{name: "case9", args: args{ages: []int{87, 35, 43, 45, 45, 71, 15, 1, 91, 25}}, want: [2]int{87, 91}},
		{name: "case10", args: args{ages: []int{39, 9, 3, 97, 37, 27, 71, 71, 67, 51}}, want: [2]int{71, 97}},
		{name: "case11", args: args{ages: []int{19, 5, 43, 13, 75, 89, 43, 89, 25, 49}}, want: [2]int{89, 89}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoOldestAges(tt.args.ages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoOldestAges() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_equableTriangle for the function that is used in CheckIfATriangle.go
func Test_equableTriangle(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "default", args: args{a: 5, b: 12, c: 13}, want: true},
		{name: "case1", args: args{a: 2, b: 3, c: 4}, want: false},
		{name: "case2", args: args{a: 6, b: 8, c: 10}, want: true},
		{name: "case3", args: args{a: 7, b: 15, c: 20}, want: true},
		{name: "case4", args: args{a: 17, b: 17, c: 30}, want: false},
		{name: "case5", args: args{a: 7, b: 10, c: 12}, want: false},
		{name: "case6", args: args{a: 6, b: 11, c: 12}, want: false},
		{name: "case7", args: args{a: 25, b: 25, c: 45}, want: false},
		{name: "case8", args: args{a: 13, b: 37, c: 30}, want: false},
		{name: "case9", args: args{a: 6, b: 25, c: 29}, want: true},
		{name: "case10", args: args{a: 10, b: 11, c: 18}, want: false},
		{name: "case11", args: args{a: 73, b: 9, c: 80}, want: false},
		{name: "case12", args: args{a: 12, b: 35, c: 37}, want: false},
		{name: "case1", args: args{a: 120, b: 109, c: 13}, want: false},
		{name: "case1", args: args{a: 9, b: 10, c: 17}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equableTriangle(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("equableTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_gps for the function that is used in SpeedControl.go
func Test_gps(t *testing.T) { //SpeedControl.go
	type args struct {
		s int
		x []float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "default", args: args{s: 15, x: []float64{0.0, 0.19, 0.5, 0.75, 1.0, 1.25, 1.5, 1.75, 2.0, 2.25}}, want: 74},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gps(tt.args.s, tt.args.x); got != tt.want {
				t.Errorf("gps() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_twoToOne for the function that is used in twoToOne.go
func Test_twoToOne(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "default", args: args{s1: "xyaabbbccccdefww", s2: "xxxxyyyyabklmopq"}, want: "abcdefklmopqwxy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoToOne(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("twoToOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
