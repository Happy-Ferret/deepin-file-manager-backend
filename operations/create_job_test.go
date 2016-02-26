/**
 * Copyright (C) 2015 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package operations_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"os/exec"
	"path/filepath"
	. "pkg.deepin.io/service/file-manager-backend/operations"
	"testing"
)

func TestCreateDirectory(t *testing.T) {
	SkipConvey("create directory on /tmp", t, func() {
		destDir := "./testdata/create"
		So(NewCreateDirectoryJob(destDir, "", skipMock).Execute(), ShouldBeNil)
		So(NewCreateDirectoryJob(destDir, "", skipMock).Execute(), ShouldBeNil)
		// TODO: filter keep_me
		files, _ := filepath.Glob(destDir + "/*")
		exec.Command("rmdir", files...).Run()
	})
}

func TestCreateFile(t *testing.T) {
	SkipConvey("create a file without a specific name", t, func() {
		destDir := "./testdata/create"
		So(NewCreateFileJob(destDir, "", []byte{}, skipMock).Execute(), ShouldBeNil)
		So(NewCreateFileJob(destDir, "", []byte{}, skipMock).Execute(), ShouldBeNil)
		files, _ := filepath.Glob(destDir + "/*")
		exec.Command("rm", files...).Run()
	})

	SkipConvey("create a file with a specific name", t, func() {
		destDir := "./testdata/create"
		So(NewCreateFileJob(destDir, "xxxxx", []byte{}, skipMock).Execute(), ShouldBeNil)
		So(NewCreateFileJob(destDir, "xxxxx", []byte{}, skipMock).Execute(), ShouldBeNil)
		files, _ := filepath.Glob(destDir + "/xxxxx*")
		exec.Command("rm", files...).Run()
	})

	SkipConvey("create a file with some init content", t, func() {
		destDir := "./testdata/create"
		So(NewCreateFileJob(destDir, "xxxxx", []byte("xxxxxxx"), skipMock).Execute(), ShouldBeNil)
		files, _ := filepath.Glob(destDir + "/xxxxx*")
		exec.Command("rm", files...).Run()
	})
}

func TestCreateFileFromTemplate(t *testing.T) {
	SkipConvey("create a file from template", t, func() {
		destDir := "./testdata/create"
		templateURL := "/home/liliqiang/Templates/newPowerPoint.ppt"
		So(NewCreateFileFromTemplateJob(destDir, templateURL, skipMock).Execute(), ShouldBeNil)
		So(NewCreateFileFromTemplateJob(destDir, templateURL, skipMock).Execute(), ShouldBeNil)
		files, _ := filepath.Glob(destDir + "/*")
		exec.Command("rm", files...).Run()
	})
}
