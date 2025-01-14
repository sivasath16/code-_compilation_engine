package backend

type CodeRequest struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

type Task struct {
	ID       string `json:"id"`
	Code     string `json:"code"`
	Language string `json:"language"`
}

type ExecutionResult struct {
	ID     string `json:"id"`
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}
