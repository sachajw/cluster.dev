package tfmodule

import (
	"github.com/apex/log"
	"github.com/shalb/cluster.dev/pkg/modules/shell/terraform/base"
	"github.com/shalb/cluster.dev/pkg/project"
)

// Factory factory for s3 backends.
type Factory struct {
}

const unitKind string = "terraform"

func NewEmptyUnit() UnitTfModule {
	unit := UnitTfModule{
		Unit:     base.NewEmptyUnit(),
		StatePtr: &UnitTfModule{},
		UnitKind: unitKind,
	}
	return unit
}

func NewUnit(spec map[string]interface{}, stack *project.Stack) (*UnitTfModule, error) {
	unit := NewEmptyUnit()
	cUnit, err := base.NewUnit(spec, stack)
	if err != nil {
		log.Debug(err.Error())
		return nil, err
	}
	unit.Unit = *cUnit
	err = unit.ReadConfig(spec, stack)
	if err != nil {
		log.Debug(err.Error())
		return nil, err
	}
	return &unit, nil
}

// New creates new unit.
func (f *Factory) New(spec map[string]interface{}, stack *project.Stack) (project.Unit, error) {
	return NewUnit(spec, stack)
}

// NewFromState creates new unit from state data.
func (f *Factory) NewFromState(spec map[string]interface{}, modKey string, p *project.StateProject) (project.Unit, error) {
	mod := NewEmptyUnit()
	err := mod.LoadState(spec, modKey, p)
	if err != nil {
		log.Debug(err.Error())
		return nil, err
	}
	return &mod, nil
}

func init() {
	modDrv := Factory{}
	log.Debugf("Registering unit driver '%v'", unitKind)
	if err := project.RegisterUnitFactory(&modDrv, unitKind); err != nil {
		log.Trace("Can't register unit driver '" + unitKind + "'.")
	}
}
