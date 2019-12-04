package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"math"
	u "github.com/realr3fo/tkai_circles/utils"
)

type Circle struct {
	gorm.Model
	Radius float32 `json:"radius"`
	Area   float32 `json:"area"`
	Owner  string  `json:"owner"`
}

type Tube struct {
	gorm.Model
	CircleID int     `json:"circleId"`
	Height   float32 `json:"height"`
	Volume   float32 `json:"volume"`
	Owner    string  `json:"owner"`
}

type Ball struct {
	gorm.Model
	CircleID int     `json:"circleId"`
	Volume   float32 `json:"volume"`
	Owner    string  `json:"owner"`
}

func (circle *Circle) GetCircleArea(ownerID uint) (map[string]interface{}, error) {
	if circle.Radius == 0 {
		err := errors.New("circle radius cannot be empty")
		return nil, err
	}

	radius := circle.Radius
	area := math.Pi * radius * radius

	circle.Area = area

	accounts := make([]*Account, 0)
	err := GetDB().Table("accounts").Where("id = ?", ownerID).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	ownerAccount := accounts[0]

	circle.Owner = ownerAccount.Username

	GetDB().Create(circle)

	resp := u.Message(true, "success")
	resp["circle"] = circle
	return resp, nil

}

func (tube *Tube) GetTubeVolume(ownerID uint) (map[string]interface{}, error) {
	if tube.CircleID == 0 {
		err := errors.New("circle id cannot be empty")
		return nil, err
	}
	if tube.Height == 0 {
		err := errors.New("tube height cannot be empty")
		return nil, err
	}

	circleID := tube.CircleID
	circles := make([]*Circle, 0)
	{
		err := GetDB().Table("circles").Where("id = ?", circleID).Find(&circles).Error
		if err != nil {
			return nil, err
		}
	}
	circleObj := circles[0]
	tubeVolume := circleObj.Area * tube.Height
	tube.Volume = tubeVolume

	accounts := make([]*Account, 0)
	err := GetDB().Table("accounts").Where("id = ?", ownerID).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	ownerAccount := accounts[0]
	tube.Owner = ownerAccount.Username

	GetDB().Create(tube)

	resp := u.Message(true, "success")
	resp["tube"] = tube
	return resp, nil
}

func (ball *Ball) GetBallVolume(ownerID uint) (map[string]interface{}, error) {
	if ball.CircleID == 0 {
		err := errors.New("circle id cannot be empty")
		return nil, err
	}

	circleID := ball.CircleID
	circles := make([]*Circle, 0)
	{
		err := GetDB().Table("circles").Where("id = ?", circleID).Find(&circles).Error
		if err != nil {
			return nil, err
		}
	}
	circleObj := circles[0]
	circleArea := circleObj.Area
	circleRadius := circleObj.Radius

	ballVolume := float32(4)/float32(3) * circleArea * circleRadius
	ball.Volume = ballVolume

	accounts := make([]*Account, 0)
	err := GetDB().Table("accounts").Where("id = ?", ownerID).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	ownerAccount := accounts[0]
	ball.Owner = ownerAccount.Username

	GetDB().Create(ball)

	resp := u.Message(true, "success")
	resp["ball"] = ball
	return resp, nil
}
