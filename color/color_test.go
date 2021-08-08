package color

import (
	"testing"
)

func TestBlack(t *testing.T) {
	type args struct {
		msg  string
		conf []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestBlack1", args: args{msg: "Black"}, want: "\033[0;0;30mBlack\033[0m"},
		{name: "TestBlack underline", args: args{msg: "Black", conf: []int{TypeUnderline}}, want: "\033[4;0;30mBlack\033[0m"},
		{name: "TestBlack underline and black background", args: args{msg: "Black", conf: []int{TypeUnderline, WHITE + 10}}, want: "\033[4;47;30mBlack\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Black(tt.args.msg, tt.args.conf...); got != tt.want {
				t.Errorf("Black() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlue(t *testing.T) {
	type args struct {
		msg  string
		conf []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestBlue1", args: args{msg: "Blue"}, want: "\033[0;0;34mBlue\033[0m"},
		{name: "TestBlue underline", args: args{msg: "Blue", conf: []int{TypeUnderline}}, want: "\033[4;0;34mBlue\033[0m"},
		{name: "TestBlue underline and black background", args: args{msg: "Blue", conf: []int{TypeUnderline, BLACK + 10}}, want: "\033[4;40;34mBlue\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Blue(tt.args.msg, tt.args.conf...); got != tt.want {
				t.Errorf("Blue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCyan(t *testing.T) {
	type args struct {
		msg  string
		conf []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestCyan1", args: args{msg: "Cyan"}, want: "\033[0;0;36mCyan\033[0m"},
		{name: "TestCyan underline", args: args{msg: "Cyan", conf: []int{TypeUnderline}}, want: "\033[4;0;36mCyan\033[0m"},
		{name: "TestCyan underline and black background", args: args{msg: "Cyan", conf: []int{TypeUnderline, BLACK + 10}}, want: "\033[4;40;36mCyan\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cyan(tt.args.msg, tt.args.conf...); got != tt.want {
				t.Errorf("Cyan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreen(t *testing.T) {
	type args struct {
		msg  string
		conf []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestGreen1", args: args{msg: "Green"}, want: "\033[0;0;32mGreen\033[0m"},
		{name: "TestGreen underline", args: args{msg: "Green", conf: []int{TypeUnderline}}, want: "\033[4;0;32mGreen\033[0m"},
		{name: "TestGreen underline and black background", args: args{msg: "Green", conf: []int{TypeUnderline, BLACK + 10}}, want: "\033[4;40;32mGreen\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Green(tt.args.msg, tt.args.conf...); got != tt.want {
				t.Errorf("Green() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMagenta(t *testing.T) {
	type args struct {
		msg  string
		conf []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestMagenta1", args: args{msg: "Magenta"}, want: "\033[0;0;35mMagenta\033[0m"},
		{name: "TestMagenta underline", args: args{msg: "Magenta", conf: []int{TypeUnderline}}, want: "\033[4;0;35mMagenta\033[0m"},
		{name: "TestMagenta underline and black background", args: args{msg: "Magenta", conf: []int{TypeUnderline, BLACK + 10}}, want: "\033[4;40;35mMagenta\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Magenta(tt.args.msg, tt.args.conf...); got != tt.want {
				t.Errorf("Magenta() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRed(t *testing.T) {
	type args struct {
		msg  string
		conf []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestRed1", args: args{msg: "Red"}, want: "\033[0;0;31mRed\033[0m"},
		{name: "TestRed underline", args: args{msg: "Red", conf: []int{TypeUnderline}}, want: "\033[4;0;31mRed\033[0m"},
		{name: "TestRed underline and black background", args: args{msg: "Red", conf: []int{TypeUnderline, BLACK + 10}}, want: "\033[4;40;31mRed\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Red(tt.args.msg, tt.args.conf...); got != tt.want {
				t.Errorf("Red() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWhite(t *testing.T) {
	type args struct {
		msg  string
		conf []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestWhite1", args: args{msg: "White"}, want: "\033[0;0;37mWhite\033[0m"},
		{name: "TestWhite underline", args: args{msg: "White", conf: []int{TypeUnderline}}, want: "\033[4;0;37mWhite\033[0m"},
		{name: "TestWhite underline and black background", args: args{msg: "White", conf: []int{TypeUnderline, BLACK + 10}}, want: "\033[4;40;37mWhite\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := White(tt.args.msg, tt.args.conf...); got != tt.want {
				t.Errorf("White() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYellow(t *testing.T) {
	type args struct {
		msg  string
		conf []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestYellow1", args: args{msg: "Yellow"}, want: "\033[0;0;33mYellow\033[0m"},
		{name: "TestYellow underline", args: args{msg: "Yellow", conf: []int{TypeUnderline}}, want: "\033[4;0;33mYellow\033[0m"},
		{name: "TestYellow underline and black background", args: args{msg: "Yellow", conf: []int{TypeUnderline, BLACK + 10}}, want: "\033[4;40;33mYellow\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Yellow(tt.args.msg, tt.args.conf...); got != tt.want {
				t.Errorf("Yellow() = %v, want %v", got, tt.want)
			}
		})
	}
}
