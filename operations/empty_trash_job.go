/**
 * Copyright (C) 2015 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package operations

import (
	"gir/gio-2.0"
	"pkg.deepin.io/dde/api/soundutils"
	"strings"
)

// EmptyTrashJob is a job to empty trash.
type EmptyTrashJob struct {
	*CommonJob
	trashDirs     []*gio.File
	shouldConfirm bool
	// doneCallback OpCallback
	// doneCallbackData interface{}
}

func (job *EmptyTrashJob) init(uiDelegate IUIDelegate) {
	// TODO: inhibit power manager

	job.CommonJob = newCommon(uiDelegate)
	job.trashDirs = []*gio.File{gio.FileNewForUri("trash:")}
	job.progressUnit = AmountUnitSumOfFilesAndDirs
}

func (job *EmptyTrashJob) finalize() {
	for _, trashDir := range job.trashDirs {
		trashDir.Unref()
	}

	job.CommonJob.finalize()
}

// delete files and directories on trash directory.
//
// @param delFile delete the file or directory, it always be true except for trashDir.
// @param delChildren delete children of directories. If a file is deleted, false should be passed.
func (job *EmptyTrashJob) deleteTrashFile(target *gio.File, delFile bool, delChildren bool) {
	if job.isAborted() {
		return
	}

	if delChildren {
		enumerator, _ := target.EnumerateChildren(strings.Join(
			[]string{
				gio.FileAttributeStandardName,
				gio.FileAttributeStandardType,
			}, ","),
			gio.FileQueryInfoFlagsNofollowSymlinks,
			job.cancellable,
		)
		if enumerator != nil {
			for !job.isAborted() {
				info, err := enumerator.NextFile(job.cancellable)
				if info == nil || err != nil {
					break
				}

				child := target.GetChild(info.GetName())
				job.deleteTrashFile(child, true, info.GetFileType() == gio.FileTypeDirectory)

				info.Unref()
				child.Unref()
			}

			enumerator.Close(job.cancellable)
			enumerator.Unref()
		}
	}

	if !job.isAborted() && delFile {
		target.Delete(job.cancellable)
	}
}

// Execute EmptyTrashJob and finalize resources.
func (job *EmptyTrashJob) Execute() {
	defer finishJob(job)

	soundutils.PlaySystemSound(soundutils.EventEmptyTrash, "", false)

	confirmed := true
	if job.shouldConfirm {
		// TODO: docs
		confirmed = job.uiDelegate.AskDeleteConfirmation(Tr("empty???"), Tr(""), Tr(""))
	}

	if confirmed {
		for _, trashDir := range job.trashDirs {
			job.deleteTrashFile(trashDir, false, true)
		}
	}

}

// NewEmptyTrashJob creates a new empty trash job.
func NewEmptyTrashJob(shouldConfirm bool, uiDelegate IUIDelegate) *EmptyTrashJob {
	job := &EmptyTrashJob{
		shouldConfirm: shouldConfirm,
	}

	job.init(uiDelegate)

	return job
}
