package types

type GymAttendanceResponse struct {
	Description                  string `json:"description"`
	TotalPeopleInGym             int    `json:"totalPeopleInGym"`
	TotalPeopleInClasses         int    `json:"totalPeopleInClasses"`
	TotalPeopleSuffix            string `json:"totalPeopleSuffix,omitempty"`
	IsApproximate                bool   `json:"isApproximate"`
	AttendanceTime               string `json:"attendanceTime"`
	LastRefreshed                string `json:"lastRefreshed"`
	LastRefreshedPeopleInClasses string `json:"lastRefreshedPeopleInClasses"`
	MaximumCapacity              int    `json:"maximumCapacity"`
}
