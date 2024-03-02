package types

type MemberResponse struct {
	ID               int    `json:"id"`
	CompoundMemberID string `json:"compoundMemberId"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	HomeGymID        int    `json:"homeGymId"`
	HomeGymName      string `json:"homeGymName"`
	EmailAddress     string `json:"emailAddress"`
	GymAccessPIN     string `json:"gymAccessPin"`
	DateOfBirth      string `json:"dateofBirth"`
	MobileNumber     string `json:"mobileNumber"`
	PostCode         string `json:"postCode"`
	MembershipName   string `json:"membershipName"`
	MembershipLevel  int    `json:"membershipLevel"`
	SuspendedReason  int    `json:"suspendedReason"`
	MemberStatus     int    `json:"memberStatus"`
}
