package models

type Specs struct {
	MaxSpeed           float64 `json:"maxspeed"`
	MaxSteeringFore    float64 `json:"maxsteeringforce"`
	MaxAngularVelocity float64 `json:"maxangularvelocity"`
	VisionRadius       float64 `json:"visionradius"`
	VisionAngle        float64 `json:"visionangle"`
	DragForce          float64 `json:"dragforce"`
}
