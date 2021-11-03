package echo

import (
	"io"
	"strings"
	"testing"
)

type ReaderWriterCloserMock struct {
	data          []byte
	expectedError error
}

func (r *ReaderWriterCloserMock) Read(p []byte) (n int, err error) {
	if r.expectedError != nil {
		return 0, r.expectedError
	}
	if strings.Compare(string(p[:len(r.data)]), string(r.data)) == 0 {
		return 0, io.EOF
	}
	copy(p, r.data)
	return len(r.data), nil
}

func (r *ReaderWriterCloserMock) Write(p []byte) (n int, err error) {
	copy(r.data, p)
	return len(r.data), nil
}

func (r *ReaderWriterCloserMock) Close() error {
	return nil
}

func NewReaderWriterCloserMock(message []byte, err error) *ReaderWriterCloserMock {
	return &ReaderWriterCloserMock{
		data:          message,
		expectedError: err,
	}
}

func TestDo(t *testing.T) {
	type args struct {
		rwc io.ReadWriteCloser
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
				rwc: NewReaderWriterCloserMock([]byte("TextMessage"), nil),
			},
			wantErr: false,
			want:    "TextMessage",
		},
		{
			name: "Passing in nil",
			args: args{
				rwc: NewReaderWriterCloserMock(nil, io.ErrUnexpectedEOF),
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
