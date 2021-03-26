package history

type History struct {
	Size    int
	Index   int
	History []string
}

func New(size int) *History {
	return &History{
		Size:    size,
		History: make([]string, size),
	}
}

func (h *History) Push(line string) {
	h.History[h.Index] = line
	h.Index++
	if h.Index == h.Size {
		h.Index = 0
	}
}

func (h *History) Get(i int) string {
	// if h.index-i-1 < 0 {
	//     return h.history[h.size+h.index-i-1]
	// }
	// return h.history[h.index-1-i]
	return h.History[h.Size-((h.Index+i)%h.Size)-1]
}
