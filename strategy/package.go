package strategy

// Strategy is the authentication interface
type Strategy interface {
	GetName() string
	Register() bool
	Remove() bool
	Validate() bool
}
