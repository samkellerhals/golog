package utils

var ActivityFields = map[string][]string{
	"hike":  {"Mountain", "Location", "Description"},
	"climb": {"Route", "Length", "Description"},
}

/*
E.g. ActivityTypes:
		Hiking:
			MountainName: string
			Latitude: float64
			Longitude: float64
			Location: string
			Route: string
			Elevation: string
			DurationHours: int8
			Description: string
		Climbing:
			RouteName: string
			Lat: string
			Lon: string
			Location: string
			Elevation: string
			ClimbingType: string
			Description: string
*/
