package counter

type Config struct {
	CaseSensitive bool `toml:"case_sensitive"`
	WithNumeric   bool `toml:"with_numeric"`
}
