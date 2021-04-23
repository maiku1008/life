package engine_test

import (
	"life/engine"
	"testing"
)

func TestGetNextGenerationState(t *testing.T) {
	t.Parallel()

	// location of the cell we wish to check
	x, y := 1, 1
	// size of the universe
	width, height := 3, 3

	tests := []struct {
		name          string
		state         bool
		neighbours    [][]int // locations of neighbours to check
		expectedState bool
	}{
		{
			name:          "A live cell with no neighbors",
			state:         true,
			neighbours:    [][]int{},
			expectedState: false,
		},
		{
			name:          "A live cell with fewer than two live neighbours dies",
			state:         true,
			neighbours:    [][]int{{x - 1, y - 1}},
			expectedState: false,
		},
		{
			name:          "A live cell with two live neighbours lives",
			state:         true,
			neighbours:    [][]int{{x - 1, y - 1}, {x - 1, y}},
			expectedState: true,
		},
		{
			name:          "A live cell with three live neighbours lives",
			state:         true,
			neighbours:    [][]int{{x - 1, y - 1}, {x - 1, y}, {x, y - 1}},
			expectedState: true,
		},
		{
			name:          "A live cell with more than three live neighbours dies",
			state:         true,
			neighbours:    [][]int{{x - 1, y - 1}, {x - 1, y}, {x, y - 1}, {x, y + 1}},
			expectedState: false,
		},
		{
			name:          "A dead cell with exactly three live neighbours becomes a live cell",
			state:         false,
			neighbours:    [][]int{{x - 1, y - 1}, {x - 1, y}, {x, y - 1}},
			expectedState: true,
		},
	}
	for _, tt := range tests {
		u := engine.NewUniverse(width, height)
		if tt.state {
			u.Resurrect(x, y)
		}
		for _, neighbour := range tt.neighbours {
			u.Resurrect(neighbour[0], neighbour[1])
		}
		gotState := u.GetNextGenerationState(x, y)
		if gotState != tt.expectedState {
			t.Errorf("%s: got=%t, expected=%t", tt.name, gotState, tt.expectedState)
		}
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	initial := 1
	u := engine.NewUniverse(8, 8)
	u.Init(initial)

	got := u.GetLiving()
	if initial != got {
		t.Errorf("error got=%d, initial=%d", got, initial)
	}
	u.Update()
	got = u.GetLiving()
	if initial == got {
		t.Errorf("error got=%d, initial=%d", got, initial)
	}
}
