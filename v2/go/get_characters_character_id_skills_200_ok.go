package esiv2

/*
200 ok object */
type GetCharactersCharacterIdSkills200Ok struct {
	/*
	 skill_id integer */
	SkillId int32 `json:"skill_id,omitempty"`
	/*
	 skillpoints_in_skill integer */
	SkillpointsInSkill int64 `json:"skillpoints_in_skill,omitempty"`
	/*
	 current_skill_level integer */
	CurrentSkillLevel int32 `json:"current_skill_level,omitempty"`
}
