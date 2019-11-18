package edge_RL_messages

import "time"

// Redefine time units for uint64
const (
	Nanosecond  		 = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

const(
	// Report from edges
	EdgeStatMsg = "EdgeStat"
	EdgeReportMsg = "EdgeReport"
	EdgeConnMsg = "EdgeConnection"
	ServiceStatMsg = "ServiceStat"

	// Request from edges
	EdgeAdjacentReq = "EdgeAdjacentReq"
)

type LifeCheck struct {
	LifeCheckMsg	string
}

type MessageWrapper struct{
	MessageType		string
	Data			interface{}
}

type EdgeStat struct{
	EdgeId		uint32
	EdgeMac		string
	EdgeIp		string
	PerfCpu		float32
	PerfMem		float32
	//BenchRes 	BenchmarkRecords
	//CtrlParams	ControlParameters
	ParamExp	float32
	ParamCoeff	float32
	ParamBias	float32
	LastUpdate	time.Time
}

func GetDefaultEdgeStat() EdgeStat{
	return EdgeStat{
		EdgeId:     2,
		EdgeMac:    "00:11:22:33:44:55",
		EdgeIp:     "165.132.106.7",
		ParamCoeff:	0.,
		ParamExp:	0.,
		ParamBias:	0.,
		PerfCpu:	0.,
		PerfMem:	0.,

		LastUpdate: time.Now(),
	}
}

type EdgeAdjacent struct{
	EdgeId		uint32
	AdjEdgeId1	uint32
	AdjEdgeId2	uint32
	AdjEdgeId3	uint32
	AdjEdgeId4	uint32
	AdjEdgeId5	uint32
}

func GetDefaultEdgeAdjacent() EdgeAdjacent{
	return EdgeAdjacent{
		EdgeId:     0,
		AdjEdgeId1: 0,
		AdjEdgeId2: 0,
		AdjEdgeId3: 0,
		AdjEdgeId4: 0,
		AdjEdgeId5: 0,
	}
}

type EdgeReports struct {
	EdgeId		uint32
	ServiceId	uint32
	ProcTime	uint64
	TimeStamp	time.Time
}

func GetDefaultEdgeReport() EdgeReports{
	return EdgeReports{
		EdgeId:    0,
		ServiceId: 0,
		ProcTime:  500 * Millisecond,
		TimeStamp: time.Now(),
	}
}

type ServiceStat struct{
	ServiceId		uint32
	ServiceName		string
	ServiceImage	string
	DefaultCpu		float32
	DefaultMem		float32
	DefaultTime		uint64
	LatencyReq		uint64
}

func GetDefaultServiceStat() ServiceStat{
	return ServiceStat{
		ServiceId:    1,
		ServiceName:  "DefaultService",
		ServiceImage: "gcr.io/knative-samples/helloworld-go",
		DefaultCpu:   10.,
		DefaultMem:   10.,
		DefaultTime:  100 * Millisecond,
		LatencyReq:   100 * Millisecond,
	}
}

type EdgeConnection struct{
	EdgeId		uint32
	OtherEdgeId	uint32
	Rtt			uint64
	TimeStamp	time.Time
}

func GetDefaultEdgeConnection() EdgeConnection{
	return EdgeConnection{
		EdgeId:      1,
		OtherEdgeId: 2,
		Rtt:         100 * Millisecond,
		TimeStamp:   time.Now(),
	}
}

//type BenchmarkRecords struct {
//	Records []BenchmarkResult
//}
//
//type BenchmarkResult struct {
//	Resource	string
//	Value		float32
//}
//
//var ResourceListV1 = [...]string{"CPU", "MEM"}
//
//type ControlParameters struct{
//	ParamExp	float32
//	ParamCoeff	float32
//	ParamBias	float32
//}

