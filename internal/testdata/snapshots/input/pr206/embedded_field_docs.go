package pr206

type Base struct{}

type Container struct {
	// Base is embedded to inherit shared fields.
	Base
}
