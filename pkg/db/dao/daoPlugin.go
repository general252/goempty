package dao

import (
	"fmt"
	model "github.com/general252/goempty/pkg/db"
	"github.com/jinzhu/gorm"
)

type Plugin struct {
	tab model.MPlugin
}

func (c *Plugin) Add(obj *model.MPlugin) (int64, error) {
	var db = model.GetDataBase()

	if err := db.Create(obj).Error; err != nil {
		return 0, err
	}

	return obj.Id, nil
}

func (c *Plugin) Has(name string) (bool, error) {
	var db = model.GetDataBase()

	var count int64
	err := db.Model(&c.tab).Where(&model.MPlugin{Name: name}).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count >= 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func (c *Plugin) Get(name string) (*model.MPlugin, error) {
	var db = model.GetDataBase()

	var obj = &model.MPlugin{}
	r := db.Model(&c.tab).Where(&model.MPlugin{Name: name}).Find(obj)
	if r.Error != nil {
		return nil, r.Error
	}
	if r.RowsAffected != 1 {
		return nil, fmt.Errorf("not found")
	}

	return obj, nil
}

func (c *Plugin) GetByKeyId(keyId int64) (*model.MPlugin, error) {
	var db = model.GetDataBase()

	var obj = &model.MPlugin{}
	r := db.Model(&c.tab).Where(&model.MPlugin{Id: keyId}).Find(obj)
	if r.Error != nil {
		return nil, r.Error
	}
	if r.RowsAffected != 1 {
		return nil, fmt.Errorf("not found")
	}

	return obj, nil
}

func (c *Plugin) Del(keyId int64) error {
	var db = model.GetDataBase()

	r := db.Model(&c.tab).Delete(&model.MPlugin{Id: keyId})
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected != 1 {
	}

	return nil
}

func (c *Plugin) Update(keyId int64, obj *model.MPlugin) error {
	var db = model.GetDataBase()

	r := db.Model(&c.tab).Where(&model.MPlugin{Id: keyId}).Update(obj)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

type PluginFilter struct {
	Name        string `json:"name"`         // 名称
	Info        string `json:"info"`         // 简介
	InstallPath string `json:"install_path"` // 安装路径
	Match       string `json:"match"`        // 模糊匹配
}

func (c *Plugin) Query(limit int, offset int, filter *PluginFilter) ([]model.MPlugin, int64, error) {
	var db = model.GetDataBase()

	var result []model.MPlugin
	var filterSQL = func() *gorm.DB {
		r := db.Model(&c.tab)
		if len(filter.Name) > 0 {
			r = r.Where("name LIKE ?", fmt.Sprintf("%%%v%%", filter.Name))
		}
		if len(filter.Info) > 0 {
			r = r.Where("info LIKE ?", fmt.Sprintf("%%%v%%", filter.Info))
		}
		if len(filter.InstallPath) > 0 {
			r = r.Where("install_path LIKE ?", fmt.Sprintf("%%%v%%", filter.InstallPath))
		}
		if len(filter.Match) > 0 {
			r = r.Where("name LIKE ?", fmt.Sprintf("%%%v%%", filter.Match)).
				Or("info LIKE ?", fmt.Sprintf("%%%v%%", filter.Match)).
				Or("install_path LIKE ?", fmt.Sprintf("%%%v%%", filter.Match))
		}

		return r
	}

	rSQL := filterSQL()
	rSQL = rSQL.Limit(limit).Offset(offset).Find(&result)
	if rSQL.Error != nil {
		return nil, 0, rSQL.Error
	}

	var count int64
	rSQLCount := filterSQL()
	_ = rSQLCount.Count(&count)

	return result, count, nil
}
