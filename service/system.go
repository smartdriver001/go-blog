/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-05-07
 * Time: 22:12
 */
package service

import (
	"github.com/izghua/go-blog/common"
	"github.com/izghua/go-blog/conf"
	"github.com/izghua/go-blog/entity"
	"github.com/izghua/zgh"
)

func GetSystemList() (system *entity.ZSystems,err error) {
	system = new(entity.ZSystems)
	_,err = conf.SqlServer.Get(system)
	if err != nil {
		zgh.ZLog().Error("message","service.GetSystemList","err",err.Error())
		return
	}
	if system.Id <= 0 {
		systemInsert := entity.ZSystems{
			Theme:conf.Theme,
			Title: conf.Title,
			Keywords: conf.Keywords,
			Description: conf.Description,
			RecordNumber: conf.RecordNumber,
		}
		_,err = conf.SqlServer.Insert(systemInsert)
		if err != nil {
			zgh.ZLog().Error("message","service.GetSystemList","err",err.Error())
			return
		}
		_,err = conf.SqlServer.Get(system)
		if err != nil {
			zgh.ZLog().Error("message","service.GetSystemList","err",err.Error())
			return
		}
	}
	return
}

func SystemUpdate(sId int,ss common.ConsoleSystem) error {
	systemUpdate := entity.ZSystems{
		Title: ss.Title,
		Keywords: ss.Keywords,
		Description: ss.Description,
		RecordNumber: ss.RecordNumber,
		Theme: ss.Theme,
	}
	_,err := conf.SqlServer.Id(sId).Update(&systemUpdate)
	return err
}