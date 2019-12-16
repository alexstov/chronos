package rtm

import (
	guuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

// Capture benchmarks API Capture
type Capture struct {
	TxnID  guuid.UUID
	Type   API
	OpsMap map[OpsID]*Monitor
}

// Init adds all operations to the Capture map at once,
// required when defer is used to reference the operations map since
// defer value and parameters to the call are evaluate at the time of the call.
// See GoLand doumentation for defer.
func (txn *Capture) init() {
	switch txn.Type {
	case APIPostReplacementOffer:
		txn.OpsMap[OpsTotal] = NewMonitor(OpsTotal)
		txn.OpsMap[OpsCustomer] = NewMonitor(OpsCustomer)
		txn.OpsMap[OpsMongoDB] = NewMonitor(OpsMongoDB)
		txn.OpsMap[OpsSendEmail] = NewMonitor(OpsSendEmail)
		break
	case APIPostReplacementPurchase:
		txn.initAllMonitors()
		break
	}
}

func (txn *Capture) initAllMonitors() {
	txn.OpsMap[OpsTotal] = NewMonitor(OpsTotal)
	txn.OpsMap[OpsUnknown] = NewMonitor(OpsUnknown)
	txn.OpsMap[OpsCustomer] = NewMonitor(OpsCustomer)
	txn.OpsMap[OpsMongoDB] = NewMonitor(OpsMongoDB)
	txn.OpsMap[OpsSendEmail] = NewMonitor(OpsSendEmail)
	txn.OpsMap[OpsAddress] = NewMonitor(OpsAddress)
	txn.OpsMap[OpsPayment] = NewMonitor(OpsPayment)
	txn.OpsMap[OpsWeDeliver] = NewMonitor(OpsWeDeliver)
}

// Start starts operation RTM capture
func (txn *Capture) Start(opsID OpsID) {
	if txn == nil {
		return
	}

	var ops *Monitor
	var has bool

	if ops, has = txn.OpsMap[opsID]; !has {
		ops = NewMonitor(opsID)
		txn.OpsMap[opsID] = ops
	}
	ops.Start(opsID)
}

// Elapsed adds operation RTM elapsed time since the last start
func (txn *Capture) Elapsed(opsID OpsID) {
	var ops *Monitor
	var has bool

	if ops, has = txn.OpsMap[opsID]; !has {
		log.Errorf("RTM mismatched call for %v", opsID.String())
		//ops = NewMonitor(opsID)
	}
	ops.Elapsed()
}

// LogWarn logs Capture monitors as Warn message.
func (txn *Capture) LogWarn() {
	log.Warningf("RTM %v: TxnID <%v>, %v <%dms>, %v <%dms>, %v <%dms>, %v <%dms>, %v <%dms>, %v <%dms>, %v <%dms>, %v <%dms>.",
		txn.Type, txn.TxnID,
		OpsTotal.String(), txn.OpsMap[OpsTotal].Milliseconds(),
		OpsUnknown.String(), txn.OpsMap[OpsUnknown].Milliseconds(),
		OpsMongoDB.String(), txn.OpsMap[OpsMongoDB].Milliseconds(),
		OpsCustomer.String(), txn.OpsMap[OpsCustomer].Milliseconds(),
		OpsSendEmail.String(), txn.OpsMap[OpsSendEmail].Milliseconds(),
		OpsAddress.String(), txn.OpsMap[OpsAddress].Milliseconds(),
		OpsPayment.String(), txn.OpsMap[OpsPayment].Milliseconds(),
		OpsWeDeliver.String(), txn.OpsMap[OpsWeDeliver].Milliseconds())
}

// LogInfo logs Capture monitors as Info message.
func (txn *Capture) LogInfo() {
	log.Infof("RTM %v: TxnID <%v>, %v <%dms>, %v <%dms>, %v <%dms>, %v <%dms>, %v <%dms>, %v <%dms>, %v <%dms>, %v <%dms>.",
		txn.Type, txn.TxnID,
		OpsTotal.String(), txn.OpsMap[OpsTotal].Milliseconds(),
		OpsUnknown.String(), txn.OpsMap[OpsUnknown].Milliseconds(),
		OpsMongoDB.String(), txn.OpsMap[OpsMongoDB].Milliseconds(),
		OpsCustomer.String(), txn.OpsMap[OpsCustomer].Milliseconds(),
		OpsSendEmail.String(), txn.OpsMap[OpsSendEmail].Milliseconds(),
		OpsAddress.String(), txn.OpsMap[OpsAddress].Milliseconds(),
		OpsPayment.String(), txn.OpsMap[OpsPayment].Milliseconds(),
		OpsWeDeliver.String(), txn.OpsMap[OpsWeDeliver].Milliseconds())
}

// Finish captures elapsed monitors, logs Capture info, and ends it.
func (txn *Capture) Finish() {

	var calcTotal time.Duration
	for k, v := range txn.OpsMap {
		if v.running {
			if k != OpsTotal {
				log.Warnf("RTM monitor %v is still running. Capturing the monitor.", k.String())
			}

			txn.Elapsed(k)
		}

		if k != OpsTotal && k != OpsUnknown {
			calcTotal += v.elapsed
		}
	}

	txn.OpsMap[OpsUnknown].elapsed = txn.OpsMap[OpsTotal].elapsed - calcTotal

	txn.LogInfo()
	// Global.EndTransaction(txn.TxnID)
}
