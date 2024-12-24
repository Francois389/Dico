package mot

import "testing"

func Test_sort(t *testing.T) {
	type args struct {
		mot string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test sortLetter with niche",
			args: args{mot: "niche"},
			want: "cehin",
		},
		{
			name: "Test sortLetter with chien",
			args: args{mot: "chien"},
			want: "cehin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortLetter(tt.args.mot); got != tt.want {
				t.Errorf("sortLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
