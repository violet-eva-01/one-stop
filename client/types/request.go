// Package types @author: Violet-Eva @date  : 2025/6/10 @notes :
package types

type ReqBody struct {
	AppKey           string   `json:"appKey"`
	Sign             string   `json:"sign"`
	WorksheetId      string   `json:"worksheetId,omitempty"`
	ViewId           string   `json:"viewId,omitempty"`
	RowId            string   `json:"rowId,omitempty"`
	PageSize         int      `json:"pageSize,omitempty"`
	PageIndex        int      `json:"pageIndex,omitempty"`
	ListType         int      `json:"listType,omitempty"`
	Controls         []string `json:"controls,omitempty"`
	Filters          []Filter `json:"filters,omitempty"`
	SortId           string   `json:"sortId,omitempty"`
	Asc              bool     `json:"isAsc,omitempty"`
	NotGetTotal      bool     `json:"notGetTotal,omitempty"`
	UseControlId     bool     `json:"useControlId,omitempty"`
	GetSystemControl bool     `json:"getSystemControl,omitempty"`
}

func NewReqBody() *ReqBody {
	return &ReqBody{}
}

func (req *ReqBody) SetAppKey(appKey string) *ReqBody {
	req.AppKey = appKey
	return req
}

func (req *ReqBody) SetSign(sign string) *ReqBody {
	req.Sign = sign
	return req
}

func (req *ReqBody) SetWorksheetId(worksheetId string) *ReqBody {
	req.WorksheetId = worksheetId
	return req
}

func (req *ReqBody) SetViewId(viewId string) *ReqBody {
	req.ViewId = viewId
	return req
}

func (req *ReqBody) SetRowId(rowId string) *ReqBody {
	req.RowId = rowId
	return req
}

func (req *ReqBody) SetPageSize(pageSize int) *ReqBody {
	req.PageSize = pageSize
	return req
}

func (req *ReqBody) SetPageIndex(pageIndex int) *ReqBody {
	req.PageIndex = pageIndex
	return req
}

func (req *ReqBody) SetListType(listType int) *ReqBody {
	req.ListType = listType
	return req
}

func (req *ReqBody) SetControls(controls ...string) *ReqBody {
	req.Controls = controls
	return req
}

func (req *ReqBody) SetFilters(filters ...Filter) *ReqBody {
	req.Filters = filters
	return req
}

func (req *ReqBody) SetSortId(sortId string) *ReqBody {
	req.SortId = sortId
	return req
}

func (req *ReqBody) IsAsc() *ReqBody {
	req.Asc = true
	return req
}

func (req *ReqBody) IsSetNotGetTotal() *ReqBody {
	req.NotGetTotal = true
	return req
}

func (req *ReqBody) IsSetUseControlId() *ReqBody {
	req.UseControlId = true
	return req
}

func (req *ReqBody) IsGetSystemControl() *ReqBody {
	req.GetSystemControl = true
	return req
}
