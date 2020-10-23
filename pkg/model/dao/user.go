package dao

import (
	"github.com/general252/goempty/pkg/model"
	"github.com/general252/gout/ulog"
	"gorm.io/gorm"
)

// NewUserDao user表
func NewUserDao() *userDao {
	return &userDao{}
}

type userDao struct {
	tab model.MUser
}

func (c *userDao) Add(obj *model.MUser) (int64, error) {
	var db = model.GetDataBase()

	if err := db.Create(obj).Error; err != nil {
		ulog.Error("%v", err)
		return 0, err
	}

	return obj.Id, nil
}

func (c *userDao) Has(name string) (bool, error) {
	var db = model.GetDataBase()

	var count int64
	err := db.Model(&c.tab).Where(&model.MUser{UserName: name}).Count(&count).Error
	if err != nil {
		ulog.Error("%v", err)
		return false, err
	}
	if count >= 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func (c *userDao) Get(name string) (*model.MUser, error) {
	var db = model.GetDataBase()

	var obj = &model.MUser{}
	r := db.Model(&c.tab).Where(&model.MUser{UserName: name}).Find(obj)
	if r.Error != nil {
		return nil, r.Error
	}

	if r.RowsAffected == 0 {
		return nil, nil
	}

	return obj, nil
}

func (c *userDao) GetByKeyId(keyId int64) (*model.MUser, error) {
	var db = model.GetDataBase()

	var obj = &model.MUser{}
	r := db.Model(&c.tab).Where(&model.MUser{Id: keyId}).Find(obj)
	if r.Error != nil {
		ulog.Error("%v", r.Error)
		return nil, r.Error
	}
	if r.RowsAffected == 0 {
		return nil, nil
	}

	return obj, nil
}

func (c *userDao) Del(keyId int64) error {
	var db = model.GetDataBase()

	r := db.Model(&c.tab).Delete(&model.MUser{Id: keyId})
	if r.Error != nil {
		ulog.Error("%v", r.Error)
		return r.Error
	}
	if r.RowsAffected != 1 {
	}

	return nil
}

func (c *userDao) Update(keyId int64, obj *model.MUser) error {
	var db = model.GetDataBase()

	r := db.Model(&c.tab).Where(&model.MUser{Id: keyId}).Updates(obj)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

type Filter struct {
	UserName string `json:"user_name"`
}

func (c *userDao) Query(limit int, offset int, filter *Filter) ([]model.MUser, int64, error) {
	var db = model.GetDataBase()

	var result []model.MUser
	var filterSQL = func() *gorm.DB {
		r := db.Model(&c.tab)
		if len(filter.UserName) > 0 {
			r = r.Where("name LIKE ?", "%"+filter.UserName+"%")
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
