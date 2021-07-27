package service

import (
	"gorm.io/gorm"
	"time"
	"v2/app/model"
)

type Company struct {
	*gorm.DB
}

func NewCompanyService(db *gorm.DB) *Company {
	return &Company{db}
}

func (c *Company) Companies(page, limit int, scopes ...func(tx *gorm.DB) *gorm.DB) (*model.Response, error) {

	response := model.NewResponse("companies", page, limit, &model.Companies{})

	if err := Paginate(c.DB, response, scopes...); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Company) FindById(scopes ...func(tx *gorm.DB) *gorm.DB) (*model.Company, error) {
	var company *model.Company

	if err := c.Table("companies").Scopes(scopes...).
		Find(&company).Error; err != nil {
		return nil, err
	}

	return company, nil
}

func (c *Company) FindByName(scopes ...func(tx *gorm.DB) *gorm.DB) (model.Companies, error) {
	var companies model.Companies

	if err := c.Table("companies").Scopes(scopes...).
		Find(&companies).Error; err != nil {
		return nil, err
	}

	return companies, nil
}

func (c *Company) FindByINN(scopes ...func(tx *gorm.DB) *gorm.DB) (*model.Company, error) {
	var company *model.Company

	if err := c.Table("companies").Scopes(scopes...).
		Find(&company).Error; err != nil {
		return nil, err
	}

	return company, nil
}

func (c *Company) FindByUserId(scopes ...func(tx *gorm.DB) *gorm.DB) (model.Companies, error) {
	var companies model.Companies

	if err := c.Table("companies").Scopes(scopes...).
		Find(&companies).Error; err != nil {
		return nil, err
	}

	return companies, nil
}

func (c *Company) FindByBranchId(scopes ...func(tx *gorm.DB) *gorm.DB) (*model.Company, error) {

	var (
		company *model.Company
		branch  *model.Branch
		tx      = c.DB.Begin()
		err     error
	)

	commit(tx, func() error {
		return err
	})

	if err = tx.Table("branches").Scopes(scopes...).Find(&branch).Error; err != nil {
		return nil, err
	}

	if err = tx.Table("companies").Where("id=?", branch.CompanyId).
		Find(&company).Error; err != nil {
		return nil, err
	}

	return company, nil
}

func (c *Company) Create(company *model.Company) (*model.Company, error) {
	if err := c.Table("companies").Omit("Id").
		Create(&company).Error; err != nil {
		return nil, err
	}
	return company, nil
}

func (c *Company) Update(company *model.Company, scopes ...func(tx *gorm.DB) *gorm.DB) (*model.Company, error) {
	if err := c.Table("companies").Scopes(scopes...).Omit("Id").
		Updates(&company).Error; err != nil {
		return nil, err
	}
	return company, nil
}

func (c *Company) Delete(scopes ...func(tx *gorm.DB) *gorm.DB) error {
	return c.Table("companies").Scopes(scopes...).Update("deleted_at", time.Now()).Error
}

func (c *Company) Exists(scopes ...func(tx *gorm.DB) *gorm.DB) (bool, error) {
	var id int64

	if err := c.Table("companies").Scopes(scopes...).Select("id").Scan(&id).Error; err != nil {
		return false, err
	}

	return id != 0, nil
}
