package expand

import (
	"bytes"
	"fmt"
	"github.com/stafiprotocol/go-substrate-rpc-client/scale"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"math"
)

/*
计算手续费
*/

type CalcFee struct {
	polynomial     *WeightToFeeCoefficient
	multiplier     int64
	per_byte_fee   int64
	base_fee       int64
	adjust_len_fee bool
}

func NewFee(meta *types.Metadata) (*CalcFee, error) {
	cf := new(CalcFee)
	md, err := NewMetadataExpand(meta)
	if err != nil {
		return nil, fmt.Errorf("new metadate expand error:%v", err)
	}
	_, tbfBytes, err := md.MV.GetConstants("TransactionPayment", "TransactionByteFee")
	if err != nil {
		return nil, fmt.Errorf("get TransactionByteFee constants error: %v", err)
	}
	var tbf types.U128
	decoder := scale.NewDecoder(bytes.NewReader(tbfBytes))
	err = decoder.Decode(&tbf)
	if err != nil {
		return nil, fmt.Errorf("decoder TransactionByteFee error: %v", err)
	}
	cf.per_byte_fee = int64(tbf.Uint64())

	_, wtfBytes, err := md.MV.GetConstants("TransactionPayment", "WeightToFee")
	if err != nil {
		return nil, fmt.Errorf("get WeightToFee constants error: %v", err)
	}
	decoder3 := scale.NewDecoder(bytes.NewReader(wtfBytes))
	vec := new(Vec)
	var wtfc WeightToFeeCoefficient
	err = vec.ProcessVec(*decoder3, wtfc)
	if err != nil {
		return nil, fmt.Errorf("decode WeightToFee error: %v", err)
	}
	vv := vec.Value[0]
	cf.polynomial = vv.(*WeightToFeeCoefficient)

	_, ebwBytes, err := md.MV.GetConstants("System", "ExtrinsicBaseWeight")
	if err != nil {
		return nil, fmt.Errorf("get ExtrinsicBaseWeight constants error: %v", err)
	}

	var w types.U32
	decoder2 := scale.NewDecoder(bytes.NewReader(ebwBytes))
	err = decoder2.Decode(&w)
	if err != nil {
		return nil, fmt.Errorf("decode ExtrinsicBaseWeight error: %v", err)
	}
	extrinsicBaseWeight := int64(w)
	cf.base_fee = cf.weight_to_fee(extrinsicBaseWeight)
	cf.adjust_len_fee = false //use V2
	return cf, nil
}

func NewCalcFee(polynomial *WeightToFeeCoefficient, extrinsic_base_weight, per_byte_fee, multiplier int64) *CalcFee {
	cf := new(CalcFee)
	cf.multiplier = multiplier
	cf.polynomial = polynomial
	cf.per_byte_fee = per_byte_fee
	cf.base_fee = cf.weight_to_fee(extrinsic_base_weight)
	cf.adjust_len_fee = false //v2-->false
	return cf
}

func (cf *CalcFee) SetMultiplier(multiplier int64) {
	cf.multiplier = multiplier
}

func (cf *CalcFee) CalcPartialFee(weight, len int64) int64 {
	unadjusted_len_fee := cf.per_byte_fee * len
	unadjusted_weight_fee := cf.weight_to_fee(weight)
	var (
		len_fee, adjustable_fee int64
	)
	if cf.adjust_len_fee {
		len_fee = 0
		adjustable_fee = unadjusted_len_fee + unadjusted_weight_fee
	} else {
		len_fee = unadjusted_len_fee
		adjustable_fee = unadjusted_weight_fee
	}
	fmt.Println("adjustable_fee: ", adjustable_fee)
	adjusted_fee := cf.calc(adjustable_fee)
	var result int64
	result += cf.base_fee
	result += len_fee
	result += adjusted_fee
	fmt.Println("BaseFee: ", cf.base_fee)
	fmt.Println("LengthFee: ", len_fee)
	fmt.Println("AdjuestFee: ", adjusted_fee)
	return result
}

func (cf *CalcFee) calc(adjustable_fee int64) int64 {
	//m:=big.NewInt(cf.multiplier)
	//af:=big.NewInt(adjustable_fee)
	//x:=m.Mod(m,af)
	//
	return int64(math.Mod(float64(cf.multiplier), float64(adjustable_fee)))
}
func (cf *CalcFee) weight_to_fee(weight int64) int64 {
	if cf.polynomial == nil {
		return 0
	}
	weight = int64(math.Pow(float64(weight), float64(cf.polynomial.Degree)))
	frac := int64(cf.polynomial.CoeffFrac.Value) * weight
	integer := cf.polynomial.CoeffInteger.Int64() * weight
	var acc int64
	if cf.polynomial.Negative {
		acc -= frac
		acc -= integer
	} else {
		acc += frac
		acc += integer
	}
	return acc
}
