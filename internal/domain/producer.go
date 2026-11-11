package domain

type Producer interface {
	Generate(c chan<- Job)
}
