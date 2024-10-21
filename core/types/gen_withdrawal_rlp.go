// Code generated by rlpgen. DO NOT EDIT.

package types

import "https://github.com/harbour-tech/go-ethereum-arbitrum/rlp"
import "io"

func (obj *Withdrawal) EncodeRLP(_w io.Writer) error {
	w := rlp.NewEncoderBuffer(_w)
	_tmp0 := w.List()
	w.WriteUint64(obj.Index)
	w.WriteUint64(obj.Validator)
	w.WriteBytes(obj.Address[:])
	w.WriteUint64(obj.Amount)
	w.ListEnd(_tmp0)
	return w.Flush()
}
