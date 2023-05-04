package core

type RegisterType int64

const (
	SPIDER RegisterType = iota
	PIPE
)

var register_types = []RegisterType{
	SPIDER,
	PIPE,
}

type RegisterCenter struct {
	pool map[RegisterType]map[string]interface{}
}

func NewRC() RegisterCenter {
	pool := new(map[RegisterType]map[string]interface{})
	for _, r_type := range register_types {
		(*pool)[r_type] = map[string]interface{}{}
	}
	return RegisterCenter{pool: *pool}
}

func (rc *RegisterCenter) GetByType(Type RegisterType) map[string]interface{} {
	return rc.pool[Type]
}

var rc = NewRC()

func Register(Type RegisterType, name string, inter interface{}) {
	interfaces := rc.GetByType(Type)
	interfaces[name] = inter
}
