package sol

import "testing"

func TestIsNumberPalindrome(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "131",
			args: args{input: 131},
			want: true,
		},
		{
			name: "123",
			args: args{input: 123},
			want: false,
		},
		{
			name: "1221",
			args: args{input: 1221},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumberPalindrome(tt.args.input); got != tt.want {
				t.Errorf("IsNumberPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
