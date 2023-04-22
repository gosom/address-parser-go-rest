package addressparser

import "errors"

var ErrAddressUnparsable = errors.New("address is unparsable")

// Address is a struct for an address
type Address struct {
	// venue name e.g. "Brooklyn Academy of Music", and building names e.g. "Empire State Building"
	House string `json:"house,omitempty"`
	// for category queries like "restaurants", etc.
	Category string `json:"category,omitempty"`
	// phrases like "in", "near", etc. used after a category phrase to help with parsing queries like "restaurants in Brooklyn"
	Near string `json:"near,omitempty"`
	// usually refers to the external (street-facing) building number. In some countries this may be a compount, hyphenated number which also includes an apartment number, or a block number (a la Japan), but libpostal will just call it the house_number for simplicity.
	HouseNumber string `json:"house_number,omitempty"`
	// street name(s)
	Road string `json:"road,omitempty"`
	// an apartment, unit, office, lot, or other secondary unit designator
	Unit string `json:"unit,omitempty"`
	// expressions indicating a floor number e.g. "3rd Floor", "Ground Floor", etc.
	Level string `json:"level,omitempty"`
	// numbered/lettered staircase
	Staircase string `json:"staircase,omitempty"`
	// numbered/lettered entrance
	Entrance string `json:"entrance,omitempty"`
	// post office box: typically found in non-physical (mail-only) addresses
	PoBox string `json:"po_box,omitempty"`
	// postal codes used for mail sorting
	Postcode string `json:"postcode,omitempty"`
	// usually an unofficial neighborhood name like "Harlem", "South Bronx", or "Crown Heights"
	Suburb string `json:"suburb,omitempty"`
	//  these are usually boroughs or districts within a city that serve some official purpose e.g. "Brooklyn" or "Hackney" or "Bratislava IV"
	CityDistrict string `json:"city_district,omitempty"`
	// any human settlement including cities, towns, villages, hamlets, localities, etc.
	City string `json:"city,omitempty"`
	// named islands e.g. "Maui"
	Island string `json:"island,omitempty"`
	// usually a second-level administrative division or county.
	StateDistrict string `json:"state_district,omitempty"`
	// a first-level administrative division. Scotland, Northern Ireland, Wales, and England in the UK are mapped to "state" as well (convention used in OSM, GeoPlanet, etc.)
	State string `json:"state,omitempty"`
	// informal subdivision of a country without any political status
	CountryRegion string `json:"country_region,omitempty"`
	// sovereign nations and their dependent territories, anything with an ISO-3166 code.
	Country string `json:"country,omitempty"`
	// currently only used for appending “West Indies” after the country name, a pattern frequently used in the English-speaking Caribbean e.g. “Jamaica, West Indies”
	WorldRegion string `json:"world_region,omitempty"`
	// Components is the raw response from libpostal
	Components []AddressComponent `json:"components"`
}

// AddressComponent is a struct for an address component
type AddressComponent struct {
	// Label is the label of the component as defined by libpostal
	Label string `json:"label"`
	// Value is the value of the component as defined by libpostal
	Value string `json:"value"`
}

// AddressParserInput is a struct for the input to the address parser
type AddressParserInput struct {
	// the address to parse
	Address string `json:"address" validate:"required"`
	// the language of the address. Leave empty if you don't know
	Language string `json:"language,omitempty"`
	// the country of the address. Leave empty if you don't know
	Country string `json:"country,omitempty"`
	// if true then the responses will be title Cased. Default behavior of libpostal is not to do that.
	TitleCase bool `json:"title_case,omitempty"`
}

// AddressParser is an interface for the address parser
type AddressParser interface {
	Parse(input AddressParserInput) (Address, error)
}
