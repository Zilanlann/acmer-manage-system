package model

import (
	"github.com/zilanlann/acmer-manage-system/server/global"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Site struct {
	gorm.Model
	Host       string   `gorm:"not null;unique" json:"host" binding:"required"`
	Name       string   `gorm:"not null;unique" json:"name" binding:"required"`
	SiteTypeID uint     `gorm:"not null" json:"siteTypeID" binding:"required"`
	SiteType   SiteType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"siteType"`
	Desc       string   `gorm:"type:text" json:"desc"`
}

type SiteType struct {
	gorm.Model
	Name string `gorm:"not null;unique" json:"name" binding:"required"`
	Desc string `gorm:"type:text" json:"desc"`
}

// Create inserts a new Site record into the database.
func (s *Site) Create() error {
	return global.DB.Create(&s).Error
}

// Update modifies an existing Site record in the database.
func (s *Site) Update() error {
	return global.DB.Updates(&s).Error
}

// Delete removes a Site record from the database, including associations.
func (s *Site) Delete() error {
	return global.DB.Select(clause.Associations).Unscoped().Delete(&s).Error
}

// Get retrieves a Site record from the database by its ID, including its associated SiteType.
func (s *Site) Get() error {
	return global.DB.Preload("SiteType").Where("id = ?", s.ID).First(&s).Error
}

// GetAllSiteList retrieves all Site records from the database, including their associated SiteTypes.
func GetAllSiteList() ([]Site, error) {
	var sites []Site
	err := global.DB.Preload("SiteType").Find(&sites).Error
	return sites, err
}

// Create inserts a new SiteType record into the database.
func (st *SiteType) Create() error {
	return global.DB.Create(&st).Error
}

// Update modifies an existing SiteType record in the database.
func (st *SiteType) Update() error {
	return global.DB.Updates(&st).Error
}

// Delete removes a SiteType record from the database, including associations.
func (st *SiteType) Delete() error {
	return global.DB.Select(clause.Associations).Unscoped().Delete(&st).Error
}

// Get retrieves a SiteType record from the database by its ID.
func (st *SiteType) Get() error {
	return global.DB.Where("id = ?", st.ID).First(&st).Error
}

// GetAllSiteTypeList retrieves all SiteType records from the database.
func GetAllSiteTypeList() ([]SiteType, error) {
	var siteTypes []SiteType
	err := global.DB.Find(&siteTypes).Error
	return siteTypes, err
}
