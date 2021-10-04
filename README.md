# gpx2czml

A package to convert `.gpx` file data to `.czml` file data for use with the Cesium ecosystem

## Usage

```go
import (
	"io/ioutil"
	"os"

	"github.com/cocncannon/czml"
	"github.com/cconcannon/gpx"
	"github.com/cconcannon/gpx2czml"
)
```

```go
// use a package to unmarshal GPX data from the XML file
// i.e. [github.com/cconcannon/gpx](https://github.com/cconcannon/gpx)

file, err := os.Open("input.xml")
data, err := ioutil.ReadAll(file)

if err != nil {
	fmt.Println(err)
	return
}

var g gpx.Gpx
err = gpx.Unmarshal(data, &g)

// transform the data from GPX to CZML properties
// Polyline is just one type of valid material to use

c := gpx2czml.CreatePolyline("abc123-packet-id", "Packet Name", g)

// write the data to file using a package to marshal CZML data to JSON
// i.e [github.com/cconcannon/czml](https://github.com/cconcannon/czml)

transform, err := czml.Marshal(c)

if err != nil {
	fmt.Println(err)
}

ioutil.WriteFile("output.czml", transform, 0644)
```

## Dependencies

I created these two Go modules to handle the `.gpx` and `.czml` data interfaces that are required for the conversion:

- [github.com/cconcannon/gpx](https://github.com/cconcannon/gpx)
- [github.com/cconcannon/czml](https://github.com/cconcannon/czml)