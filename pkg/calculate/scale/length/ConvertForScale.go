package length

import "math"

func ConvertForScale(length float64, sourceUnit string, scale float64, destinationUnit string) (float64, error) {
	sourceFactor, sourceUnitError := GetFactorForUnit(sourceUnit)
	if sourceUnitError != nil {
		return length, sourceUnitError
	}

	destinationFactor, destinationUnitError := GetFactorForUnit(destinationUnit)
	if destinationUnitError != nil {
		return length, destinationUnitError
	}

	normalizedLength := length
	if sourceFactor > 0 {
		normalizedLength = length * math.Abs(float64(sourceFactor))
	} else {
		normalizedLength = length / math.Abs(float64(sourceFactor))
	}


	scaledLength := normalizedLength * scale

	if destinationFactor < 0 {
		return scaledLength * math.Abs(float64(destinationFactor)), nil
	}

	return scaledLength / math.Abs(float64(destinationFactor)), nil
}
