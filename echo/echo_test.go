package echo

import (
	"io"
	"strings"
	"testing"
)

type ReaderWriterMock struct {
	data          []byte
	expectedError error
}

func (r *ReaderWriterMock) Read(p []byte) (n int, err error) {
	if r.expectedError != nil {
		return 0, r.expectedError
	}
	if strings.Compare(string(p[:len(r.data)]), string(r.data)) == 0 {
		return 0, io.EOF
	}
	copy(p, r.data)
	return len(r.data), nil
}

func (r *ReaderWriterMock) Write(p []byte) (n int, err error) {
	copy(r.data, p)
	return len(r.data), nil
}

func NewReaderWriterMock(message []byte, err error) *ReaderWriterMock {
	return &ReaderWriterMock{
		data:          message,
		expectedError: err,
	}
}

func TestDo(t *testing.T) {
	type args struct {
		rwc io.ReadWriter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "Copy text message",
			args: args{
				rwc: NewReaderWriterMock([]byte("TextMessage"), nil),
			},
			wantErr: false,
			want:    "TextMessage",
		},
		{
			name: "Passing in nil",
			args: args{
				rwc: NewReaderWriterMock(nil, io.ErrUnexpectedEOF),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Do(tt.args.rwc); (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				out := make([]byte, len(tt.want))
				n, err := tt.args.rwc.Read(out)
				if err != nil {
					t.Fatalf("Failed when writing %v", err)
				}
				if strings.Compare(string(out[:n]), tt.want) != 0 {
					t.Fatalf("Failed when comparing the output %v", tt.want)
				}
			}
		})
	}
}
