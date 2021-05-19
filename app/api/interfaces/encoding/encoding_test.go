package IDENT

import "testing"

func TestEncodeBase64(t *testing.T) {
	type args struct {
		savePath string
		fileNama string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeBase64(tt.args.savePath, tt.args.fileNama); got != tt.want {
				t.Errorf("EncodeBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}
