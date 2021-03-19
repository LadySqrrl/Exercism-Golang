package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	var planetAge float64

	switch planet {
	case "Mercury":
		planetAge = (seconds / 31557600) / 0.2408467

	case "Venus":
		planetAge = (seconds / 31557600) / 0.61519726

	case "Earth":
		planetAge = seconds / 31557600

	case "Mars":
		planetAge = (seconds / 31557600) / 1.8808158

	case "Jupiter":
		planetAge = (seconds / 31557600) / 11.862615

	case "Saturn":
		planetAge = (seconds / 31557600) / 29.447498

	case "Uranus":
		planetAge = (seconds / 31557600) / 84.016846

	case "Neptune":
		planetAge = (seconds / 31557600) / 164.79132

	}

	return planetAge
}
