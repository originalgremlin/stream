package configuration

type Configuration map[string]interface{}

func NewConfiguration() Configuration {
	// TODO: get conf from yml file and command-line args
	// TODO: get all modules and module values from conf
	// TODO: panic if configuration can't be loaded
	return make(map[string]interface{})
}

func Validate() (bool, error) {
	// TODO: return true if the configuration is valid
	// TODO: otherwise return false and the reason for the error
	return true, nil
}

func (c Configuration) Int(key string) int {
	val, _ := c[key].(int)
	return val
}

func (c Configuration) String(key string) string {
	val, _ := c[key].(string)
	return val
}
