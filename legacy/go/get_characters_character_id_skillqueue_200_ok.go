package esilegacy

import "time"

/*
200 ok object */
type GetCharactersCharacterIdSkillqueue200Ok struct {
	/*
	 skill_id integer */
	SkillId int32 `json:"skill_id,omitempty"`
	/*
	 finish_date string */
	FinishDate time.Time `json:"finish_date,omitempty"`
	/*
	 start_date string */
	StartDate time.Time `json:"start_date,omitempty"`
	/*
	 finished_level integer */
	FinishedLevel int32 `json:"finished_level,omitempty"`
	/*
	 queue_position integer */
	QueuePosition int32 `json:"queue_position,omitempty"`
	/*
	 training_start_sp integer */
	TrainingStartSp int32 `json:"training_start_sp,omitempty"`
	/*
	 level_end_sp integer */
	LevelEndSp int32 `json:"level_end_sp,omitempty"`
	/*
	 Amount of SP that was in the skill when it started training it's current level. Used to calculate % of current level complete. */
	LevelStartSp int32 `json:"level_start_sp,omitempty"`
}
