package esilegacy

/*
200 ok object */
type GetUniverseGraphicsGraphicIdOk struct {
	/*
	 graphic_id integer */
	GraphicId int32 `json:"graphic_id,omitempty"`
	/*
	 graphic_file string */
	GraphicFile string `json:"graphic_file,omitempty"`
	/*
	 sof_race_name string */
	SofRaceName string `json:"sof_race_name,omitempty"`
	/*
	 sof_fation_name string */
	SofFationName string `json:"sof_fation_name,omitempty"`
	/*
	 sof_dna string */
	SofDna string `json:"sof_dna,omitempty"`
	/*
	 sof_hull_name string */
	SofHullName string `json:"sof_hull_name,omitempty"`
	/*
	 collision_file string */
	CollisionFile string `json:"collision_file,omitempty"`
	/*
	 icon_folder string */
	IconFolder string `json:"icon_folder,omitempty"`
}
