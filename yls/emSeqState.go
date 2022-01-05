package yls

type EmSeqState int

const (
	SEQ_SUCCESS  EmSeqState = 0
	SEQ_NOMEMBER EmSeqState = 1001
	SEQ_NODISH   EmSeqState = 1002
)
