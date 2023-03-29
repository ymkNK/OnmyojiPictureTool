package cmd

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/cobra"
)

func Test_syncPic(t *testing.T) {
	type args struct {
		c    *cobra.Command
		args []string
	}
	type testConfig struct {
		args args
	}
	Convey("test", t, func() {
		//your mock code...
		Convey("test case1", func() {
			tt := testConfig{
				args: args{
					//c : &cobra.Command{},
					args: []string{},
				},
			}
			syncPic(tt.args.c, tt.args.args)
		})
	})
}
