/*
 * MetaSV API Spec
 *
 * API definition for MetaSV provided apis
 *
 * API version: 2.1.4
 * Contact: heqiming@metasv.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metasv

import (
	"encoding/json"
)

// SensibleNftSellV2Utxo Sensible nft sell v2 Utxo belongs to the specified address
type SensibleNftSellV2Utxo struct {
	// Address string of this utxo
	Address *string `json:"address,omitempty"`
	// Address calculated from contract hash(p2ch).
	ContractAddress *string `json:"contractAddress,omitempty"`
	// Txid for this utxo.
	Txid *string `json:"txid,omitempty"`
	// Output index for the Utxo.
	TxIndex *int32 `json:"txIndex,omitempty"`
	// Codehash of this utxo.
	CodeHash *string `json:"codeHash,omitempty"`
	// Genesis of this utxo.
	Genesis *string `json:"genesis,omitempty"`
	// The index of this NFT.
	TokenIndex *int64 `json:"tokenIndex,omitempty"`
	// the price of nft.
	Price *int64 `json:"price,omitempty"`
	// The address to receive fees.
	FeeAddressPkh *string `json:"feeAddressPkh,omitempty"`
	// feeRate
	FeeRate *int32 `json:"feeRate,omitempty"`
	// nftId
	NftId *string `json:"nftId,omitempty"`
	// The address pkh of seller
	SellerAddressPkh *string `json:"sellerAddressPkh,omitempty"`
	// Bsv value of the utxo(Irrelavant to token value)
	Satoshi *int64 `json:"satoshi,omitempty"`
	// Bsv value of the utxo(In string format)
	SatoshiString *string `json:"satoshiString,omitempty"`
	// The height of this utxo, -1 for unconfirmed utxo.
	Height *int32 `json:"height,omitempty"`
	// Is current nft transfered into sell contract, If not ready, the following fields will be null
	IsReady *bool `json:"isReady,omitempty"`
	// SensibleId of the token
	SensibleId *string `json:"sensibleId,omitempty"`
	// The metanet tx describing the nft.
	MetaTxid *string `json:"metaTxid,omitempty"`
	// Symbol of the token.
	MetaOutputIndex *int32 `json:"metaOutputIndex,omitempty"`
	// The total supply of this NFT.
	TokenSupply *int64 `json:"tokenSupply,omitempty"`
	// Flag used for paging
	Flag *string `json:"flag,omitempty"`
}

// NewSensibleNftSellV2Utxo instantiates a new SensibleNftSellV2Utxo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSensibleNftSellV2Utxo() *SensibleNftSellV2Utxo {
	this := SensibleNftSellV2Utxo{}
	return &this
}

// NewSensibleNftSellV2UtxoWithDefaults instantiates a new SensibleNftSellV2Utxo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSensibleNftSellV2UtxoWithDefaults() *SensibleNftSellV2Utxo {
	this := SensibleNftSellV2Utxo{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *SensibleNftSellV2Utxo) SetAddress(v string) {
	o.Address = &v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetContractAddress() string {
	if o == nil || o.ContractAddress == nil {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetContractAddressOk() (*string, bool) {
	if o == nil || o.ContractAddress == nil {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasContractAddress() bool {
	if o != nil && o.ContractAddress != nil {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *SensibleNftSellV2Utxo) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *SensibleNftSellV2Utxo) SetTxid(v string) {
	o.Txid = &v
}

// GetTxIndex returns the TxIndex field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetTxIndex() int32 {
	if o == nil || o.TxIndex == nil {
		var ret int32
		return ret
	}
	return *o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetTxIndexOk() (*int32, bool) {
	if o == nil || o.TxIndex == nil {
		return nil, false
	}
	return o.TxIndex, true
}

// HasTxIndex returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasTxIndex() bool {
	if o != nil && o.TxIndex != nil {
		return true
	}

	return false
}

// SetTxIndex gets a reference to the given int32 and assigns it to the TxIndex field.
func (o *SensibleNftSellV2Utxo) SetTxIndex(v int32) {
	o.TxIndex = &v
}

// GetCodeHash returns the CodeHash field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetCodeHash() string {
	if o == nil || o.CodeHash == nil {
		var ret string
		return ret
	}
	return *o.CodeHash
}

// GetCodeHashOk returns a tuple with the CodeHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetCodeHashOk() (*string, bool) {
	if o == nil || o.CodeHash == nil {
		return nil, false
	}
	return o.CodeHash, true
}

// HasCodeHash returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasCodeHash() bool {
	if o != nil && o.CodeHash != nil {
		return true
	}

	return false
}

// SetCodeHash gets a reference to the given string and assigns it to the CodeHash field.
func (o *SensibleNftSellV2Utxo) SetCodeHash(v string) {
	o.CodeHash = &v
}

// GetGenesis returns the Genesis field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetGenesis() string {
	if o == nil || o.Genesis == nil {
		var ret string
		return ret
	}
	return *o.Genesis
}

// GetGenesisOk returns a tuple with the Genesis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetGenesisOk() (*string, bool) {
	if o == nil || o.Genesis == nil {
		return nil, false
	}
	return o.Genesis, true
}

// HasGenesis returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasGenesis() bool {
	if o != nil && o.Genesis != nil {
		return true
	}

	return false
}

// SetGenesis gets a reference to the given string and assigns it to the Genesis field.
func (o *SensibleNftSellV2Utxo) SetGenesis(v string) {
	o.Genesis = &v
}

// GetTokenIndex returns the TokenIndex field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetTokenIndex() int64 {
	if o == nil || o.TokenIndex == nil {
		var ret int64
		return ret
	}
	return *o.TokenIndex
}

// GetTokenIndexOk returns a tuple with the TokenIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetTokenIndexOk() (*int64, bool) {
	if o == nil || o.TokenIndex == nil {
		return nil, false
	}
	return o.TokenIndex, true
}

// HasTokenIndex returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasTokenIndex() bool {
	if o != nil && o.TokenIndex != nil {
		return true
	}

	return false
}

// SetTokenIndex gets a reference to the given int64 and assigns it to the TokenIndex field.
func (o *SensibleNftSellV2Utxo) SetTokenIndex(v int64) {
	o.TokenIndex = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetPrice() int64 {
	if o == nil || o.Price == nil {
		var ret int64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetPriceOk() (*int64, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given int64 and assigns it to the Price field.
func (o *SensibleNftSellV2Utxo) SetPrice(v int64) {
	o.Price = &v
}

// GetFeeAddressPkh returns the FeeAddressPkh field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetFeeAddressPkh() string {
	if o == nil || o.FeeAddressPkh == nil {
		var ret string
		return ret
	}
	return *o.FeeAddressPkh
}

// GetFeeAddressPkhOk returns a tuple with the FeeAddressPkh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetFeeAddressPkhOk() (*string, bool) {
	if o == nil || o.FeeAddressPkh == nil {
		return nil, false
	}
	return o.FeeAddressPkh, true
}

// HasFeeAddressPkh returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasFeeAddressPkh() bool {
	if o != nil && o.FeeAddressPkh != nil {
		return true
	}

	return false
}

// SetFeeAddressPkh gets a reference to the given string and assigns it to the FeeAddressPkh field.
func (o *SensibleNftSellV2Utxo) SetFeeAddressPkh(v string) {
	o.FeeAddressPkh = &v
}

// GetFeeRate returns the FeeRate field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetFeeRate() int32 {
	if o == nil || o.FeeRate == nil {
		var ret int32
		return ret
	}
	return *o.FeeRate
}

// GetFeeRateOk returns a tuple with the FeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetFeeRateOk() (*int32, bool) {
	if o == nil || o.FeeRate == nil {
		return nil, false
	}
	return o.FeeRate, true
}

// HasFeeRate returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasFeeRate() bool {
	if o != nil && o.FeeRate != nil {
		return true
	}

	return false
}

// SetFeeRate gets a reference to the given int32 and assigns it to the FeeRate field.
func (o *SensibleNftSellV2Utxo) SetFeeRate(v int32) {
	o.FeeRate = &v
}

// GetNftId returns the NftId field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetNftId() string {
	if o == nil || o.NftId == nil {
		var ret string
		return ret
	}
	return *o.NftId
}

// GetNftIdOk returns a tuple with the NftId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetNftIdOk() (*string, bool) {
	if o == nil || o.NftId == nil {
		return nil, false
	}
	return o.NftId, true
}

// HasNftId returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasNftId() bool {
	if o != nil && o.NftId != nil {
		return true
	}

	return false
}

// SetNftId gets a reference to the given string and assigns it to the NftId field.
func (o *SensibleNftSellV2Utxo) SetNftId(v string) {
	o.NftId = &v
}

// GetSellerAddressPkh returns the SellerAddressPkh field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetSellerAddressPkh() string {
	if o == nil || o.SellerAddressPkh == nil {
		var ret string
		return ret
	}
	return *o.SellerAddressPkh
}

// GetSellerAddressPkhOk returns a tuple with the SellerAddressPkh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetSellerAddressPkhOk() (*string, bool) {
	if o == nil || o.SellerAddressPkh == nil {
		return nil, false
	}
	return o.SellerAddressPkh, true
}

// HasSellerAddressPkh returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasSellerAddressPkh() bool {
	if o != nil && o.SellerAddressPkh != nil {
		return true
	}

	return false
}

// SetSellerAddressPkh gets a reference to the given string and assigns it to the SellerAddressPkh field.
func (o *SensibleNftSellV2Utxo) SetSellerAddressPkh(v string) {
	o.SellerAddressPkh = &v
}

// GetSatoshi returns the Satoshi field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetSatoshi() int64 {
	if o == nil || o.Satoshi == nil {
		var ret int64
		return ret
	}
	return *o.Satoshi
}

// GetSatoshiOk returns a tuple with the Satoshi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetSatoshiOk() (*int64, bool) {
	if o == nil || o.Satoshi == nil {
		return nil, false
	}
	return o.Satoshi, true
}

// HasSatoshi returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasSatoshi() bool {
	if o != nil && o.Satoshi != nil {
		return true
	}

	return false
}

// SetSatoshi gets a reference to the given int64 and assigns it to the Satoshi field.
func (o *SensibleNftSellV2Utxo) SetSatoshi(v int64) {
	o.Satoshi = &v
}

// GetSatoshiString returns the SatoshiString field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetSatoshiString() string {
	if o == nil || o.SatoshiString == nil {
		var ret string
		return ret
	}
	return *o.SatoshiString
}

// GetSatoshiStringOk returns a tuple with the SatoshiString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetSatoshiStringOk() (*string, bool) {
	if o == nil || o.SatoshiString == nil {
		return nil, false
	}
	return o.SatoshiString, true
}

// HasSatoshiString returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasSatoshiString() bool {
	if o != nil && o.SatoshiString != nil {
		return true
	}

	return false
}

// SetSatoshiString gets a reference to the given string and assigns it to the SatoshiString field.
func (o *SensibleNftSellV2Utxo) SetSatoshiString(v string) {
	o.SatoshiString = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetHeight() int32 {
	if o == nil || o.Height == nil {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetHeightOk() (*int32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *SensibleNftSellV2Utxo) SetHeight(v int32) {
	o.Height = &v
}

// GetIsReady returns the IsReady field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetIsReady() bool {
	if o == nil || o.IsReady == nil {
		var ret bool
		return ret
	}
	return *o.IsReady
}

// GetIsReadyOk returns a tuple with the IsReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetIsReadyOk() (*bool, bool) {
	if o == nil || o.IsReady == nil {
		return nil, false
	}
	return o.IsReady, true
}

// HasIsReady returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasIsReady() bool {
	if o != nil && o.IsReady != nil {
		return true
	}

	return false
}

// SetIsReady gets a reference to the given bool and assigns it to the IsReady field.
func (o *SensibleNftSellV2Utxo) SetIsReady(v bool) {
	o.IsReady = &v
}

// GetSensibleId returns the SensibleId field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetSensibleId() string {
	if o == nil || o.SensibleId == nil {
		var ret string
		return ret
	}
	return *o.SensibleId
}

// GetSensibleIdOk returns a tuple with the SensibleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetSensibleIdOk() (*string, bool) {
	if o == nil || o.SensibleId == nil {
		return nil, false
	}
	return o.SensibleId, true
}

// HasSensibleId returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasSensibleId() bool {
	if o != nil && o.SensibleId != nil {
		return true
	}

	return false
}

// SetSensibleId gets a reference to the given string and assigns it to the SensibleId field.
func (o *SensibleNftSellV2Utxo) SetSensibleId(v string) {
	o.SensibleId = &v
}

// GetMetaTxid returns the MetaTxid field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetMetaTxid() string {
	if o == nil || o.MetaTxid == nil {
		var ret string
		return ret
	}
	return *o.MetaTxid
}

// GetMetaTxidOk returns a tuple with the MetaTxid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetMetaTxidOk() (*string, bool) {
	if o == nil || o.MetaTxid == nil {
		return nil, false
	}
	return o.MetaTxid, true
}

// HasMetaTxid returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasMetaTxid() bool {
	if o != nil && o.MetaTxid != nil {
		return true
	}

	return false
}

// SetMetaTxid gets a reference to the given string and assigns it to the MetaTxid field.
func (o *SensibleNftSellV2Utxo) SetMetaTxid(v string) {
	o.MetaTxid = &v
}

// GetMetaOutputIndex returns the MetaOutputIndex field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetMetaOutputIndex() int32 {
	if o == nil || o.MetaOutputIndex == nil {
		var ret int32
		return ret
	}
	return *o.MetaOutputIndex
}

// GetMetaOutputIndexOk returns a tuple with the MetaOutputIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetMetaOutputIndexOk() (*int32, bool) {
	if o == nil || o.MetaOutputIndex == nil {
		return nil, false
	}
	return o.MetaOutputIndex, true
}

// HasMetaOutputIndex returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasMetaOutputIndex() bool {
	if o != nil && o.MetaOutputIndex != nil {
		return true
	}

	return false
}

// SetMetaOutputIndex gets a reference to the given int32 and assigns it to the MetaOutputIndex field.
func (o *SensibleNftSellV2Utxo) SetMetaOutputIndex(v int32) {
	o.MetaOutputIndex = &v
}

// GetTokenSupply returns the TokenSupply field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetTokenSupply() int64 {
	if o == nil || o.TokenSupply == nil {
		var ret int64
		return ret
	}
	return *o.TokenSupply
}

// GetTokenSupplyOk returns a tuple with the TokenSupply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetTokenSupplyOk() (*int64, bool) {
	if o == nil || o.TokenSupply == nil {
		return nil, false
	}
	return o.TokenSupply, true
}

// HasTokenSupply returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasTokenSupply() bool {
	if o != nil && o.TokenSupply != nil {
		return true
	}

	return false
}

// SetTokenSupply gets a reference to the given int64 and assigns it to the TokenSupply field.
func (o *SensibleNftSellV2Utxo) SetTokenSupply(v int64) {
	o.TokenSupply = &v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *SensibleNftSellV2Utxo) GetFlag() string {
	if o == nil || o.Flag == nil {
		var ret string
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensibleNftSellV2Utxo) GetFlagOk() (*string, bool) {
	if o == nil || o.Flag == nil {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *SensibleNftSellV2Utxo) HasFlag() bool {
	if o != nil && o.Flag != nil {
		return true
	}

	return false
}

// SetFlag gets a reference to the given string and assigns it to the Flag field.
func (o *SensibleNftSellV2Utxo) SetFlag(v string) {
	o.Flag = &v
}

func (o SensibleNftSellV2Utxo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.ContractAddress != nil {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	if o.Txid != nil {
		toSerialize["txid"] = o.Txid
	}
	if o.TxIndex != nil {
		toSerialize["txIndex"] = o.TxIndex
	}
	if o.CodeHash != nil {
		toSerialize["codeHash"] = o.CodeHash
	}
	if o.Genesis != nil {
		toSerialize["genesis"] = o.Genesis
	}
	if o.TokenIndex != nil {
		toSerialize["tokenIndex"] = o.TokenIndex
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.FeeAddressPkh != nil {
		toSerialize["feeAddressPkh"] = o.FeeAddressPkh
	}
	if o.FeeRate != nil {
		toSerialize["feeRate"] = o.FeeRate
	}
	if o.NftId != nil {
		toSerialize["nftId"] = o.NftId
	}
	if o.SellerAddressPkh != nil {
		toSerialize["sellerAddressPkh"] = o.SellerAddressPkh
	}
	if o.Satoshi != nil {
		toSerialize["satoshi"] = o.Satoshi
	}
	if o.SatoshiString != nil {
		toSerialize["satoshiString"] = o.SatoshiString
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.IsReady != nil {
		toSerialize["isReady"] = o.IsReady
	}
	if o.SensibleId != nil {
		toSerialize["sensibleId"] = o.SensibleId
	}
	if o.MetaTxid != nil {
		toSerialize["metaTxid"] = o.MetaTxid
	}
	if o.MetaOutputIndex != nil {
		toSerialize["metaOutputIndex"] = o.MetaOutputIndex
	}
	if o.TokenSupply != nil {
		toSerialize["tokenSupply"] = o.TokenSupply
	}
	if o.Flag != nil {
		toSerialize["flag"] = o.Flag
	}
	return json.Marshal(toSerialize)
}

type NullableSensibleNftSellV2Utxo struct {
	value *SensibleNftSellV2Utxo
	isSet bool
}

func (v NullableSensibleNftSellV2Utxo) Get() *SensibleNftSellV2Utxo {
	return v.value
}

func (v *NullableSensibleNftSellV2Utxo) Set(val *SensibleNftSellV2Utxo) {
	v.value = val
	v.isSet = true
}

func (v NullableSensibleNftSellV2Utxo) IsSet() bool {
	return v.isSet
}

func (v *NullableSensibleNftSellV2Utxo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensibleNftSellV2Utxo(val *SensibleNftSellV2Utxo) *NullableSensibleNftSellV2Utxo {
	return &NullableSensibleNftSellV2Utxo{value: val, isSet: true}
}

func (v NullableSensibleNftSellV2Utxo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensibleNftSellV2Utxo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
