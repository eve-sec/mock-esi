package esidev

/*
last_week object */
type GetFwLeaderboardsCorporationsLastWeek struct {
	/*
	 Amount of kills */
	Amount int32 `json:"amount,omitempty"`
	/*
	 corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`
}
