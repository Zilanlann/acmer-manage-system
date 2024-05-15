package site_service

import "github.com/zilanlann/acmer-manage-system/server/model"

type SiteList struct {
	Sites []model.Site `json:"list"`
	Total int          `json:"total"`
}

func (s *SiteList) Get() (err error) {
	s.Sites, err = model.GetAllSiteList()
	s.Total = len(s.Sites)
	return
}

type SiteTypeList struct {
	SiteTypes []model.SiteType `json:"list"`
	Total     int              `json:"total"`
}

func (s *SiteTypeList) Get() (err error) {
	s.SiteTypes, err = model.GetAllSiteTypeList()
	s.Total = len(s.SiteTypes)
	return
}

type Site struct {
	model.Site
}

func (s *Site) Create() (err error) {
	err = s.Site.Create()
	return
}

func (s *Site) Update() (err error) {
	err = s.Site.Update()
	return
}

func (s *Site) Delete() (err error) {
	err = s.Site.Delete()
	return
}

type SiteType struct {
	model.SiteType
}

func (s *SiteType) Create() (err error) {
	err = s.SiteType.Create()
	return
}

func (s *SiteType) Update() (err error) {
	err = s.SiteType.Update()
	return
}

func (s *SiteType) Delete() (err error) {
	err = s.SiteType.Delete()
	return
}
