package email

type Generator interface {
	Generate() (string, error)
}
