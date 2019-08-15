package models

import (
	"fmt"

	u "../utils"
	"github.com/jinzhu/gorm"
)

//Meetup is a stuct
type Meetup struct {
	gorm.Model
	Title        string `json:"title"`
	SubTitle     string `json:"subtitle"`
	Description  string `json:"description"`
	ImageURL     string `json:"imageUrl"`
	Address      string `json:"address"`
	ContactEmail string `json:"contactEmail"`
	IsLiked      bool   `json:"isLiked"`
}

// id: 1,
// title: "Coding bootcamp",
// subtitle: "learn to code quickly",
// description: "talking about coding",
// imageUrl:
// 	"https://iceclog.com/wp-content/uploads/2016/09/596px-Internet1.jpg",
// address: "123 fake street",
// contactEmail: "fake@1.com",
// isLiked: false

func isEmpty(values []string) bool {
	for _, value := range values {
		if value == "" {
			return false
		}
	}
	return true
}

//Validate validates the meetup struc
func (meetup *Meetup) Validate() (map[string]interface{}, bool) {

	if meetup.Title == "" || meetup.SubTitle == "" || meetup.Description == "" || meetup.ImageURL == "" || meetup.Address == "" || meetup.ContactEmail == "" {
		return u.Message(false, "Please enter all fields"), false
	}
	// if contact.Phone == "" {
	// 	return u.Message(false, "Please enter contact number"), false
	// }

	// if contact.UserId <= 0 {
	// 	return u.Message(false, "Please login"), false
	// }

	return u.Message(true, "success"), true
}

//DeleteMeetup deltes meetup based on id
func DeleteMeetup(id uint) map[string]interface{} {

	GetDB().Table("meetups").Where("id = ?", id).Delete(id)

	res := u.Message(true, "success")

	return res
}

//ToggleLike updates meetup like based on id
func ToggleLike(id uint) map[string]interface{} {
	meetup := &Meetup{}
	err := GetDB().Table("meetups").Where("id = ?", id).First(meetup).Error
	if err != nil {
		return nil
	}

	if meetup.IsLiked == true {
		GetDB().Table("meetups").Where("id = ?", id).Update("is_liked", false)
	} else {
		GetDB().Table("meetups").Where("id = ?", id).Update("is_liked", true)
	}

	res := u.Message(true, "success")

	return res
}

// Create Creates meetup based on data sent via api
func (meetup *Meetup) Create() map[string]interface{} {
	if res, ok := meetup.Validate(); !ok {
		return res
	}

	GetDB().Create(meetup)

	res := u.Message(true, "success")
	res["meetup"] = meetup
	return res
}

// UpdateMeetup updates meetup based on data sent via api
func (meetup *Meetup) UpdateMeetup() map[string]interface{} {

	if res, ok := meetup.Validate(); !ok {
		return res
	}
	fmt.Println(meetup)
	GetDB().First(&meetup)

	res := u.Message(true, "success")
	res["meetup"] = meetup
	return res
}

//GetMeetup gets one meetup based on get request with id
func GetMeetup(id uint) *Meetup {
	meetup := &Meetup{}
	err := GetDB().Table("meetups").Where("id = ?", id).First(meetup).Error
	if err != nil {
		return nil
	}

	return meetup
}

//GetMeetups gets all the meetups
func GetMeetups() []*Meetup {
	meetups := make([]*Meetup, 0)
	err := GetDB().Table("meetups").Order("ID DESC").Find(&meetups).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return meetups
}
