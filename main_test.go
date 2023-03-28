package main

import "testing"

func Test_optionallyTranslateToSingular(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "arg_pods", args: args{in: "pods"}, want: "pod"},
		{name: "arg_pod", args: args{in: "pod"}, want: "pod"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optionallyTranslateToSingular(tt.args.in); got != tt.want {
				t.Errorf("optionallyTranslateToSingular() = %v, want %v", got, tt.want)
			}
		})
	}
}
