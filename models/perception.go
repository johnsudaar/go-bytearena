package models

type Perception struct {
	Vision     []*Item `json:"vision"`
	Energy     float64 `json:"energy"`
	Velocity   Vector2 `json:"velocity"`
	Azimuth    float64 `json:"azimuth"`
	BodyRadius float64 `json:"bodyradius"`
}

type Item struct {
	Tag      string  `json:"tag"`
	NearEdge Vector2 `json:"nearedge"`
	Center   Vector2 `json:"center"`
	FarEdge  Vector2 `json:"faredge"`
	Velocity Vector2 `json:"velocity"`
}
