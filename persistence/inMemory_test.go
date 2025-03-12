package persistence

import "testing"

func TestInMemory(t *testing.T) {

	p := NewInMemory[Item]()

	SuitePersistencer(p, t)
	SuiteOptimisticLocking(p, t)
}
