/*
 * Unit Servise API
 *
 * This is TEST API for my service
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EditUnit struct {
	// FULL NAME unit
	Name string `json:"name"`
	// force name unit
	ForceName string `json:"forceName"`
	// count heals point unit
	Hp int32 `json:"hp"`
	// initiative unit
	Initiative int32 `json:"initiative"`
	// ability to shoot unit
	Bs int32 `json:"bs"`
	// ability to fight unit
	Fs int32 `json:"fs"`
	// Additionat ability soldes
	AdditionalRule string `json:"additionalRule"`
}
