package esiv1

/*
200 ok object */
type GetCorporationsCorporationIdShareholders200Ok struct {
	/*
	 shareholder_id integer */
	ShareholderId int32 `json:"shareholder_id,omitempty"`
	/*
	 shareholder_type string */
	ShareholderType string `json:"shareholder_type,omitempty"`
	/*
	 share_count integer */
	ShareCount int64 `json:"share_count,omitempty"`
}