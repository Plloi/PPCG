package dailycodingproblem

import "testing"

func Test_problem253(t *testing.T) {
	type args struct {
		InputString string
		k           int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Base",
			args: args{
				InputString: "thisisazigzag",
				k: 4,
			},
			want: "meep",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem253(tt.args.InputString, tt.args.k); got != tt.want {
				t.Errorf("problem253() = %v, want %v", got, tt.want)
			}
		})
	}
}
