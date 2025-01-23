

package engine

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

//EngineAPIError是共識和執行之間的標準化錯誤訊息
//用戶端，也包含 Geth 可能包含的任何自訂錯誤訊息。
type EngineAPIError struct {
	code int
	msg  string
	err  error
}

func (e *EngineAPIError) ErrorCode() int { return e.code }
func (e *EngineAPIError) Error() string  { return e.msg }
func (e *EngineAPIError) ErrorData() interface{} {
	if e.err == nil {
		return nil
	}
	return struct {
		Error string `json:"err"`
	}{e.err.Error()}
}

// With 傳回帶有​​新的嵌入自訂資料欄位的錯誤副本。
func (e *EngineAPIError) With(err error) *EngineAPIError {
	return &EngineAPIError{
		code: e.code,
		msg:  e.msg,
		err:  err,
	}
}

var (
	_ rpc.Error     = new(EngineAPIError)
	_ rpc.DataError = new(EngineAPIError)
)

var (
    //引擎 API 在以下呼叫中傳回 VALID：
	//-newPayloadV1：如果有效負載已知或剛剛驗證並執行
	//-forkchoiceUpdateV1：如果鏈接受重組（如果它是陳舊的，可能會忽略）
	VALID = "VALID"

    //在以下呼叫中引擎 API 傳回 INVALID：
	//-newPayloadV1：如果有效負載無法在本機鏈頂部執行
	//-forkchoiceUpdateV1：如果新頭未知，則預合併或重組失敗
	INVALID = "INVALID"

    //SYNCING 由引擎 API 在以下呼叫中傳回：
	//-newPayloadV1：如果在活動同步之上接受有效負載
	//-forkchoiceUpdateV1：如果新的頭之前已經見過，但不是鏈的一部分
	SYNCING = "SYNCING"

    //ACCEPTED 由引擎 API 在以下呼叫中傳回：
	//-newPayloadV1：如果有效負載被接受，但未處理（側鏈）
	ACCEPTED = "ACCEPTED"

	GenericServerError       = &EngineAPIError{code: -32000, msg: "伺服器錯誤"}
	UnknownPayload           = &EngineAPIError{code: -38001, msg: "未知有效負載"}
	InvalidForkChoiceState   = &EngineAPIError{code: -38002, msg: "無效的 forkchoice 狀態"}
	InvalidPayloadAttributes = &EngineAPIError{code: -38003, msg: "無效負載屬性"}
	TooLargeRequest          = &EngineAPIError{code: -38004, msg: "請求太大"}
	InvalidParams            = &EngineAPIError{code: -32602, msg: "無效參數"}
	UnsupportedFork          = &EngineAPIError{code: -38005, msg: "無支撐前叉"}

	STATUS_INVALID         = ForkChoiceResponse{PayloadStatus: PayloadStatusV1{Status: INVALID}, PayloadID: nil}
	STATUS_SYNCING         = ForkChoiceResponse{PayloadStatus: PayloadStatusV1{Status: SYNCING}, PayloadID: nil}
	INVALID_TERMINAL_BLOCK = PayloadStatusV1{Status: INVALID, LatestValidHash: &common.Hash{}}
)
