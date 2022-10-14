package blockchain

type TxOutput struct {
	Value  int
	PubKey string
}

type TxInput struct {
	ID  []byte // references older output
	Out int    //index of output if there are many outputs
	Sig string
}

// to check if the account i.e data owns the info inside the output which is referenced by the input
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

// to check if the account i.e data owns info inside output
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
