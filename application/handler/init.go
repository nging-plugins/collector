/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package handler

import (
	"github.com/coscms/webcore/library/module"
	routeRegistry "github.com/coscms/webcore/registry/route"
	"github.com/webx-top/echo"
)

func RegisterRoute(r module.Router) {
	r.Backend().RegisterToGroup(`/collector`, registerRoute)
}

func registerRoute(g echo.RouteRegister) {
	metaHandler := routeRegistry.MetaHandler
	g.Route(`GET,POST`, `/export`, metaHandler(Export, echo.H{`name`: `导出管理`}))
	g.Route(`GET,POST`, `/export_log`, metaHandler(ExportLog, echo.H{`name`: `日子列表`}))
	g.Route(`GET,POST`, `/export_log_view/:id`, metaHandler(ExportLogView, echo.H{`name`: `日志详情`}))
	g.Route(`GET,POST`, `/export_log_delete`, metaHandler(ExportLogDelete, echo.H{`name`: `删除日志`}))
	g.Route(`GET,POST`, `/export_add`, metaHandler(ExportAdd, echo.H{`name`: `添加导出规则`}))
	g.Route(`GET,POST`, `/export_edit`, metaHandler(ExportEdit, echo.H{`name`: `修改导出规则`}))
	g.Route(`GET,POST`, `/export_edit_status`, metaHandler(ExportEditStatus, echo.H{`name`: `修改导出规则`}))
	g.Route(`GET,POST`, `/export_delete`, metaHandler(ExportDelete, echo.H{`name`: `删除导出规则`}))
	g.Route(`GET,POST`, `/history`, metaHandler(History, echo.H{`name`: `历史记录`}))
	g.Route(`GET,POST`, `/history_view`, metaHandler(HistoryView, echo.H{`name`: `查看历史内容`}))
	g.Route(`GET,POST`, `/history_delete`, metaHandler(HistoryDelete, echo.H{`name`: `删除历史记录`}))
	g.Route(`GET,POST`, `/rule`, metaHandler(Rule, echo.H{`name`: `规则列表`}))
	g.Route(`GET,POST`, `/rule_add`, metaHandler(RuleAdd, echo.H{`name`: `添加规则`}))
	g.Route(`GET,POST`, `/rule_edit`, metaHandler(RuleEdit, echo.H{`name`: `修改规则`}))
	g.Route(`GET,POST`, `/rule_delete`, metaHandler(RuleDelete, echo.H{`name`: `删除规则`}))
	g.Route(`GET,POST`, `/rule_collect`, metaHandler(RuleCollect, echo.H{`name`: `采集`}))
	g.Route(`GET,POST`, `/group`, metaHandler(Group, echo.H{`name`: `任务分组列表`}))
	g.Route(`GET,POST`, `/group_add`, metaHandler(GroupAdd, echo.H{`name`: `添加分组`}))
	g.Route(`GET,POST`, `/group_edit`, metaHandler(GroupEdit, echo.H{`name`: `修改分组`}))
	g.Route(`GET,POST`, `/group_delete`, metaHandler(GroupDelete, echo.H{`name`: `删除分组`}))
	g.Route(`GET,POST`, `/regexp_test`, metaHandler(RegexpTest, echo.H{`name`: `测试正则表达式`}))
}
