package main

import (
	"reflect"
	"testing"
)

//Test_abbrevName for the function that is used in AbbreviateATwoWordName.go
func Test_abbrevName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "default", args: args{name: "Sam Harris"}, want: "S.H"},
		{name: "case1", args: args{name: "Patrick Feenan"}, want: "P.F"},
		{name: "case2", args: args{name: "Evan Cole"}, want: "E.C"},
		{name: "case3", args: args{name: "P Favuzzi"}, want: "P.F"},
		{name: "case4", args: args{name: "David Mendieta"}, want: "D.M"},
		{name: "case5", args: args{name: "sam harris"}, want: "S.H"},
		{name: "case6", args: args{name: "Patrick feenan"}, want: "P.F"},
		{name: "case6", args: args{name: "evan Cole"}, want: "E.C"},
		{name: "case6", args: args{name: "P F"}, want: "P.F"},
		{name: "case6", args: args{name: "DAvid mEndieta"}, want: "D.M"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abbrevName(tt.args.name); got != tt.want {
				t.Errorf("abbrevName() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_points for the function that is used in TotalAmountOfPoints.go
func Test_points(t *testing.T) {
	type args struct {
		games []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "default0", args: args{games: []string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}}, want: 30},
		{name: "default1", args: args{games: []string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"}}, want: 10},
		{name: "default2", args: args{games: []string{"0:1", "0:2", "0:3", "0:4", "1:2", "1:3", "1:4", "2:3", "2:4", "3:4"}}, want: 0},
		{name: "default3", args: args{games: []string{"1:0", "2:0", "3:0", "4:0", "2:1", "1:3", "1:4", "2:3", "2:4", "3:4"}}, want: 15},
		{name: "default4", args: args{games: []string{"1:0", "2:0", "3:0", "4:4", "2:2", "3:3", "1:4", "2:3", "2:4", "3:4"}}, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := points(tt.args.games); got != tt.want {
				t.Errorf("points() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_multiply for the function that is used in Multiply.go
func Test_multiply(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "default", args: args{a: 2, b: 3}, want: 6},
		{name: "case1", args: args{a: 10, b: 9}, want: 90},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_evenOrOdd for the function that is used in EvenOrOdd.go
func Test_evenOrOdd(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "default", args: args{number: 1}, want: "Odd"},
		{name: "case1", args: args{number: 2}, want: "Even"},
		{name: "case2", args: args{number: 0}, want: "Even"},
		{name: "case3", args: args{number: 42}, want: "Even"},
		{name: "case4", args: args{number: 13}, want: "Odd"},
		{name: "case5", args: args{number: 99}, want: "Odd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evenOrOdd(tt.args.number); got != tt.want {
				t.Errorf("evenOrOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_multipleOfIndex for the function that is used in MultipleOfIndex.go
func Test_multipleOfIndex(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "default", args: args{ints: []int{22, -6, 32, 82, 9, 25}}, want: []int{-6, 32, 25}},
		{name: "case1", args: args{ints: []int{68, -1, 1, -7, 10, 10}}, want: []int{-1, 10}},
		{name: "case2", args: args{ints: []int{11, -11}}, want: []int{-11}},
		{name: "case3", args: args{ints: []int{-56, -85, 72, -26, -14, 76, -27, 72, 35, -21, -67, 87, 0, 21, 59, 27, -92, 68}}, want: []int{-85, 72, 0, 68}},
		{name: "case4", args: args{ints: []int{28, 38, -44, -99, -13, -54, 77, -51}}, want: []int{38, -44, -99}},
		{name: "case5", args: args{ints: []int{-1, -49, -1, 67, 8, -60, 39, 35}}, want: []int{-49, 8, -60, 35}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multipleOfIndex(tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("multipleOfIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_makeNegative for the function that is used in ReturnNegative.go
func Test_makeNegative(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "default", args: args{x: 1}, want: -1},
		{name: "case1", args: args{x: -5}, want: -5},
		{name: "case2", args: args{x: 0}, want: 0},
		{name: "case3", args: args{x: 42}, want: -42},
		{name: "case4", args: args{x: -9}, want: -9},
		{name: "case5", args: args{x: -1}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeNegative(tt.args.x); got != tt.want {
				t.Errorf("makeNegative() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_repeatStr for the function that is used in StringRepeat.go
func Test_repeatStr(t *testing.T) {
	type args struct {
		repititions int
		value       string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "default", args: args{repititions: 4, value: "a"}, want: "aaaa"},
		{name: "case1", args: args{repititions: 3, value: "hello "}, want: "hello hello hello "},
		{name: "case2", args: args{repititions: 2, value: "abc"}, want: "abcabc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatStr(tt.args.repititions, tt.args.value); got != tt.want {
				t.Errorf("repeatStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_century for the function that is used in CenturyFromYear.go
func Test_century(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "default", args: args{year: 1705}, want: 18},
		{name: "case1", args: args{year: 1900}, want: 19},
		{name: "case2", args: args{year: 1601}, want: 17},
		{name: "case3", args: args{year: 2000}, want: 20},
		{name: "case4", args: args{year: 89}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := century(tt.args.year); got != tt.want {
				t.Errorf("century() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_opposite for the function that is used in OppositeNumber.go
func Test_opposite(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "default", args: args{value: 1}, want: -1},
		{name: "case1", args: args{value: 0}, want: 0},
		{name: "case2", args: args{value: 5}, want: -5},
		{name: "case3", args: args{value: -5}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := opposite(tt.args.value); got != tt.want {
				t.Errorf("opposite() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_positiveSum for the function that is used in SumOfPositive.go
func Test_positiveSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "default", args: args{numbers: []int{1, -4, 7, 12}}, want: 20},
		{name: "case1", args: args{numbers: []int{1, 2, 3, 4, 5}}, want: 15},
		{name: "case2", args: args{numbers: []int{1, -2, 3, 4, 5}}, want: 13},
		{name: "case3", args: args{numbers: []int{}}, want: 0},
		{name: "case4", args: args{numbers: []int{-1, -2, -3, -4, -5}}, want: 0},
		{name: "case5", args: args{numbers: []int{-1, 2, 3, 4, -5}}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := positiveSum(tt.args.numbers); got != tt.want {
				t.Errorf("positiveSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_greet for the function that is used in Function1.go
func Test_greet(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		if got := greet(); got != "hello world" {
			t.Errorf("greet() = %v, want %v", got, "hello world")
		}
	})
}

//Test_removeChar for the function that is used in RemoveFirstAndLastCharacter.go
func Test_removeChar(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "default", args: args{word: "eloquent"}, want: "loquen"},
		{name: "case1", args: args{word: "country"}, want: "ountr"},
		{name: "case2", args: args{word: "person"}, want: "erso"},
		{name: "case3", args: args{word: "place"}, want: "lac"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeChar(tt.args.word); got != tt.want {
				t.Errorf("removeChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_solution for the function that is used in ReversedStrings.go
func Test_solution(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "default", args: args{word: "world"}, want: "dlrow"},
		{name: "case1", args: args{word: "hello"}, want: "olleh"},
		{name: "case2", args: args{word: ""}, want: ""},
		{name: "case3", args: args{word: "h"}, want: "h"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.word); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_inAscOrder for the function that is used in AreTheNumbersInOrder.go
func Test_inAscOrder(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "default", args: args{numbers: []int{1, 2}}, want: true},
		{name: "case1", args: args{numbers: []int{-2, -1}}, want: true},
		{name: "case2", args: args{numbers: []int{45, 987}}, want: true},
		{name: "case3", args: args{numbers: []int{-1094, -45}}, want: true},
		{name: "case4", args: args{numbers: []int{2, 1}}, want: false},
		{name: "case5", args: args{numbers: []int{-1, -2}}, want: false},
		{name: "case6", args: args{numbers: []int{987, 45}}, want: false},
		{name: "case7", args: args{numbers: []int{-45, -1094}}, want: false},
		{name: "case8", args: args{numbers: []int{1, 2, 3}}, want: true},
		{name: "case9", args: args{numbers: []int{1, 3, 2}}, want: false},
		{name: "case10", args: args{numbers: []int{2, 1, 3}}, want: false},
		{name: "case11", args: args{numbers: []int{2, 3, 1}}, want: false},
		{name: "case12", args: args{numbers: []int{3, 1, 2}}, want: false},
		{name: "case13", args: args{numbers: []int{3, 2, 1}}, want: false},
		{name: "case14", args: args{numbers: []int{-666, -33, -10}}, want: true},
		{name: "case15", args: args{numbers: []int{-33, -10, -666}}, want: false},
		{name: "case16", args: args{numbers: []int{-10, -33, -666}}, want: false},
		{name: "case17", args: args{numbers: []int{-10, -666, -33}}, want: false},
		{name: "case18", args: args{numbers: []int{-33, -666, -10}}, want: false},
		{name: "case19", args: args{numbers: []int{-666, -10, -33}}, want: false},
		{name: "case20", args: args{numbers: []int{1, 4, 13, 97, 508, 1047, 20058}}, want: true},
		{name: "case21", args: args{numbers: []int{56, 98, 123, 67, 742, 1024, 32, 90969}}, want: false},
		{name: "case22", args: args{numbers: []int{1, 4, 13, 97, 508, 1047, 20058}}, want: true},
		{name: "case23", args: args{numbers: []int{56, 98, 123, 67, 742, 1024, 32, 9096}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inAscOrder(tt.args.numbers); got != tt.want {
				t.Errorf("inAscOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_boolToWord for the function that is used in ConvertBoolean.go
func Test_boolToWord(t *testing.T) {
	type args struct {
		word bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "yes", args: args{word: true}, want: "Yes"},
		{name: "no", args: args{word: false}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := boolToWord(tt.args.word); got != tt.want {
				t.Errorf("boolToWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_multiplyFunc3 for the function that is used in Function3.go
func Test_multiplyFunc3(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "default", args: args{a: 2, b: 2}, want: 4},
		{name: "case1", args: args{a: 4, b: 5}, want: 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplyFunc3(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("multiplyFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_maps for the function that is used in Beginner.go
func Test_maps(t *testing.T) {
	type args struct {
		x []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "default", args: args{x: []int{1, 2, 3}}, want: []int{2, 4, 6}},
		{name: "case1", args: args{x: []int{4, 1, 1, 1, 4}}, want: []int{8, 2, 2, 2, 8}},
		{name: "case2", args: args{x: []int{2, 2, 2, 2, 2, 2}}, want: []int{4, 4, 4, 4, 4, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maps(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maps() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_squareOrSquareRoot for the function that is used in ToSquare(root)OrNotToSquare(root).go
func Test_squareOrSquareRoot(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "default", args: args{arr: []int{4, 3, 9, 7, 2, 1}}, want: []int{2, 9, 3, 49, 4, 1}},
		{name: "case1", args: args{arr: []int{100, 101, 5, 5, 1, 1}}, want: []int{10, 10201, 25, 25, 1, 1}},
		{name: "case2", args: args{arr: []int{1, 2, 3, 4, 5, 6}}, want: []int{1, 4, 9, 2, 25, 36}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareOrSquareRoot(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("squareOrSquareRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_combat for the function that is used in Grasshopper.go
func Test_combat(t *testing.T) {
	type args struct {
		health float64
		damage float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{name: "default", args: args{health: 100.0, damage: 50.0}, want: 50.0},
		{name: "Dead", args: args{health: 1.0, damage: 100.0}, want: 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combat(tt.args.health, tt.args.damage); got != tt.want {
				t.Errorf("combat() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_howManyDalmatians for the function that is used in 101Dalmatians.go
func Test_howManyDalmatians(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "default", args: args{number: 26}, want: "More than a handful!"},
		{name: "case1", args: args{number: 8}, want: "Hardly any"},
		{name: "case2", args: args{number: 14}, want: "More than a handful!"},
		{name: "case3", args: args{number: 80}, want: "Woah that's a lot of dogs!"},
		{name: "case4", args: args{number: 100}, want: "Woah that's a lot of dogs!"},
		{name: "case5", args: args{number: 50}, want: "More than a handful!"},
		{name: "case6", args: args{number: 10}, want: "Hardly any"},
		{name: "case7", args: args{number: 101}, want: "101 DALMATIONS!!!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := howManyDalmatians(tt.args.number); got != tt.want {
				t.Errorf("howManyDalmatians() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_isPalindrome for the function that is used in IsItAPalindrome.go
func Test_isPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "default", args: args{str: "a"}, want: true},
		{name: "case1", args: args{str: "aba"}, want: true},
		{name: "case2", args: args{str: "abba"}, want: true},
		{name: "case3", args: args{str: "hello"}, want: false},
		{name: "case4", args: args{str: "Bob"}, want: true},
		{name: "case5", args: args{str: "Madam"}, want: true},
		{name: "case6", args: args{str: "AbBa"}, want: true},
		{name: "case7", args: args{str: ""}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.str); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_dnaToRNA for the function that is used in DNAToRNAConversion.go
func Test_dnaToRNA(t *testing.T) {
	type args struct {
		dna string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "default", args: args{dna: "GCAT"}, want: "GCAU"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dnaToRNA(tt.args.dna); got != tt.want {
				t.Errorf("dnaToRNA() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Test_hero for the function that is used in IsHeGonnaSurvive.go
func Test_hero(t *testing.T) {
	type args struct {
		bullets int
		dragons int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "default", args: args{bullets: 10, dragons: 5}, want: true},
		{name: "case1", args: args{bullets: 7, dragons: 4}, want: false},
		{name: "case2", args: args{bullets: 4, dragons: 5}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hero(tt.args.bullets, tt.args.dragons); got != tt.want {
				t.Errorf("hero() = %v, want %v", got, tt.want)
			}
		})
	}
}
