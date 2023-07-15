package datatype

type Camera struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	RoverId  int    `json:"rover_id"`
	FullName string `json:"full_name"`
}

type CameraSimple struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

type Rover struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	LandingDate string         `json:"landing_date"`
	LaunchDate  string         `json:"launch_date"`
	Status      string         `json:"status"`
	MaxSol      int            `json:"max_sol"`
	MaxDate     string         `json:"max_date"`
	TotalPhotos int            `json:"total_photos"`
	Cameras     []CameraSimple `json:"cameras"`
}

type Photo struct {
	Id        int    `json:"id"`
	Sol       int    `json:"sol"`
	Camera    Camera `json:"camera"`
	ImgSrc    string `json:"img_src"`
	EarthDate string `json:"earth_date"`
	Rover     Rover  `json:"rover"`
}

type Data struct {
	Data []Photo `json:"photos"`
}
