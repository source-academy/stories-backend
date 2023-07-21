package userenums

type LoginProvider uint

const (
	LoginProviderGitHub LoginProvider = iota
	LoginProviderNUSNET
)
