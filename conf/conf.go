package conf

type Configuration map[string]interface{}

func New() Configuration {
	// TODO: get conf from yml file and command-line args
	// TODO: get all modules and module values from conf
	return
}

func (c Configuration) Int(key string) int {
	val, _ := c[key].(int)
	return val
}

func (c Configuration) String(key string) string {
	val, _ := c[key].(string)
	return val
}

