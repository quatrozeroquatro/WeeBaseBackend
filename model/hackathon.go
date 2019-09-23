package model

import "log"

type Hackathon struct {
	ID         int64  `gorm:"column:hackathon_id" json:"id,omitempty"`
	Image      []byte `gorm:"column:image" json:"image,omitempty"`
	Title      string `gorm:"column:title" json:"title,omitempty"`
	Onboarding string `gorm:"column:onboarding" json:"onboarding,omitempty"`
}

//CreateHackathon: criar um hackathon
func (dsd *WeeHackDB) CreateHackathon(hackathon *Hackathon) error {
	result := dsd.Db.Table("public.hackathon").Create(hackathon)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//GetScouts: retorna um hackathon
func (dsd *WeeHackDB) GetHackathon(id int) (*Hackathon, error) {
	hackathon := Hackathon{}
	result := dsd.Db.Table("public.hackathon").Where("id = ?", id).First(&hackathon)
	if result.Error != nil {
		log.Println("error on get data from hackathon", result.Error)
		return nil, result.Error
	}
	return &hackathon, nil
}

//GetUsers: retorna todos os hackathon
func (dsd *WeeHackDB) GetAllHackathons() (*[]Hackathon, error) {
	var hackathons []Hackathon
	result := dsd.Db.Table("public.hackathon").Find(&hackathons)
	if result.Error != nil {
		log.Println("error on get data from hackathon", result.Error)
		return nil, result.Error
	}
	return &hackathons, nil
}