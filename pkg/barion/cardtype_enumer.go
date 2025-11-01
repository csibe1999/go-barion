package barion

import (
	"encoding/json"
	"fmt"
)

func (i CardType) String() string {
	switch i {
	case CardTypeUnknown:
		return "Unknown"
	case CardTypeMasterCard:
		return "MasterCard"
	case CardTypeVisa:
		return "Visa"
	case CardTypeAmericanExpress:
		return "AmericanExpress"
	case CardTypeElectron:
		return "Electron"
	case CardTypeMaestro:
		return "Maestro"
	case CardTypeDinersClub:
		return "DinersClub"
	case CardTypeDiscover:
		return "Discover"
	default:
		return fmt.Sprintf("CardType(%d)", i)
	}
}

var _CardTypeValues = []CardType{
	CardTypeUnknown,
	CardTypeMasterCard,
	CardTypeVisa,
	CardTypeAmericanExpress,
	CardTypeElectron,
	CardTypeMaestro,
	CardTypeDinersClub,
	CardTypeDiscover,
}

var _CardTypeNameToValueMap = map[string]CardType{
	"Unknown":         CardTypeUnknown,
	"MasterCard":      CardTypeMasterCard,
	"Visa":            CardTypeVisa,
	"AmericanExpress": CardTypeAmericanExpress,
	"Electron":        CardTypeElectron,
	"Maestro":         CardTypeMaestro,
	"DinersClub":      CardTypeDinersClub,
	"Discover":        CardTypeDiscover,
}

// CardTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CardTypeString(s string) (CardType, error) {
	if val, ok := _CardTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CardType values", s)
}

// CardTypeValues returns all values of the enum
func CardTypeValues() []CardType {
	return _CardTypeValues
}

// IsACardType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CardType) IsACardType() bool {
	for _, v := range _CardTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CardType
func (i CardType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CardType
func (i *CardType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CardType should be a string, got %s", data)
	}

	var err error
	*i, err = CardTypeString(s)
	return err
}
