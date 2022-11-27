package spec

type MockRandom struct {
	returnValues []int
	index        int
}

func NewMockRandom(returnValues []int) *MockRandom {
	return &MockRandom{returnValues: returnValues, index: 0}
}

func (r *MockRandom) RandInt(min int, max int) int {
	result := r.returnValues[r.index]

	r.index = (r.index + 1) % len(r.returnValues)

	return result
}

func (r *MockRandom) RandArrayEntry(arr []string) string {
	i := r.RandInt(0, len(arr))

	return arr[i]
}
