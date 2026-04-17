package pr218

// Deprecated: Use NewGreeting instead.
const OldGreeting = "hello"

const NewGreeting = "hi"

// Deprecated: Use Add instead.
func OldAdd(a, b int) int {
	return a + b
}

func Add(a, b int) int {
	return a + b
}

// Deprecated: Use Server instead.
type OldServer struct {
	// Deprecated: Use Addr instead.
	Host string
	Addr string
}

type Server struct {
	Addr string
}

func UseDeprecated() {
	_ = OldGreeting
	_ = NewGreeting

	_ = OldAdd(1, 2)
	_ = Add(1, 2)

	_ = OldServer{}
	_ = Server{}
}
