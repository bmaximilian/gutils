package length

import "errors"

func GetFactorForUnit(unit string) (int, error)  {
	if len(unit) == 1 {
		return 1, nil
	}

	unitFactorString := unit[:1]

	unitFactor, err := factorIntForString(unitFactorString)
	if err != nil {
		return 1, err
	}

	return unitFactor, nil
}

func factorIntForString(factorUnit string) (int, error) {
	switch factorUnit {
	case "m":
		return -1000, nil
	case "c":
		return -100, nil
	case "":
		return 1, nil
	case "d":
		return 10, nil
	case "k":
		return 1000, nil
	default:
		return 1, errors.New("the factor is not valid")
	}
}
