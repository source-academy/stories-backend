package userenums

type LoginProvider uint

const (
	LoginProviderGitHub LoginProvider = iota
	LoginProviderNUSNET
)

// We cannot name it String() because it will conflict with the String() method
func (provider LoginProvider) ToString() string {
	switch provider {
	case LoginProviderGitHub:
		return "github"
	case LoginProviderNUSNET:
		return "luminus" // for legacy reasons
	}
	return "unknown"
}
