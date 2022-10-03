package fixtures

import (
	m "github.com/RedHatInsights/sources-api-go/model"
)

var (
	uid1 = "5eebe172-7baa-4280-823f-19e597d091e9"
	uid2 = "31b5338b-685d-4056-ba39-d00b4d7f19cc"
	uid3 = "36be1c27-ef96-42b0-9a13-72240b18cf83"
	uid4 = "1c8b6c9a-af40-11ec-b909-0242ac120002"
	uid5 = "5cbb40a8-f66a-4efb-8ed2-5f18c59ff7ca"
	uid6 = "f6d1e4ae-781c-4be0-a4ed-8935af5d9f47"
)

var TestSourceData = []m.Source{
	{
		ID:                 1,
		Name:               "Source1",
		SourceTypeID:       1,
		TenantID:           1,
		AvailabilityStatus: "available",
		Uid:                &uid1,
	},
	{
		ID:                 2,
		Name:               "Source2",
		SourceTypeID:       1,
		TenantID:           1,
		AvailabilityStatus: "unavailable",
		Uid:                &uid2,
	},
	{
		ID:                 100,
		Name:               "Source3 for TestSourceDelete()",
		SourceTypeID:       1,
		TenantID:           1,
		AvailabilityStatus: "available",
		Uid:                &uid3,
	},
	{
		ID:                 4,
		Name:               "Source4",
		SourceTypeID:       2,
		TenantID:           1,
		AvailabilityStatus: "available",
		Uid:                &uid4,
	},
	{
		ID:                 101,
		Name:               "Source5 without applications",
		SourceTypeID:       2,
		TenantID:           1,
		AvailabilityStatus: "available",
		Uid:                &uid5,
	},
	{
		ID:                 6,
		Name:               "Source6 Satellite",
		SourceTypeID:       5,
		TenantID:           1,
		AvailabilityStatus: "available",
		Uid:                &uid6,
	},
}
