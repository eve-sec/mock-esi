package esilatest

/*
200 ok object */
type GetCharactersCharacterIdKillmailsRecent200Ok struct {
	/*
	 ID of this killmail */
	KillmailId int32 `json:"killmail_id,omitempty"`
	/*
	 A hash of this killmail */
	KillmailHash string `json:"killmail_hash,omitempty"`
}