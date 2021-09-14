package models

type School struct {
	IDCity            string      `json:"idCity"`
	Name              string      `json:"name"`
	CompanyFancy      string      `json:"companyFancy"`
	TaxIDRegistration interface{} `json:"taxIdRegistration"`
	Phone             string      `json:"phone"`
	Active            bool        `json:"active"`
	CanReschedule     bool        `json:"canReschedule"`
	ReschedulingFee   float64     `json:"reschedulingFee"`
	Image             string      `json:"image"`
	Rating            float64     `json:"rating"`
}
