package core

type ServType int64

type ServHub map[ServType]map[string]interface{}

const (
	SPIDER ServType = iota
	PIPE
	TASK
)

var service_types = []ServType{
	SPIDER,
	PIPE,
	TASK,
}
var Reg Registry

type Registry struct {
	hub ServHub
}

type Serv interface {
	GetName() string
	GetType() []ServType
}

func NewRegistry() Registry {
	hub := new(ServHub)
	for _, r_type := range service_types {
		(*hub)[r_type] = map[string]interface{}{}
	}
	return Registry{hub: *hub}
}

func (reg *Registry) GetByType(servType ServType) map[string]interface{} {
	return reg.hub[servType]
}

func (reg *Registry) Register(serv Serv) {
	for _, typ := range serv.GetType() {
		interfaces := reg.GetByType(typ)
		interfaces[serv.GetName()] = serv
	}
}
func (reg *Registry) UnRegister(serv Serv) {
	for _, typ := range serv.GetType() {
		interfaces := reg.GetByType(typ)
		delete(interfaces, serv.GetName())
	}
}

func init() {
	Reg = NewRegistry()
}
