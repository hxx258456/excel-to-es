package excel

import "testing"

func TestReadExcel(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test case",
			args: args{
				filepath: "../testdata/院校(1).xlsx",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//ReadExcel()
		})
	}
}
