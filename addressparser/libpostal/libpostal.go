package libpostal

import (
	"strings"

	"github.com/gosom/kit/logging"
	postal "github.com/openvenues/gopostal/parser"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/gosom/address-parser-go-rest/addressparser"
)

var _ addressparser.AddressParser = (*libPostalParser)(nil)

type libPostalParser struct {
	log logging.Logger
}

func (o *libPostalParser) Parse(input addressparser.AddressParserInput) (addressparser.Address, error) {
	components := postal.ParseAddressOptions(input.Address, postal.ParserOptions{
		Language: input.Language,
		Country:  input.Country,
	})
	if len(components) == 0 {
		return addressparser.Address{}, addressparser.ErrAddressUnparsable
	}
	address := addressparser.Address{}
	tag := language.Und
	if input.Language != "" {
		if r, err := language.Parse("de"); err == nil {
			tag = r
		}
	}
	houseNumberFound := false
	for i := range components {
		if input.TitleCase {
			components[i].Value = cases.Title(tag, cases.NoLower).String(components[i].Value)
		}
		switch components[i].Label {
		case "house":
			address.House = components[i].Value
		case "category":
			address.Category = components[i].Value
		case "near":
			address.Near = components[i].Value
		case "house_number":
			if !houseNumberFound {
				address.HouseNumber = components[i].Value
				houseNumberFound = true
			}
		case "road":
			address.Road = components[i].Value
		case "unit":
			address.Unit = components[i].Value
		case "level":
			address.Level = components[i].Value
		case "staircase":
			address.Staircase = components[i].Value
		case "entrance":
			address.Entrance = components[i].Value
		case "po_box":
			address.PoBox = components[i].Value
		case "postcode":
			address.Postcode = components[i].Value
		case "suburb":
			address.Suburb = components[i].Value
		case "city_district":
			address.CityDistrict = components[i].Value
		case "city":
			address.City = components[i].Value
		case "island":
			address.Island = components[i].Value
		case "state_district":
			address.StateDistrict = components[i].Value
		case "state":
			address.State = components[i].Value
		case "country_region":
			address.CountryRegion = components[i].Value
		case "country":
			address.Country = components[i].Value
		case "world_region":
			address.WorldRegion = components[i].Value
		default:
			o.log.Warn("Unknown component", "component", components[i].Label)
		}

		component := addressparser.AddressComponent{
			Label: components[i].Label,
			Value: components[i].Value,
		}
		address.Components = append(address.Components, component)
	}
	address.HouseNumber = strings.TrimSpace(address.HouseNumber)
	return address, nil
}

func NewLibPostalParser(log logging.Logger) *libPostalParser {
	return &libPostalParser{log: log}
}
