package onmyoji

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRefreshOnmyojiMedia(t *testing.T) {
	type args struct {
		existedFileMap map[string]bool
	}
	type testConfig struct {
		args    args
		wantErr bool
	}
	Convey("test", t, func() {
		//your mock code...
		Convey("test case1", func() {
			tt := testConfig{
				args: args{
					existedFileMap: map[string]bool{},
				},
				wantErr: false,
			}
			err := RefreshOnmyojiMedia(tt.args.existedFileMap, "/Users/bytedance/Documents/personal/yys/horizontal")
			t.Logf("err: %+v", err)

			So(err != nil, ShouldEqual, tt.wantErr)
		})
	})
}
