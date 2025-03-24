package api

import "errors"

type Trips [][2]string

// Takes a slice with Trips and return slice with strings representing the itinerary if the
// itinerary is possible and the input is not empty, else error
func TripsToItinerary(trips Trips) ([]string, error) {
	var recursion func(Trips, []string) ([]string, error)
	maxRecursionDepth := len(trips)
	count := 0
	recursion = func(trips Trips, sortedTrips []string) ([]string, error) {
		// all trips are sorted and appended to sortedTrips, return the result
		if len(trips) == 0 {
			return sortedTrips, nil
		}

		// if recursiondepth is larger then the length of trips, it's an impossible itinerary
		if count >= maxRecursionDepth {
			return nil, errors.New("impossible itinerary")
		}

		count += 1
		var unsortedTrips Trips

		for _, trip := range trips {
			origin, destination := trip[0], trip[1]
			// origin of the trip is equal to last element of slice - append destination
			if origin == sortedTrips[len(sortedTrips)-1] {
				sortedTrips = append(sortedTrips, destination)
				// destination of the trip is equal to first element of slice - prepend origin
			} else if destination == sortedTrips[0] {
				sortedTrips = append([]string{origin}, sortedTrips...)
			} else {
				unsortedTrips = append(unsortedTrips, trip)
			}
		}
		return recursion(unsortedTrips, sortedTrips)
	}

	if len(trips) < 1 {
		return nil, errors.New("the array containing the trips can't be empty")
	}

	// uses the first trip as starting point for constructing the itinerary
	return recursion(trips[1:], []string{trips[0][0], trips[0][1]})
}
