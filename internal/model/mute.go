package model

import "time"

type Mute struct {
	Username              string    `json:"username" bson:"username"`
	Muter                 string    `json:"muter" bson:"muter"`
	MuteUntil             time.Time `json:"muteUntil" bson:"mute_until"`
	Reason                string    `json:"reason" bson:"reason"`
	Message               string    `json:"message" bson:"message"`
	AdditionalInformation string    `json:"additionalInformation" bson:"additional_information"`
}

type Unmute struct {
	Username              string `json:"username" bson:"username"`
	Reason                string `json:"reason" bson:"reason"`
	Unmuter               string `json:"unmuter" bson:"unmuter"`
	AdditionalInformation string `json:"additionalInformation" bson:"additional_information"`
}
