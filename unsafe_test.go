//go:build !appengine

package fasttemplate

import (
	"reflect"
	"testing"
)

func Test_unsafeBytes2String(t *testing.T) {
	tests := []struct {
		name string
		args []byte
		want string
	}{
		{
			name: "regular bytes",
			args: []byte("string"),
			want: "string",
		},
		{
			name: "unicode",
			args: []byte("😉"),
			want: "😉",
		},
		{
			name: "empty",
			args: []byte{},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unsafeBytes2String(tt.args); got != tt.want {
				t.Errorf("unsafeBytes2String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unsafeString2Bytes(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []byte
	}{
		{
			name: "regular string",
			args: "string",
			want: []byte("string"),
		},
		{
			name: "unicode",
			args: "😉",
			want: []byte("😉"),
		},
		{
			name: "empty",
			args: "",
			want: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := unsafeString2Bytes(tt.args); !reflect.DeepEqual(gotB, tt.want) {
				t.Errorf("unsafeString2Bytes() = %v, want %v (is nil %t)", gotB, tt.want, gotB == nil)
			}
		})
	}
}
