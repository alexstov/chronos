package rtm

import (
	guuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// API operation ID
type API int

const (
	// APIUnknown unknown API
	APIUnknown API = iota
	// APIPostReplacementOffer post replacement offer
	APIPostReplacementOffer
	// APIPostReplacementPurchase post replacement purchase
	APIPostReplacementPurchase
	// APIGetReplacementVerification ...
	APIGetReplacementVerification
	// APIGetReplacementDetails ...
	APIGetReplacementDetails
	// APIGetReplacementProducts ...
	APIGetReplacementProducts
	// APIPostReplacementConfirmation ...
	APIPostReplacementConfirmation
	// APIProfileDisabler ...
	APIProfileDisabler
	// APIProfileEnabler ...
	APIProfileEnabler
	// APIGetHealthCheck ...
	APIGetHealthCheck
	// APIGetApplianceFilters ...
	APIGetApplianceFilters
	// APIPutSpecialPrices ...
	APIPutSpecialPrices
	// APILoadCatalog ...
	APILoadCatalog
)

func (a API) String() string {
	return [...]string{
		"Unknown",
		"PostReplacementOffer",
		"PostReplacementPurchase",
		"GetReplacementVerification",
		"GetReplacementDetails",
		"GetReplacementProducts",
		"PostReplacementConfirmation",
		"ProfileDisabler",
		"ProfileEnabler",
		"GetHealthCheck",
		"GetApplianceFilters",
		"PutSpecialPrices",
		"LoadCatalog"}[a]
}

// Context service context
type Context struct {
	Map map[guuid.UUID]*Capture
}

// Global RTM API Capture context
var Global *Context

// NewContext creates new RTM context
func NewContext() *Context {
	var ctx Context
	ctx.Map = map[guuid.UUID]*Capture{}

	return &ctx
}

// StartCapture generates new Capture uuid per API call,
// this is typically a router handle function call.
func (ctx *Context) StartCapture(api API) (capture *Capture, err error) {
	var capID guuid.UUID

	if capID, err = guuid.NewRandom(); err != nil {
		log.Error("RTM NewRandom failed to generate Capture ID for ", api) // TODO: get API name.
		return nil, err
	}

	cap := Capture{TxnID: capID, Type: api, OpsMap: make(map[OpsID]*Monitor)}
	Global.Map[capID] = &cap
	cap.initAllMonitors()
	cap.OpsMap[OpsTotal].Start(OpsTotal)

	return &cap, err
}

// EndCapture frees Capture RTM memory
func (ctx *Context) EndCapture(CaptureID guuid.UUID) {
	_, ok := ctx.Map[CaptureID]
	if ok {
		delete(ctx.Map, CaptureID)
		return
	}
	log.Warn("RTM trying to free unaccounted Capture ", CaptureID) // TODO: get API name.
}
