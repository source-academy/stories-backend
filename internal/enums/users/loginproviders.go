package userenums

type LoginProvider uint

const (
	LoginProviderGitHub LoginProvider = iota
	LoginProviderNUSNET
)

func (provider LoginProvider) String() string {
	switch provider {
	case LoginProviderGitHub:
		return "github"
	case LoginProviderNUSNET:
		return "luminus" // for legacy reasons
	}
	return "unknown"
}
