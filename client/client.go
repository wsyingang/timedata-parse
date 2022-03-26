package client

var ClientMap map[string]Client
func init(){
	ClientMap=make(map[string]Client)
	// register entry for map
}
type Client interface {
	WriteData(tsData TimeSeriesData)
	WriteDataBatch(tsDataBatch TimeSeriesDataBatch)
	GetData(param QueryParam) TimeSeriesData
	GetDataBatch(batchParam QueryParamBatch) TimeSeriesDataBatch
	Server
}

func NewClient(){
	NewServer()
}
