package validation

import entity "kwick/model/entity"

func CompareAddresses(ra, da entity.Address) bool {
	if ra.BuildingName == da.BuildingName && ra.BuildingNumber == da.BuildingNumber && ra.Street == da.Street && ra.City == da.City && ra.Country == da.Country && ra.PostalCode == da.PostalCode && ra.UserId == da.UserId {
		return false
	}
	return true
}
