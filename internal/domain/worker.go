package domain

type Worker interface {
	Do(Job) Result
}
