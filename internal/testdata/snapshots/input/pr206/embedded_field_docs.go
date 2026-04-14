package pr206

// Base provides shared fields.
type Base struct {
	// ID uniquely identifies the entity.
	ID int
}

type Container struct {
	// Base is embedded to inherit shared fields.
	Base

	// Extra is a container-specific field.
	Extra string
}
