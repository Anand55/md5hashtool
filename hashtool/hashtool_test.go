package hashtool

import (
	"flag"
	"os"
	"testing"
)

func TestStart(t *testing.T) {
	cases := []struct {
		name string
		args []string
	}{
		{
			name: "Args with parallel flag ",
			args: []string{"", "-parallel", "1", "google.com"},
		},
		{
			name: "Args without parallel flag",
			args: []string{"", "google.com", "facebook.com"},
		},
	}
	for _, tc := range cases {
		flag.CommandLine = flag.NewFlagSet(tc.name, flag.ExitOnError)
		os.Args = tc.args
		Start()
	}
}
