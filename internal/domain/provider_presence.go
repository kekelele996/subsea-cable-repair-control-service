package domain

type ProviderHandle struct{ Provider SafetyProvider }

func (h ProviderHandle) Available() bool { return h.Provider != nil }
