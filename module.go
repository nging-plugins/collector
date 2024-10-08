package collector

import (
	"github.com/coscms/webcore/library/cron"
	"github.com/coscms/webcore/library/module"

	"github.com/nging-plugins/collector/application/handler"
	"github.com/nging-plugins/collector/application/library/setup"
)

const ID = `collector`

var Module = module.Module{
	TemplatePath: map[string]string{
		ID: `collector/template/backend`,
	},
	AssetsPath: []string{
		`collector/public/assets`,
	},
	SQLCollection: setup.RegisterSQL,
	Navigate:      RegisterNavigate,
	Route:         handler.RegisterRoute,
	CronJobs: []*cron.Jobx{
		{
			Name:         `collect_page`,
			RunnerGetter: handler.CollectPageJob,
			Example:      `>collect_page:1`,
			Description:  `网页采集`,
		},
	},
	DBSchemaVer: 0.2000,
}
