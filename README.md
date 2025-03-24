# iternaryparser

A minimal Go web server that accepts a JSON array of string pairs representing trips and returns the iternary as a JSON array of strings


### Run the server

```bash
go run cmd/server/main.go
```

The server starts on port `8080` by default.

### `POST /trips`

**Request Body** (`Content-Type: application/json`):

```json
[
    ["LAX", "DXB"], 
    ["JFK", "LAX"], 
    ["SFO", "SJC"], 
    ["DXB", "SFO"]
]
```

**Response:**

```json
[
    "JFK",
    "LAX",
    "DXB",
    "SFO",
    "SJC"
]
```

Returns a reconstructed itinerary from the unordered list of trips.

## Project Structure

```
cmd/server         # Main entry point
internal/server    # Server setup and routing
internal/api       # Handler logic and transformation
```

## Improvements

The current algorithm is not efficent and has polynomal time complexity, approximately n^2. I haven't spend time trying to optimise it, since I got the impression that I was to implement a quick solution rather than an optimal one. A linear-time algorithm is very easy to implement using maps. ChatGPT suggest the following alternative:

```go
func TripsToItinerary(trips [][]string) ([]string, error) {
    if len(trips) == 0 {
        return nil, errors.New("no trips provided")
    }
    
    // Build the map and a set of destination airports.
    fromTo := make(map[string]string)
    dests := make(map[string]bool)
    for _, trip := range trips {
        if len(trip) != 2 {
            return nil, errors.New("each trip must have exactly 2 elements")
        }
        fromTo[trip[0]] = trip[1]
        dests[trip[1]] = true
    }
    
    // Find the starting airport: one that is not a destination.
    start := ""
    for _, trip := range trips {
        if !dests[trip[0]] {
            start = trip[0]
            break
        }
    }
    if start == "" {
        return nil, errors.New("unable to find a valid starting airport")
    }
    
    // Build the itinerary by following the chain.
    itinerary := []string{start}
    for {
        next, ok := fromTo[start]
        if !ok {
            break
        }
        itinerary = append(itinerary, next)
        start = next
    }
    
    return itinerary, nil
}
```