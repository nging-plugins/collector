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
	"encoding/json"

	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/echo"

	"github.com/coscms/webcore/library/backend"
	"github.com/coscms/webcore/library/common"

	"github.com/nging-plugins/collector/application/dbschema"
	"github.com/nging-plugins/collector/application/model"
)

func History(c echo.Context) error {
	pageID := c.Queryx(`pageId`).Uint()
	parentID := c.Queryx(`parentId`).Uint64()
	m := model.NewCollectorHistory(c)
	cond := db.NewCompounds()
	if parentID > 0 {
		err := m.Get(nil, `id`, parentID)
		if err != nil {
			return err
		}
		pageID = m.PageId
		cond.Add(db.Cond{`parent_id`: parentID})
	} else {
		if pageID > 0 {
			cond.Add(db.Cond{`page_id`: pageID})
		}
		cond.Add(db.Cond{`parent_id`: 0})
	}
	_, err := common.PagingWithLister(c, common.NewLister(m, nil, func(r db.Result) db.Result {
		return r.OrderBy(`-id`)
	}, cond))
	ret := common.Err(c, err)
	c.Set(`listData`, m.Objects())
	var positions []dbschema.NgingCollectorHistory
	var pageRule *dbschema.NgingCollectorPage
	if pageID > 0 {
		pageM := model.NewCollectorPage(c)
		pageM.Get(func(r db.Result) db.Result {
			return r.Select(`id`, `type`, `name`, `content_type`)
		}, `id`, pageID)
		if pageM.Id > 0 {
			pageRule = pageM.NgingCollectorPage
		}
		mw := func(r db.Result) db.Result {
			return r.Select(`page_id`, `title`, `parent_id`, `id`)
		}
		if m.Id < 1 {
			m.Get(mw, `page_id`, pageID)
		}
		if m.Id > 0 {
			positions, _ = m.Positions(mw, m.Id)
		}
	}
	c.Set(`positions`, positions)
	c.Set(`pageRule`, pageRule)
	return c.Render(`collector/history`, ret)
}

func HistoryView(c echo.Context) error {
	ident := c.Form(`ident`)
	data := c.Data()
	if !com.IsAlphaNumericUnderscore(ident) {
		return c.JSON(data.SetInfo(c.T(`无效参数`), 0))
	}
	b, err := common.ReadCache(`colloctor`, ident+`.json`)
	if err != nil {
		return c.JSON(data.SetError(err))
	}
	data.SetData(json.RawMessage(b))
	return c.JSON(data)
}

func HistoryDelete(c echo.Context) error {
	id := c.Form(`id`)
	m := model.NewCollectorHistory(c)
	cond := db.Cond{`id`: id}
	err := m.Delete(nil, cond)
	if err != nil {
		common.SendFail(c, err.Error())
	} else {
		common.SendOk(c, c.T(`操作成功`))
	}
	return c.Redirect(backend.URLFor(`/collector/history`))
}
