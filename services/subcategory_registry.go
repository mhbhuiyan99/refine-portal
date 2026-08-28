package services

import "net/url"

// SubCategory describes a supported sub-category and its API parameters.
type SubCategory struct {
	CanonicalKey string
	Params       url.Values
}

// subCategoryDefinitions contains the API parameters for each
// canonical sub-category. Parameters are defined only once.
var subCategoryDefinitions = map[string]SubCategory{
	"hotel": {
		CanonicalKey: "hotel",
		Params:       url.Values{"pt": {"5"}},
	},
	"cottage": {
		CanonicalKey: "cottage",
		Params:       url.Values{"pt": {"12"}},
	},
	"cabin": {
		CanonicalKey: "cabin",
		Params:       url.Values{"pt": {"3"}},
	},
	"villas": {
		CanonicalKey: "villas",
		Params:       url.Values{"pt": {"9"}},
	},
	"resort": {
		CanonicalKey: "resort",
		Params:       url.Values{"pt": {"7"}},
	},
	"skiChalets": {
		CanonicalKey: "skiChalets",
		Params:       url.Values{"pt": {"13"}},
	},
	"timeShares": {
		CanonicalKey: "timeShares",
		Params:       url.Values{"pt": {"16"}},
	},
	"uniqueVacationRental": {
		CanonicalKey: "uniqueVacationRental",
		Params:       url.Values{"pt": {"14-15"}},
	},
	"vacationRental": {
		CanonicalKey: "vacationRental",
		Params:       url.Values{"order": {"1"}},
	},
	"familyRental": {
		CanonicalKey: "familyRental",
		Params:       url.Values{"amenities": {"5"}},
	},
	"petFriendly": {
		CanonicalKey: "petFriendly",
		Params:       url.Values{"amenities": {"11"}},
	},
	"pools": {
		CanonicalKey: "pools",
		Params:       url.Values{"amenities": {"12"}},
	},
	"ocean": {
		CanonicalKey: "ocean",
		Params:       url.Values{"amenities": {"19-20"}},
	},
	"beachRental": {
		CanonicalKey: "beachRental",
		Params:       url.Values{"amenities": {"18-19"}},
	},
	"luxuryRental": {
		CanonicalKey: "luxuryRental",
		Params:       url.Values{"order": {"3"}},
	},
	"discount": {
		CanonicalKey: "discount",
		Params:       url.Values{"order": {"2"}},
	},
	"groupTravel": {
		CanonicalKey: "groupTravel",
		Params:       url.Values{"pax": {"5"}},
	},
	"winter": {
		CanonicalKey: "winter",
		Params:       url.Values{"isWinter": {"1"}},
	},
	"summerRentals": {
		CanonicalKey: "summerRentals",
		Params:       url.Values{"isSummer": {"1"}},
	},
	"shortTerm": {
		CanonicalKey: "shortTerm",
		Params:       url.Values{"isShortTermStays": {"1"}},
	},
	"businessTravel": {
		CanonicalKey: "businessTravel",
		Params:       url.Values{"isBusinessTravel": {"1"}},
	},
	"sustainableTravel": {
		CanonicalKey: "sustainableTravel",
		Params:       url.Values{"ecoFriendly": {"1"}},
	},
	"holidayHome": {
		CanonicalKey: "holidayHome",
		Params: url.Values{
			"sqs":   {"rentals"},
			"order": {"1"},
		},
	},
	"monthlyStays": {
		CanonicalKey: "monthlyStays",
		Params:       url.Values{"isMonthStays": {"1"}},
	},
}

// subCategorySlugAliases maps URL slugs to canonical sub-category keys.
var subCategorySlugAliases = map[string]string{
	"hotels":                  "hotel",
	"cottages":                "cottage",
	"cabins":                  "cabin",
	"villas":                  "villas",
	"resorts":                 "resort",
	"ski-chalets":             "skiChalets",
	"timeshares":              "timeShares",
	"unique-vacation-rentals": "uniqueVacationRental",
	"vacation-rentals":        "vacationRental",
	"family-rentals":          "familyRental",
	"family":                  "familyRental",
	"pet-friendly-rentals":    "petFriendly",
	"pet-friendly":            "petFriendly",
	"rentals-with-pools":      "pools",
	"oceanfront-rentals":      "ocean",
	"beach-rentals":            "beachRental",
	"beach":                    "beachRental",
	"luxury-rentals":           "luxuryRental",
	"luxury":                   "luxuryRental",
	"discount-rentals":         "discount",
	"discount":                "discount",
	"group-travel":             "groupTravel",
	"winter-rentals":           "winter",
	"summer-rentals":           "summerRentals",
	"short-term-stays":         "shortTerm",
	"business-travel":          "businessTravel",
	"sustainable-travel":       "sustainableTravel",
	"holiday-homes":            "holidayHome",
	"monthly-stays":            "monthlyStays",
}

// knownUnmappedSlugs contains sub-category routes listed by the
// specification for which no API mapping was provided.
var knownUnmappedSlugs = map[string]struct{}{
	"condos":       {},
	"boat-rentals": {},
	"rv-rentals":   {},
}

// LookupSubCategory resolves a URL slug.
//
// mapped:
//   The slug has a known API mapping.
//
// known:
//   The slug is explicitly listed as a sub-category route, even if
//   its API mapping has not been provided.
//
// If both are false, the slug is treated as a normal location segment.
func LookupSubCategory(slug string) (
	subCategory SubCategory,
	mapped bool,
	known bool,
) {
	key, ok := subCategorySlugAliases[slug]
	if ok {
		return subCategoryDefinitions[key], true, true
	}

	if _, ok := knownUnmappedSlugs[slug]; ok {
		return SubCategory{}, false, true
	}

	return SubCategory{}, false, false
}

// SubCategorySlugs returns all URL slugs registered
// as supported sub-categories.
func SubCategorySlugs() []string {
	slugs := make([]string, 0, len(subCategorySlugAliases))

	for slug := range subCategorySlugAliases {
		slugs = append(slugs, slug)
	}

	return slugs
}