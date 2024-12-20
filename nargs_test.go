package nargs

import (
	"reflect"
	"testing"
)

func TestCheckForUnusedFunctionArgs(t *testing.T) {
	defaultFlags := Flags{
		IncludeNamedReturns: false,
		IncludeReceivers:    false,
		IncludeTests:        true,
		SetExitStatus:       true,
	}

	type args struct {
		cliArgs []string
		flags   Flags
	}
	tests := []struct {
		name               string
		args               args
		wantResults        []string
		wantExitWithStatus bool
		wantErr            bool
	}{
		{
			name: "Success (file with no errors), default flags",
			args: args{
				cliArgs: []string{"testdata/success.go"},
				flags:   defaultFlags,
			},
			// Even though setExitStatus is true, no errors were found.
			// Hence, we do not want to exit with a nonzero exit code.
			wantExitWithStatus: false,
			wantErr:            false,
		},
		{
			name: "File with errors, default flags",
			args: args{
				cliArgs: []string{"testdata/test.go"},
				flags:   defaultFlags,
			},
			wantResults: []string{
				"testdata/test.go:6 funcOne contains unused parameter c\n",
				"testdata/test.go:13 funcTwo contains unused parameter z\n",
				"testdata/test.go:31 closureOne contains unused parameter v\n",
				"testdata/test.go:39 unusedFunc contains unused parameter f\n",
				"testdata/test.go:43 closureTwo contains unused parameter i\n",
			},
			wantExitWithStatus: true,
			wantErr:            false,
		},
		{
			name: "File with errors, include named returns",
			args: args{
				cliArgs: []string{"testdata/test.go"},
				flags: Flags{
					IncludeNamedReturns: true,
					IncludeReceivers:    true,
					IncludeTests:        true,
					SetExitStatus:       true,
				},
			},
			wantResults: []string{
				"testdata/test.go:6 funcOne contains unused parameter c\n",
				"testdata/test.go:13 funcTwo contains unused parameter z\n",
				"testdata/test.go:19 funcThree contains unused parameter recv\n",
				"testdata/test.go:25 funcFour contains unused parameter namedReturn\n",
				"testdata/test.go:31 closureOne contains unused parameter v\n",
				"testdata/test.go:39 unusedFunc contains unused parameter f\n",
				"testdata/test.go:43 closureTwo contains unused parameter i\n",
			},
			wantExitWithStatus: true,
			wantErr:            false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResults, gotExitWithStatus, err := CheckForUnusedFunctionArgs(tt.args.cliArgs, tt.args.flags)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckForUnusedFunctionArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("CheckForUnusedFunctionArgs()\ngot = %v,\nwant %v", gotResults, tt.wantResults)
			}
			if gotExitWithStatus != tt.wantExitWithStatus {
				t.Errorf(
					"CheckForUnusedFunctionArgs() gotExitWithStatus = %v, want %v",
					gotExitWithStatus,
					tt.wantExitWithStatus,
				)
			}
		})
	}
}
