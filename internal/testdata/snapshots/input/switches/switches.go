package switches

// CustomSwitch does the things in a switch
type CustomSwitch struct{}

// Something does some things... and stuff
func (c *CustomSwitch) Something() bool { return false }

func Switch(interfaceValue interface{}) bool {
	switch concreteValue := interfaceValue.(type) {
	case int:
		return concreteValue*3 > 10
	case bool:
		return !concreteValue
	case CustomSwitch:
		return concreteValue.Something()
	default:
		return false
	}
}
