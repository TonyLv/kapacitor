package tick

import (
	"github.com/influxdata/kapacitor/pipeline"
	"github.com/influxdata/kapacitor/tick/ast"
)

// InfluxDBOut converts the InfluxDBOut pipeline node into the TICKScript AST
type InfluxDBOut struct {
	Function
}

// NewInfluxDBOut creates a InfluxQL function builder
func NewInfluxDBOut(parents []ast.Node) *InfluxDBOut {
	return &InfluxDBOut{
		Function{
			Parents: parents,
		},
	}
}

// Build creates a InfluxDBOut ast.Node
func (n *InfluxDBOut) Build(db *pipeline.InfluxDBOutNode) (ast.Node, error) {
	n.Pipe("influxDBOut").
		Dot("cluster", db.Cluster).
		Dot("database", db.Database).
		Dot("retentionPolicy", db.RetentionPolicy).
		Dot("measurement", db.Measurement).
		Dot("writeConsistency", db.WriteConsistency).
		Dot("precision", db.Precision).
		Dot("buffer", db.Buffer).
		Dot("flushInterval", db.FlushInterval).
		DotIf("create", db.CreateFlag)

	for k, v := range db.Tags {
		n.Dot("tag", k, v)
	}

	return n.prev, n.err
}