package esiv2

/*
jump_clone object */
type GetCharactersCharacterIdClonesJumpClone struct {
	/*
	 implants array */
	Implants []int32 `json:"implants,omitempty"`
	/*
	 location_id integer */
	LocationId int64 `json:"location_id,omitempty"`
	/*
	 location_type string */
	LocationType string `json:"location_type,omitempty"`
}
