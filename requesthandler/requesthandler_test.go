package requesthandler

import (
	"sync"
	"testing"
)

func TestMakeHttpRequest(t *testing.T) {
	var wg1 sync.WaitGroup
	wg1.Add(2)
	type args struct {
		wg      *sync.WaitGroup
		urlChan <-chan string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Test",
			args: args{
				wg:      &wg1,
				urlChan: generateStrings(),
			},
			wantErr: false,
		},
		{
			name: "Negative Test",
			args: args{
				wg:      &wg1,
				urlChan: generateStringsNegative(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeHttpRequest(tt.args.wg, tt.args.urlChan); (err != nil) != tt.wantErr {
				t.Errorf("MakeHttpRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func generateStrings() <-chan string {
	rc := make(chan string)
	arr := []string{"google.com", "facebook.com", "github.com", "twitter.com"}
	go func() {
		defer close(rc)

		for i := 0; i < len(arr); i++ {
			rc <- arr[i]
		}
	}()
	return rc
}

func generateStringsNegative() <-chan string {
	rc := make(chan string)
	arr := []string{"googl"}
	go func() {
		defer close(rc)

		for i := 0; i < len(arr); i++ {
			rc <- arr[i]
		}
	}()
	return rc
}
