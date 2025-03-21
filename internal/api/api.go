package api

import "errors"

type Trips [][2]string

func TripsToItinerary(trips Trips) ([]string, error) {
	var recursion func(Trips, []string) ([]string, error)
	maxRecursionDepth := len(trips)
	count := 0
	recursion = func(trips Trips, sortedTrips []string) ([]string, error) {
		if count >= maxRecursionDepth {
			return nil, errors.New("impossible itinerary")
		}
		count += 1
		unsortedTrips := make(Trips, 0)
		if len(trips) == 0 {
			return sortedTrips, nil
		}
		for _, trip := range trips {
			if len(trip) != 2 {
				return nil, errors.New("every trip must contain a pair of strings")
			}
			if trip[0] == sortedTrips[len(sortedTrips)-1] {
				sortedTrips = append(sortedTrips, trip[1])
			} else if trip[1] == sortedTrips[0] {
				sortedTrips = append([]string{trip[0]}, sortedTrips...)
			} else {
				unsortedTrips = append(unsortedTrips, trip)
			}
		}
		return recursion(unsortedTrips, sortedTrips)
	}
	if len(trips) < 1 {
		return nil, errors.New("the array containing the trips can't be empty")
	}
	return recursion(trips[1:], []string{trips[0][0], trips[0][1]})
}
