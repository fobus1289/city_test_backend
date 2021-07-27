package service

import (
	"gorm.io/gorm"
	"time"
	"v2/app/model"
)

type Branch struct {
	*gorm.DB
}

func NewBranchService(DB *gorm.DB) *Branch {
	return &Branch{DB: DB}
}

func (b *Branch) Branches(page, limit int, scopes ...func(tx *gorm.DB) *gorm.DB) (*model.Response, error) {

	response := model.NewResponse("branches", page, limit, &model.Branches{})

	if err := Paginate(b.DB, response, scopes...); err != nil {
		return nil, err
	}

	return response, nil
}

func (b *Branch) FindById(scopes ...func(tx *gorm.DB) *gorm.DB) (*model.Branch, error) {
	var branch *model.Branch

	if err := b.Table("branches").Scopes(scopes...).
		Find(&branch).Error; err != nil {
		return nil, err
	}

	return branch, nil
}

func (b *Branch) FindByUserId(scopes ...func(tx *gorm.DB) *gorm.DB) (model.Branches, error) {
	var branches model.Branches

	if err := b.Table("branches").Scopes(scopes...).Find(&branches).Error; err != nil {
		return nil, err
	}

	return branches, nil
}

func (b *Branch) FindByName(scopes ...func(tx *gorm.DB) *gorm.DB) (model.Branches, error) {
	var branches model.Branches

	if err := b.Table("branches").Scopes(scopes...).
		Find(&branches).Error; err != nil {
		return nil, err
	}

	return branches, nil
}

func (b *Branch) FindByCompanyId(scopes ...func(tx *gorm.DB) *gorm.DB) (model.Branches, error) {
	var branches model.Branches

	if err := b.Table("branches").Scopes(scopes...).
		Find(&branches).Error; err != nil {
		return nil, err
	}

	return branches, nil
}

func (b *Branch) Create(branch *model.Branch) (*model.Branch, error) {

	branch.FixUtilDate()

	if err := b.Table("branches").Omit("Id").
		Create(&branch).Error; err != nil {
		return nil, err
	}

	return branch, nil
}

func (b *Branch) Update(branch *model.Branch, scopes ...func(tx *gorm.DB) *gorm.DB) (*model.Branch, error) {

	branch.FixUtilDate()

	if err := b.Table("branches").Scopes(scopes...).Where("id=?", branch.Id).
		Omit("Id").Updates(&branch).Error; err != nil {
		return nil, err
	}
	return branch, nil
}

func (b *Branch) Delete(scopes ...func(tx *gorm.DB) *gorm.DB) error {
	return b.Table("branches").Scopes(scopes...).Update("deleted_at", time.Now()).Error
}

func (b *Branch) Exists(scopes ...func(tx *gorm.DB) *gorm.DB) (bool, error) {
	var id int64

	if err := b.Table("branches").Scopes(scopes...).Select("id").Scan(&id).Error; err != nil {
		return false, err
	}

	return id != 0, nil
}
