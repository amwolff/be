package classifier

type Fingerprint struct{}

type Classifier interface {
	Do(f Fingerprint) (Fingerprint, bool)
}
