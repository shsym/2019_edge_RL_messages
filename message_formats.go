package edge_RL_messages

import "time"

// Redefine time units for uint64
const (
	MessageFormatVersion = 2
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
	ParamExp	float32
	ParamCoeff	float32
	ParamBias	float32
	LastUpdate	time.Time
}

func GetDefaultEdgeStat() EdgeStat{
	return EdgeStat{
		EdgeId:     0,
		EdgeMac:    "00:11:22:33:44:55",
		EdgeIp:     "123.123.123.123",
		ParamCoeff:	0.,
		ParamExp:	0.,
		ParamBias:	0.,
		PerfCpu:	0.,
		PerfMem:	0.,

		LastUpdate: time.Now(),
	}
}

type EdgeAdjacent struct{
	EdgeId			uint32
	AdjEdgeId1		uint32
	AdjEdgeAddr1	string
	AdjEdgeId2		uint32
	AdjEdgeAddr2	string
	AdjEdgeId3		uint32
	AdjEdgeAddr3	string
	AdjEdgeId4		uint32
	AdjEdgeAddr4	string
	AdjEdgeId5		uint32
	AdjEdgeAddr5	string
}

func GetDefaultEdgeAdjacent() EdgeAdjacent{
	return EdgeAdjacent{
		EdgeId:     0,
		AdjEdgeId1: 0,
		AdjEdgeAddr1: "",
		AdjEdgeId2: 0,
		AdjEdgeAddr2: "",
		AdjEdgeId3: 0,
		AdjEdgeAddr3: "",
		AdjEdgeId4: 0,
		AdjEdgeAddr4: "",
		AdjEdgeId5: 0,
		AdjEdgeAddr5: "",
	}
}

type EdgeReports struct {
	EdgeId		uint32
	ServiceName	string
	ProcTime	uint64
	ServSatis	float32		// 0 ~ 1
	ParamExp	float32
	ParamCoeff	float32
	ParamBias	float32
	Load 		float32
	TimeStamp	time.Time
}

func GetDefaultEdgeReport() EdgeReports{
	return EdgeReports{
		EdgeId:      0,
		ServiceName: "DefaultService",
		ProcTime:    500 * Millisecond,
		ServSatis:	 0.,
		ParamExp:    0.,
		ParamCoeff:  1.,
		ParamBias:   0.,
		Load:        0.5,
		TimeStamp:   time.Now(),
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

type EdgeStatReply struct{
	EdgeId		uint32
	Neighbors	string		// would be xml or json format
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

