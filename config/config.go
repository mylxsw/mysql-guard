package config

type Config struct {
	TestMode bool
	DataDir  string

	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string

	Killer              bool
	KillerMatchCommands []string
	KillerBusyTime      int

	DeadlockLogger bool

	AdanosServer []string
	AdanosToken  string
}
