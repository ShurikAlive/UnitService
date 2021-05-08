/*
 * Unit Servise API
 *
 * This is TEST API for my service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package Swagger

type Equipment struct {
	// ID equipment
	Id string `json:"id"`
	// FULL NAME equipment
	Name string `json:"name"`
	// limit equipment on one unit. -1 - unlimit
	LimitOnUnit int32 `json:"limitOnUnit"`
	// limit equipment on one team. -1 - unlimit
	LimitOnTeam int32 `json:"limitOnTeam"`
	// The role of a soldier available when selecting ammunition.
	SoldarRole string `json:"soldarRole"`
	// game rule equipment
	Rule string `json:"rule"`
	// limit equipment on game. -1 - unlimit
	Ammo int32 `json:"ammo"`
	// cost equipment in game points
	Cost int32 `json:"cost"`
}