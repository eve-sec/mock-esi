package esiv2

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/antihax/mock-esi/mockesi"
)

func init() {

	mockesi.NewRoute(
		"GetAlliances",
		"Get",
		"/v2/alliances/",
		GetAlliances,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdCorporations",
		"Get",
		"/v2/alliances/{alliance_id}/corporations/",
		GetAlliancesAllianceIdCorporations,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdIcons",
		"Get",
		"/v2/alliances/{alliance_id}/icons/",
		GetAlliancesAllianceIdIcons,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdAssetsLocations",
		"Post",
		"/v2/characters/{character_id}/assets/locations/",
		PostCharactersCharacterIdAssetsLocations,
	)

	mockesi.NewRoute(
		"PostCorporationsCorporationIdAssetsLocations",
		"Post",
		"/v2/corporations/{corporation_id}/assets/locations/",
		PostCorporationsCorporationIdAssetsLocations,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarks",
		"Get",
		"/v2/characters/{character_id}/bookmarks/",
		GetCharactersCharacterIdBookmarks,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarksFolders",
		"Get",
		"/v2/characters/{character_id}/bookmarks/folders/",
		GetCharactersCharacterIdBookmarksFolders,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendar",
		"Get",
		"/v2/characters/{character_id}/calendar/",
		GetCharactersCharacterIdCalendar,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendarEventIdAttendees",
		"Get",
		"/v2/characters/{character_id}/calendar/{event_id}/attendees/",
		GetCharactersCharacterIdCalendarEventIdAttendees,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAgentsResearch",
		"Get",
		"/v2/characters/{character_id}/agents_research/",
		GetCharactersCharacterIdAgentsResearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBlueprints",
		"Get",
		"/v2/characters/{character_id}/blueprints/",
		GetCharactersCharacterIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCorporationhistory",
		"Get",
		"/v2/characters/{character_id}/corporationhistory/",
		GetCharactersCharacterIdCorporationhistory,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFatigue",
		"Get",
		"/v2/characters/{character_id}/fatigue/",
		GetCharactersCharacterIdFatigue,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMedals",
		"Get",
		"/v2/characters/{character_id}/medals/",
		GetCharactersCharacterIdMedals,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdNotificationsContacts",
		"Get",
		"/v2/characters/{character_id}/notifications/contacts/",
		GetCharactersCharacterIdNotificationsContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPortrait",
		"Get",
		"/v2/characters/{character_id}/portrait/",
		GetCharactersCharacterIdPortrait,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdRoles",
		"Get",
		"/v2/characters/{character_id}/roles/",
		GetCharactersCharacterIdRoles,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdStandings",
		"Get",
		"/v2/characters/{character_id}/standings/",
		GetCharactersCharacterIdStandings,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdStats",
		"Get",
		"/v2/characters/{character_id}/stats/",
		GetCharactersCharacterIdStats,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdTitles",
		"Get",
		"/v2/characters/{character_id}/titles/",
		GetCharactersCharacterIdTitles,
	)

	mockesi.NewRoute(
		"PostCharactersAffiliation",
		"Post",
		"/v2/characters/affiliation/",
		PostCharactersAffiliation,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdClones",
		"Get",
		"/v2/characters/{character_id}/clones/",
		GetCharactersCharacterIdClones,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdImplants",
		"Get",
		"/v2/characters/{character_id}/implants/",
		GetCharactersCharacterIdImplants,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdContacts",
		"Delete",
		"/v2/characters/{character_id}/contacts/",
		DeleteCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdContacts",
		"Get",
		"/v2/alliances/{alliance_id}/contacts/",
		GetAlliancesAllianceIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContacts",
		"Get",
		"/v2/characters/{character_id}/contacts/",
		GetCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContacts",
		"Get",
		"/v2/corporations/{corporation_id}/contacts/",
		GetCorporationsCorporationIdContacts,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdContacts",
		"Post",
		"/v2/characters/{character_id}/contacts/",
		PostCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdContacts",
		"Put",
		"/v2/characters/{character_id}/contacts/",
		PutCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdAlliancehistory",
		"Get",
		"/v2/corporations/{corporation_id}/alliancehistory/",
		GetCorporationsCorporationIdAlliancehistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdBlueprints",
		"Get",
		"/v2/corporations/{corporation_id}/blueprints/",
		GetCorporationsCorporationIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContainersLogs",
		"Get",
		"/v2/corporations/{corporation_id}/containers/logs/",
		GetCorporationsCorporationIdContainersLogs,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdIcons",
		"Get",
		"/v2/corporations/{corporation_id}/icons/",
		GetCorporationsCorporationIdIcons,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembers",
		"Get",
		"/v2/corporations/{corporation_id}/members/",
		GetCorporationsCorporationIdMembers,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStructures",
		"Get",
		"/v2/corporations/{corporation_id}/structures/",
		GetCorporationsCorporationIdStructures,
	)

	mockesi.NewRoute(
		"GetDogmaEffectsEffectId",
		"Get",
		"/v2/dogma/effects/{effect_id}/",
		GetDogmaEffectsEffectId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFwStats",
		"Get",
		"/v2/characters/{character_id}/fw/stats/",
		GetCharactersCharacterIdFwStats,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdFwStats",
		"Get",
		"/v2/corporations/{corporation_id}/fw/stats/",
		GetCorporationsCorporationIdFwStats,
	)

	mockesi.NewRoute(
		"GetFwLeaderboards",
		"Get",
		"/v2/fw/leaderboards/",
		GetFwLeaderboards,
	)

	mockesi.NewRoute(
		"GetFwLeaderboardsCharacters",
		"Get",
		"/v2/fw/leaderboards/characters/",
		GetFwLeaderboardsCharacters,
	)

	mockesi.NewRoute(
		"GetFwLeaderboardsCorporations",
		"Get",
		"/v2/fw/leaderboards/corporations/",
		GetFwLeaderboardsCorporations,
	)

	mockesi.NewRoute(
		"GetFwStats",
		"Get",
		"/v2/fw/stats/",
		GetFwStats,
	)

	mockesi.NewRoute(
		"GetFwSystems",
		"Get",
		"/v2/fw/systems/",
		GetFwSystems,
	)

	mockesi.NewRoute(
		"GetFwWars",
		"Get",
		"/v2/fw/wars/",
		GetFwWars,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFittings",
		"Get",
		"/v2/characters/{character_id}/fittings/",
		GetCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdFittings",
		"Post",
		"/v2/characters/{character_id}/fittings/",
		PostCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFleet",
		"Get",
		"/v2/characters/{character_id}/fleet/",
		GetCharactersCharacterIdFleet,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdLocation",
		"Get",
		"/v2/characters/{character_id}/location/",
		GetCharactersCharacterIdLocation,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOnline",
		"Get",
		"/v2/characters/{character_id}/online/",
		GetCharactersCharacterIdOnline,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdShip",
		"Get",
		"/v2/characters/{character_id}/ship/",
		GetCharactersCharacterIdShip,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailLabels",
		"Get",
		"/v2/characters/{character_id}/mail/labels/",
		GetCharactersCharacterIdMailLabels,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdMailLabels",
		"Post",
		"/v2/characters/{character_id}/mail/labels/",
		PostCharactersCharacterIdMailLabels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOrders",
		"Get",
		"/v2/characters/{character_id}/orders/",
		GetCharactersCharacterIdOrders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOrders",
		"Get",
		"/v2/corporations/{corporation_id}/orders/",
		GetCorporationsCorporationIdOrders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOrdersHistory",
		"Get",
		"/v2/corporations/{corporation_id}/orders/history/",
		GetCorporationsCorporationIdOrdersHistory,
	)

	mockesi.NewRoute(
		"GetSearch",
		"Get",
		"/v2/search/",
		GetSearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSkillqueue",
		"Get",
		"/v2/characters/{character_id}/skillqueue/",
		GetCharactersCharacterIdSkillqueue,
	)

	mockesi.NewRoute(
		"GetStatus",
		"Get",
		"/v2/status/",
		GetStatus,
	)

	mockesi.NewRoute(
		"GetUniverseFactions",
		"Get",
		"/v2/universe/factions/",
		GetUniverseFactions,
	)

	mockesi.NewRoute(
		"GetUniverseStationsStationId",
		"Get",
		"/v2/universe/stations/{station_id}/",
		GetUniverseStationsStationId,
	)

	mockesi.NewRoute(
		"GetUniverseStructuresStructureId",
		"Get",
		"/v2/universe/structures/{structure_id}/",
		GetUniverseStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetUniverseSystemKills",
		"Get",
		"/v2/universe/system_kills/",
		GetUniverseSystemKills,
	)

	mockesi.NewRoute(
		"GetUniverseTypesTypeId",
		"Get",
		"/v2/universe/types/{type_id}/",
		GetUniverseTypesTypeId,
	)

	mockesi.NewRoute(
		"PostUniverseNames",
		"Post",
		"/v2/universe/names/",
		PostUniverseNames,
	)

	mockesi.NewRoute(
		"PostUiAutopilotWaypoint",
		"Post",
		"/v2/ui/autopilot/waypoint/",
		PostUiAutopilotWaypoint,
	)

}

