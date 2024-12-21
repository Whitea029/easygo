package errors

type BizError struct {
	ShowedMsg   string `json:"showedMsg"`
	InternalMsg string `json:"internalMsg"`
	Err         error
}

func (b *BizError) Error() string {
	return b.ShowedMsg
}
