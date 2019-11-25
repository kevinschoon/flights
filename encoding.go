package main

import (
	"bytes"
	"io"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding/dot"
)

type Encoding string

func (e Encoding) String() string {
	return string(e)
}

const (
	Json = Encoding("Json")
	Dot  = Encoding("Dot")
)

type Encoder interface {
	Encode(io.Writer) error
}

type dotEncoder struct {
	graph graph.Graph
}

func (d dotEncoder) Encode(w io.Writer) error {
	raw, err := dot.Marshal(d.graph, "", "", "")
	if err != nil {
		return err
	}
	_, err = bytes.NewBuffer(raw).WriteTo(w)
	return err
}

func NewEncoder(enc Encoding, g graph.Graph) Encoder {
	return dotEncoder{graph: g}
}
