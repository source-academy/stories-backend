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

func LoginProviderFromString(provider string) (LoginProvider, bool) {
	switch provider {
	case "github":
		return LoginProviderGitHub, true
	case "luminus":
		return LoginProviderNUSNET, true
	}
	// We fall back to NUSNET as default provider
	return LoginProviderNUSNET, false
}
