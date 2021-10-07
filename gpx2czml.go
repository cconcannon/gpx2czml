package gpx2czml

import (
	"strings"

	"github.com/cconcannon/czml"
	"github.com/cconcannon/gpx"
)

// CreatePolyline creates a Polyline data structure and returns a new CZML data structure with the
// Polyline added in a Packet. An id and a name for the packet must be provided.
func CreatePolyline(id, name string, g gpx.Gpx) czml.Czml {
	var c czml.Czml

	initializeCzml(&g, &c)

	p := czml.CreateEmptyPacket(id, name)

	p.AddEmptyPolyline("white")

	tracks := g.GetTracks()
	segments := tracks[0].GetSegments()
	points := segments[0].GetTrackPoints()

	for _, point := range points {
		p.Polyline.AddPoint(point.Lat, point.Lon, point.Elevation)
	}

	c.AddPacket(p)

	return c
}

func CreatePath(id, name string, g gpx.Gpx) czml.Czml {
	var c czml.Czml

	initializeCzml(&g, &c)

	p := czml.CreateEmptyPacket(id, name)

	p.AddPath("white")
	p.AddBillboard()

	tracks := g.GetTracks()
	segments := tracks[0].GetSegments()
	points := segments[0].GetTrackPoints()

	for _, point := range points {
		p.AddPosition(point.Time, point.Lat, point.Lon, point.Elevation)
	}

	startTime := points[0].Time
	stopTime := points[len(points)-1].Time
	interval := strings.Join([]string{startTime, stopTime}, "/")

	c.AddPacket(p)
	c.AddClock(interval, startTime, 50)

	return c
}

// transfer metadata from gpx, if it exists
func initializeCzml(g *gpx.Gpx, c *czml.Czml) {
	var docName string
	if g.Metadata != nil {
		docName = g.Metadata.Name
	} else {
		docName = "GPX2CZML Document"
	}

	c.InitializeDocument(docName)
}
