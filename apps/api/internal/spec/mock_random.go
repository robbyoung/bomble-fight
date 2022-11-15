package spec

type MockRandom struct {
	intToReturn int
}

func NewMockRandom(intToReturn int) *MockRandom {
	return &MockRandom{intToReturn: intToReturn}
}

func (r *MockRandom) RandInt(min int, max int) int {
	return r.intToReturn
}

func (r *MockRandom) RandArrayEntry(arr []string) string {
	i := r.RandInt(0, len(arr))

	return arr[i]
}
