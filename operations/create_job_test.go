package operations_test

import (
	. "deepin-file-manager/operations"
	. "github.com/smartystreets/goconvey/convey"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestCreateDirectory(t *testing.T) {
	Convey("create directory on /tmp", t, func() {
		destDir := "./testdata/create"
		destURL, _ := pathToURL(destDir)
		So(NewCreateDirectoryJob(destURL, "", skipMock).Execute(), ShouldBeNil)
		So(NewCreateDirectoryJob(destURL, "", skipMock).Execute(), ShouldBeNil)
		files, _ := filepath.Glob(destDir + "/*")
		exec.Command("rmdir", files...).Run()
	})
}

func TestCreateFile(t *testing.T) {
	Convey("create a file without a specific name", t, func() {
		destDir := "./testdata/create"
		urlDir, _ := pathToURL(destDir)
		So(NewCreateFileJob(urlDir, "", []byte{}, skipMock).Execute(), ShouldBeNil)
		So(NewCreateFileJob(urlDir, "", []byte{}, skipMock).Execute(), ShouldBeNil)
		files, _ := filepath.Glob(destDir + "/*")
		exec.Command("rm", files...).Run()
	})

	Convey("create a file with a specific name", t, func() {
		destDir := "./testdata/create"
		urlDir, _ := pathToURL(destDir)
		So(NewCreateFileJob(urlDir, "xxxxx", []byte{}, skipMock).Execute(), ShouldBeNil)
		So(NewCreateFileJob(urlDir, "xxxxx", []byte{}, skipMock).Execute(), ShouldBeNil)
		files, _ := filepath.Glob(destDir + "/xxxxx*")
		exec.Command("rm", files...).Run()
	})

	Convey("create a file with some init content", t, func() {
		destDir := "./testdata/create"
		urlDir, _ := pathToURL(destDir)
		So(NewCreateFileJob(urlDir, "xxxxx", []byte("xxxxxxx"), skipMock).Execute(), ShouldBeNil)
		files, _ := filepath.Glob(destDir + "/xxxxx*")
		exec.Command("rm", files...).Run()
	})
}

func TestCreateFileFromTemplate(t *testing.T) {
	Convey("create a file from template", t, func() {
		destDir := "./testdata/create"
		destURL, _ := pathToURL(destDir)
		templateURL, _ := pathToURL("/home/liliqiang/Templates/newPowerPoint.ppt")
		So(NewCreateFileFromTemplateJob(destURL, templateURL, skipMock).Execute(), ShouldBeNil)
		So(NewCreateFileFromTemplateJob(destURL, templateURL, skipMock).Execute(), ShouldBeNil)
		files, _ := filepath.Glob(destDir + "/*")
		exec.Command("rm", files...).Run()
	})
}
