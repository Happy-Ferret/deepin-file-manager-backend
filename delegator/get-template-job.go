package delegator

import (
	"net/url"
	"pkg.linuxdeepin.com/lib/dbus"
	"pkg.linuxdeepin.com/lib/operations"
)

var _GetTemplateJobCount uint64

type GetTemplateJob struct {
	dbusInfo dbus.DBusInfo
	op       *operations.GetTemplateJob
}

// GetDBusInfo returns dbus information.
func (job *GetTemplateJob) GetDBusInfo() dbus.DBusInfo {
	return job.dbusInfo
}

// NewListJob creates a new list job for dbus.
func NewGetTemplateJob(templateDirURI *url.URL) *GetTemplateJob {
	job := &GetTemplateJob{
		dbusInfo: genDBusInfo("GetTemplateJob", &_GetTemplateJobCount),
		op:       operations.NewGetTemplateJob(templateDirURI),
	}
	return job
}

func (job *GetTemplateJob) Execute() []string {
	defer dbus.UnInstallObject(job)
	return job.op.Execute()
}