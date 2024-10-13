package gt7

import "gorm.io/gorm"

// Car represents a car entity with basic and detailed information.
type Car struct {
	gorm.Model
	Name         string `json:"name" bson:"name"`
	Brand        string `json:"brand" bson:"brand"`
	GrLevel      string `json:"gr_level" bson:"gr_level"`
	ImageURL     string `json:"image_url" bson:"image_url"`
	DetailURL    string `json:"detail_url" bson:"detail_url"`
	Weight       string `json:"weight" bson:"weight"`
	Performance  string `json:"performance" bson:"performance"`
	Dimensions   string `json:"dimensions" bson:"dimensions"`
	Description  string `json:"description" bson:"description"`
	Manufacturer string `json:"manufacturer" bson:"manufacturer"`
	Category     string `json:"category" bson:"category"`
	Engine       string `json:"engine" bson:"engine"`
	Aspiration   string `json:"aspiration" bson:"aspiration"`
	Power        string `json:"power" bson:"power"`
	Torque       string `json:"torque" bson:"torque"`
	Drivetrain   string `json:"drivetrain" bson:"drivetrain"`
	Length       string `json:"length" bson:"length"`
	Width        string `json:"width" bson:"width"`
	Height       string `json:"height" bson:"height"`
}
