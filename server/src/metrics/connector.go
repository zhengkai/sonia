package metrics

var (
	connectorReq  = newCounter(`connectorReq`, `Number of requests to the connector`)
	connectorRsp  = newCounter(`connectorRsp`, `Number of responses from the connector`)
	connectorFail = newCounter(`connectorFail`, `Number of failed requests to the connector`)
)

// ConnectorReq  ...
func ConnectorReq() {
	connectorReq.Inc()
}

// ConnectorRsp ...
func ConnectorRsp() {
	connectorRsp.Inc()
}

// ConnectorFail ...
func ConnectorFail() {
	connectorFail.Inc()
}
