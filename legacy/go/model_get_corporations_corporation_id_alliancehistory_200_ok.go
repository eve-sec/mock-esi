package esilegacy

import "time"

/*
200 ok object */
type GetCorporationsCorporationIdAlliancehistory200Ok struct {
	/* */
	Alliance GetCorporationsCorporationIdAlliancehistoryAlliance `json:"alliance,omitempty"`
	/*
	 An incrementing ID that can be used to canonically establish order of records in cases where dates may be ambiguous */
	RecordId int32 `json:"record_id,omitempty"`
	/*
	 start_date string */
	StartDate time.Time `json:"start_date,omitempty"`
}
