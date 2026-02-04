package simbad

type ObjectInfo struct {
	Identifier   string
	ObjectType   string
	SpectralType string
	VMagnitude   *float64
	Parallax     *float64
	RA           *float64
	Dec          *float64
}

func (o *ObjectInfo) DistanceParsecs() *float64 {
	if o.Parallax == nil || *o.Parallax <= 0 {
		return nil
	}
	d := 1000.0 / *o.Parallax
	return &d
}
