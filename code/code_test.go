package code

import "testing"

func TestCode_Equal(t *testing.T) {
	type args struct {
		v Code
	}
	tests := []struct {
		name string
		c    Code
		args args
		want bool
	}{
		{
			name: "Code Empty equal NilCode",
			c:    "",
			args: args{
				v: NilCode,
			},
			want: true,
		},
		{
			name: "Code(0000) equal Code(0000)",
			c:    "0000",
			args: args{
				v: Code("0000"),
			},
			want: true,
		},
		{
			name: "Code Empty not equal Code(0000)",
			c:    "",
			args: args{
				v: Code("0000"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Equal(tt.args.v); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCode_IsNil(t *testing.T) {
	tests := []struct {
		name string
		c    Code
		want bool
	}{
		{
			name: "Code Empty is NilCode",
			c:    "",
			want: true,
		},
		{
			name: "Code(0000) is not NilCode",
			c:    "0000",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsNil(); got != tt.want {
				t.Errorf("IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

type input struct {
	code Code
}

func (i input) Code() Code {
	return i.code
}

func TestForm(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want Code
	}{
		{
			name: "Struct with Code",
			args: args{
				v: &input{
					code: "0000",
				},
			},
			want: Code("0000"),
		},
		{
			name: "Struct with NilCode",
			args: args{
				v: &input{
					code: NilCode,
				},
			},
			want: NilCode,
		},
		{
			name: "Variable without Code get NilCode",
			args: args{
				v: 0,
			},
			want: NilCode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := From(tt.args.v); got != tt.want {
				t.Errorf(`From() = "%v", want "%v"`, got, tt.want)
			}
		})
	}
}
