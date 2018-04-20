package esiv1

/*
200 ok object */
type GetCorporationsCorporationIdContacts200Ok struct {
	/*
	 Standing of the contact */
	Standing float32 `json:"standing,omitempty"`
	/*
	 contact_type string */
	ContactType string `json:"contact_type,omitempty"`
	/*
	 contact_id integer */
	ContactId int32 `json:"contact_id,omitempty"`
	/*
	 Whether this contact is being watched */
	IsWatched bool `json:"is_watched,omitempty"`
	/*
	 Custom label of the contact */
	LabelId int64 `json:"label_id,omitempty"`
}