package service

// Properties is a structure for providing usage and descriptions to the root command
type Properties struct {
	Usage            string
	ShortDescription string
	LongDescription  string
}

// NewProperties creates a new set of properties for the service command
func NewProperties(usage string, shortDescription string, longDescription string) Properties {
	return Properties{
		usage,
		shortDescription,
		longDescription,
	}
}
