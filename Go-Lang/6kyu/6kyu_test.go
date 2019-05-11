package main

import "testing"

//Test_multiple3And5 for the function that is used in MultipleOf3Or5.go
func Test_multiple3And5(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "default", args: args{number: 10}, want: 23},
		{name: "case1", args: args{number: 20}, want: 78},
		{name: "case2", args: args{number: 0}, want: 0},
		{name: "case3", args: args{number: 1}, want: 0},
		{name: "case4", args: args{number: 200}, want: 9168},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiple3And5(tt.args.number); got != tt.want {
				t.Errorf("multiple3And5() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_decodeMorse for the function that is used in DecodeTheMorseCode.go
func Test_decodeMorse(t *testing.T) {
	type args struct {
		morseCode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "default", args: args{morseCode: ".... . -.--   .--- ..- -.. ."}, want: "HEY JUDE"},
		{name: "case1", args: args{morseCode: ".-"}, want: "A"},
		{name: "case2", args: args{morseCode: "."}, want: "E"},
		{name: "case3", args: args{morseCode: ".."}, want: "I"},
		{name: "case4", args: args{morseCode: ". ."}, want: "EE"},
		{name: "case5", args: args{morseCode: ".   ."}, want: "E E"},
		{name: "case6", args: args{morseCode: "...---..."}, want: "SOS"},
		{name: "case7", args: args{morseCode: "... --- ..."}, want: "SOS"},
		{name: "case8", args: args{morseCode: "...   ---   ..."}, want: "S O S"},
		{name: "case9", args: args{morseCode: " . "}, want: "E"},
		{name: "case10", args: args{morseCode: "   .   . "}, want: "E E"},
		{name: "case11", args: args{morseCode: "      ...---... -.-.--   - .... .   --.- ..- .. -.-. -.-   -... .-. --- .-- -.   ..-. --- -..-   .--- ..- -- .--. ...   --- ...- . .-.   - .... .   .-.. .- --.. -.--   -.. --- --. .-.-.-  "}, want: "SOS! THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeMorse(tt.args.morseCode); got != tt.want {
				t.Errorf("decodeMorse() = %v, want %v", got, tt.want)
			}
		})
	}
}
