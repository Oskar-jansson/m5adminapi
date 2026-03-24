package models

type Event struct {
	System *SystemEvent
	Alarm  *AlarmEvent
	Area   *AreaEvent
	Ping   *PingEvent
	Error  *ErrorEvent

	// internal sdk message. Does not come from api
	InternalError error
}

type SystemEvent struct {
	Id               *uint32 `json:"id"`
	Domainid         *int32  `json:"domainId"`
	Timestamp        *string `json:"timestamp"`
	Message          *string `json:"message"`
	Eventtype        *uint32 `json:"eventtype"`
	UnitId           *uint32 `json:"unitid"`
	Evoptions        *uint32 `json:"evoptions"`
	Group            *string `json:"group"`
	Priority         *uint32 `json:"priority"`
	Unitnumber       *uint32 `json:"unitnumber"`
	Ionumber         *uint32 `json:"ionumber"`
	Domainname       *string `json:"domainname"`
	Extensionnumber  *uint32 `json:"extensionnumber"`
	Globalunitnumber *uint32 `json:"globalunitnumber"`
	Position         *string `json:"position"`
	Unitname         *string `json:"unitname"`
}

type AlarmEvent struct {
	Id               *uint32 `json:"id"`
	TimeStamp        *string `json:"timestamp"`
	Unitid           *uint32 `json:"unitid"`
	Domainid         *int32  `json:"domainid"`
	Group            *string `json:"group"`
	Priority         *uint32 `json:"priority"`
	Alarmtypetext    *string `json:"alarmtypetext"`
	IoNumber         *uint32 `json:"ionumber"`
	Alarmareaid      *int32  `json:"alarmareaid"`
	AlarmState       *uint32 `json:"alarmstate"`
	Alarmtext        *string `json:"alarmtext"`
	BaseType         *uint32 `json:"basetype"`
	AlarmType        *uint32 `json:"alarmtype"`
	DomainName       *string `json:"domainname"`
	Extensionnumber  *uint32 `json:"Extensionnumber"`
	Globalunitnumber *uint32 `json:"globalunitnumber"`
	Position         *string `json:"position"`
	UnitName         *string `json:"unitname"`
	Unitnumber       *uint32 `json:"unitnumber"`
}

type AreaEvent struct {
	Alarmareastate   *uint32 `json:"Alarmareastate"`
	Alarmoptions     *uint32 `json:"alarmoptions"`
	Alarmstateparams *uint64 `json:"alarmstateparams"`
	Armingerror      *uint32 `json:"armingerror"`
	Depzonenr        *uint32 `json:"depzonenr"`
	Zonenumber       *uint32 `json:"zonenumber"`
}

type PingEvent struct {
	Count *uint64 `json:"count"`
	Time  *string `json:"time"`
}

type ErrorEvent struct {
	Errorcode    *int    `json:"errorcode"`
	Errormessage *string `json:"errormessage"`
}
