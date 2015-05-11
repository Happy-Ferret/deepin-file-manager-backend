package operations_test

import (
	. "deepin-file-manager/operations"
	. "github.com/smartystreets/goconvey/convey"
	"os/exec"
	"pkg.linuxdeepin.com/lib/gio-2.0"
	"testing"
)

func TestChownJob(t *testing.T) {
	Convey("test chown with a file", t, func() {
		exec.Command("cp", "./testdata/chown/testfile", "./testdata/chown/afile").Run()
		defer exec.Command("rm", "./testdata/chown/afile").Run()
		u, err := pathToURL("./testdata/chown/afile")
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		// TODO: generate a not existed group.
		job := NewChownJob(u, "xx", "xx")
		job.Execute()
		So(job.HasError(), ShouldBeTrue)

		// TODO: make sure the original group is not targetGroup.
		// and permission denied won't happen.
		targetGroup := "video"
		job2 := NewChownJob(u, "", targetGroup)
		job2.Execute()
		So(job2.HasError(), ShouldBeFalse)

		f := gio.FileNewForPath("./testdata/chown/afile")
		info, _ := f.QueryInfo(gio.FileAttributeOwnerGroup, gio.FileQueryInfoFlagsNofollowSymlinks, nil)
		So(info.GetAttributeString(gio.FileAttributeOwnerGroup), ShouldEqual, targetGroup)
		info.Unref()

	})
}
