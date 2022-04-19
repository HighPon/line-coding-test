package main

import "testing"

func TestCheckMidnightTime(t *testing.T) {
	type args struct {
		recordTime string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{recordTime: "13:50:13.100"},
			want: false,
		},
		{
			name: "case 2",
			args: args{recordTime: "00:50:13.100"},
			want: true,
		},
		{
			name: "case 3",
			args: args{recordTime: "04:50:13.100"},
			want: true,
		},
		{
			name: "case 4",
			args: args{recordTime: "08:50:13.100"},
			want: false,
		},
		{
			name: "case 5",
			args: args{recordTime: "22:50:13.100"},
			want: true,
		},
		{
			name: "case 6",
			args: args{recordTime: "23:50:13.100"},
			want: true,
		},
		{
			name: "case 7",
			args: args{recordTime: "24:50:13.100"},
			want: true,
		},
		{
			name: "case 8",
			args: args{recordTime: "28:50:13.100"},
			want: true,
		},
		{
			name: "case 9",
			args: args{recordTime: "30:50:13.100"},
			want: false,
		},
		{
			name: "case 10",
			args: args{recordTime: "46:50:13.100"},
			want: true,
		},
		{
			name: "case 11",
			args: args{recordTime: "47:50:13.100"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkMidnightTime(tt.args.recordTime); got != tt.want {
				t.Errorf("checkMidnightTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcTimeFromZeroWithoutDecimal(t *testing.T) {
	type args struct {
		timeString []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{timeString: []string{"13", "50", "08"}},
			want: 49808,
		},
		{
			name: "case 2",
			args: args{timeString: []string{"3", "53", "51"}},
			want: 14031,
		},
		{
			name: "case 3",
			args: args{timeString: []string{"8", "21", "48"}},
			want: 30108,
		},
		{
			name: "case 4",
			args: args{timeString: []string{"18", "08", "59"}},
			want: 65339,
		},
		{
			name: "case 5",
			args: args{timeString: []string{"23", "34", "12"}},
			want: 84852,
		},
		{
			name: "case 6",
			args: args{timeString: []string{"25", "42", "48"}},
			want: 92568,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcTimeFromZeroWithoutDecimal(tt.args.timeString); got != tt.want {
				t.Errorf("calcTimeFromZeroWithoutDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestCalcTimeBetweenTwo(t *testing.T) {
// 	type args struct {
// 		firstDistance  string
// 		secondDistance string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			name: "case 1",
// 			args: args{timeString: ""},
// 			want: 49808,
// 		},
// 		{
// 			name: "case 2",
// 			args: args{timeString: []string{"3", "53", "51"}},
// 			want: 14031,
// 		},
// 		{
// 			name: "case 3",
// 			args: args{timeString: []string{"8", "21", "48"}},
// 			want: 30108,
// 		},
// 		{
// 			name: "case 4",
// 			args: args{timeString: []string{"18", "08", "59"}},
// 			want: 65339,
// 		},
// 		{
// 			name: "case 5",
// 			args: args{timeString: []string{"23", "34", "12"}},
// 			want: 84852,
// 		},
// 		{
// 			name: "case 6",
// 			args: args{timeString: []string{"25", "42", "48"}},
// 			want: 92568,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := calcTimeFromZeroWithoutDecimal(tt.args.timeString); got != tt.want {
// 				t.Errorf("calcTimeFromZeroWithoutDecimal() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
