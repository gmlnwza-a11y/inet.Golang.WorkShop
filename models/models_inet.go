package models

import "gorm.io/gorm"

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

type RegisterRequest struct {
	Email        string `json:"email" validate:"required,email"`
	Username     string `json:"username" validate:"required,username_valid,min=3,max=20"`
	Password     string `json:"password" validate:"required,min=6,max=20"`
	LindeID      string `json:"line_id" validate:"omitempty"`
	PhoneNumber  string `json:"phone_number" validate:"required,numeric,len=10,startswith=0"`
	BusinessType string `json:"business_type" validate:"required,business_allowed"`
	WebURL       string `json:"web_url" validate:"subdomain"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type ResultData struct {
	Data       []DogsRes `json:"data"`
	Name       string    `json:"name"`
	Count      int       `json:"count"`
	SumRed     int       `json:"sum_red"`
	SumGreen   int       `json:"sum_green"`
	SumPink    int       `json:"sum_pink"`
	SumNoColor int       `json:"sum_nocolor"`
}

type Company struct {
	gorm.Model
	Name        string `json:"name"`
	CEO         string `json:"ceo"`
	Employee    int    `json:"employee"`
	Location    string `json:"location"`
	Established int    `json:"established"`
}

type CompanyRes struct {
	Name        string `json:"name"`
	CEO         string `json:"ceo"`
	Employee    int    `json:"employee"`
	Location    string `json:"location"`
	Established int    `json:"established"`
	Type        string `json:"type"`
}

type CompanyResultData struct {
	Data  []CompanyRes `json:"data"`
	Name  string       `json:"name"`
	Count int          `json:"count"`
}
