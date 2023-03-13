package heavyapiload

type HeavyApiLoadRepository interface {
	Test() (string, error)
}
