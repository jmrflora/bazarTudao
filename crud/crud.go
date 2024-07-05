package crud

import (
	"github.com/jmrflora/bazarTudao/ent"
)

type crud struct {
	c *ent.Client
}

func NewCrud(driverName string, dataSourceName string) (*crud, error) {
	c, err := ent.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &crud{c: c}, nil
}
