package heavyapiload

type UseCase interface {
	Load() (string, error)
}