func errorOut(w http.ResponseWriter, r *http.Request, e error) {
	w.WriteHeader(http.StatusInternalServerError)

	w.Write([]byte(e.Error()))
}

func processParameters(data interface{}, input string) (v interface{}, err error) {
	switch reflect.TypeOf(data).String() {
	case "int":
		{
			v, err = strconv.Atoi(input)
			v = v.(int)
		}
	case "int8":
		{
			v, err = strconv.ParseInt(input, 10, 8)
			v = (int8)(v.(int64))
		}
	case "int16":
		{
			v, err = strconv.ParseInt(input, 10, 16)
			v = (int16)(v.(int64))
		}
	case "rune":
		{
			v, err = strconv.ParseInt(input, 10, 32)
			v = (rune)(v.(int64))
		}
	case "int32":
		{
			v, err = strconv.ParseInt(input, 10, 32)
			v = (int32)(v.(int64))
		}
	case "int64":
		{
			v, err = strconv.ParseInt(input, 10, 64)
			v = v.(int64)
		}
	case "uint8":
		{
			v, err = strconv.ParseUint(input, 10, 8)
			v = (uint8)(v.(uint64))
		}
	case "byte":
		{
			v, err = strconv.ParseUint(input, 10, 8)
			v = (byte)(v.(uint64))
		}
	case "uint16":
		{
			v, err = strconv.ParseUint(input, 10, 16)
			v = (uint16)(v.(uint64))
		}
	case "uint32":
		{
			v, err = strconv.ParseUint(input, 10, 32)
			v = (uint32)(v.(uint64))
		}
	case "uint64":
		{
			v, err = strconv.ParseUint(input, 10, 64)
			v = v.(uint64)
		}
	case "float32":
		{
			v, err = strconv.ParseFloat(input, 32)
			v = (float32)(v.(float64))
		}
	case "float64":
		{
			v, err = strconv.ParseFloat(input, 64)
			v = v.(float64)
		}
	case "bool":
		{
			v, err = strconv.ParseBool(input)
			v = v.(bool)
		}
	case "string":
		{
			v = input
		}
	}

	return
}
