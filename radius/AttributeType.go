package radius

//go:generate stringer -type=AttributeType
type AttributeType uint8

const (
	//start rfc2865
	_                                    = iota //drop the zero
	UserName               AttributeType = iota //1
	UserPassword           AttributeType = iota //2
	CHAPPassword           AttributeType = iota //3
	NASIPAddress           AttributeType = iota //4
	NASPort                AttributeType = iota //5
	ServiceType            AttributeType = iota //6
	FramedProtocol         AttributeType = iota //7
	FramedIPAddress        AttributeType = iota //8
	FramedIPNetmask        AttributeType = iota //9
	FramedRouting          AttributeType = iota //10
	FilterID               AttributeType = iota //11
	FramedMTU              AttributeType = iota //12
	FramedCompression      AttributeType = iota //13
	LoginIPHost            AttributeType = iota //14
	LoginService           AttributeType = iota //15
	LoginTCPPort           AttributeType = iota //16
	_                                    = iota //17 unassigned
	ReplyMessage           AttributeType = iota //18
	CallbackNumber         AttributeType = iota //19
	CallbackID             AttributeType = iota //20
	_                                    = iota //21 unassigned
	FramedRoute            AttributeType = iota //22
	FramedIPXNetwork       AttributeType = iota //23
	State                  AttributeType = iota //24
	Class                  AttributeType = iota //25
	VendorSpecific         AttributeType = iota
	SessionTimeout         AttributeType = iota
	IdleTimeout            AttributeType = iota
	TerminationAction      AttributeType = iota
	CalledStationID        AttributeType = iota
	CallingStationID       AttributeType = iota
	NASIdentifier          AttributeType = iota
	ProxyState             AttributeType = iota
	LoginLATService        AttributeType = iota
	LoginLATNode           AttributeType = iota
	LoginLATGroup          AttributeType = iota
	FramedAppleTalkLink    AttributeType = iota
	FramedAppleTalkNetwork AttributeType = iota
	FramedAppleTalkZone    AttributeType = iota
	AcctStatusType         AttributeType = iota
	AcctDelayTime          AttributeType = iota
	AcctInputOctets        AttributeType = iota
	AcctOutputOctets       AttributeType = iota
	AcctSessionID          AttributeType = iota
	AcctAuthentic          AttributeType = iota
	AcctSessionTime        AttributeType = iota
	AcctInputPackets       AttributeType = iota
	AcctOutputPackets      AttributeType = iota
	AcctTerminateCause     AttributeType = iota
	AcctMultiSessionID     AttributeType = iota
	AcctLinkCount          AttributeType = iota
	AcctInputGigawords     AttributeType = iota //52
	AcctOutputGigawords    AttributeType = iota
	Unassigned1            AttributeType = iota
	EventTimestamp         AttributeType = iota
	EgressVLANID           AttributeType = iota
	IngressFilters         AttributeType = iota
	EgressVLANName         AttributeType = iota
	UserPriorityTable      AttributeType = iota //59
	CHAPChallenge          AttributeType = 60
	NASPortType            AttributeType = 61
	PortLimit              AttributeType = 62
	LoginLATPort           AttributeType = 63
	//end rfc2865 rfc 2866
	TunnelType                   AttributeType = iota
	TunnelMediumType             AttributeType = iota
	TunnelClientEndpoint         AttributeType = iota
	TunnelServerEndpoint         AttributeType = iota
	AcctTunnelConnection         AttributeType = iota
	TunnelPassword               AttributeType = iota
	ARAPPassword                 AttributeType = iota
	ARAPFeatures                 AttributeType = iota
	ARAPZoneAccess               AttributeType = iota
	ARAPSecurity                 AttributeType = iota
	ARAPSecurityData             AttributeType = iota
	PasswordRetry                AttributeType = iota
	Prompt                       AttributeType = iota
	ConnectInfo                  AttributeType = iota
	ConfigurationToken           AttributeType = iota
	EAPMessage                   AttributeType = iota
	MessageAuthenticator         AttributeType = iota
	TunnelPrivateGroupID         AttributeType = iota
	TunnelAssignmentID           AttributeType = iota
	TunnelPreference             AttributeType = iota
	ARAPChallengeResponse        AttributeType = iota
	AcctInterimInterval          AttributeType = iota
	AcctTunnelPacketsLost        AttributeType = iota
	NASPortID                    AttributeType = iota
	FramedPool                   AttributeType = iota
	CUI                          AttributeType = iota
	TunnelClientAuthID           AttributeType = iota
	TunnelServerAuthID           AttributeType = iota
	NASFilterRule                AttributeType = iota
	Unassigned                   AttributeType = iota
	OriginatingLineInfo          AttributeType = iota
	NASIPv6Address               AttributeType = iota
	FramedInterfaceID            AttributeType = iota
	FramedIPv6Prefix             AttributeType = iota
	LoginIPv6Host                AttributeType = iota
	FramedIPv6Route              AttributeType = iota
	FramedIPv6Pool               AttributeType = iota
	ErrorCause                   AttributeType = iota
	EAPKeyName                   AttributeType = iota
	DigestResponse               AttributeType = iota
	DigestRealm                  AttributeType = iota
	DigestNonce                  AttributeType = iota
	DigestResponseAuth           AttributeType = iota
	DigestNextnonce              AttributeType = iota
	DigestMethod                 AttributeType = iota
	DigestURI                    AttributeType = iota
	DigestQop                    AttributeType = iota
	DigestAlgorithm              AttributeType = iota
	DigestEntityBodyHash         AttributeType = iota
	DigestCNonce                 AttributeType = iota
	DigestNonceCount             AttributeType = iota
	DigestUsername               AttributeType = iota
	DigestOpaque                 AttributeType = iota
	DigestAuthParam              AttributeType = iota
	DigestAKAAuts                AttributeType = iota
	DigestDomain                 AttributeType = iota
	DigestStale                  AttributeType = iota
	DigestHA1                    AttributeType = iota
	SIPAOR                       AttributeType = iota
	DelegatedIPv6Prefix          AttributeType = iota
	MIP6FeatureVector            AttributeType = iota
	MIP6HomeLinkPrefix           AttributeType = iota
	OperatorName                 AttributeType = iota
	LocationInformation          AttributeType = iota
	LocationData                 AttributeType = iota
	BasicLocationPolicyRules     AttributeType = iota
	ExtendedLocationPolicyRules  AttributeType = iota
	LocationCapable              AttributeType = iota
	RequestedLocationInfo        AttributeType = iota
	FramedManagementProtocol     AttributeType = iota
	ManagementTransportProtectio AttributeType = iota
	ManagementPolicyID           AttributeType = iota
	ManagementPrivilegeLevel     AttributeType = iota
	PKMSSCert                    AttributeType = iota
	PKMCACert                    AttributeType = iota
	PKMConfigSettings            AttributeType = iota
	PKMCryptosuiteList           AttributeType = iota
	PKMSAID                      AttributeType = iota
	PKMSADescriptor              AttributeType = iota
	PKMAuthKey                   AttributeType = iota
	DSLiteTunnelName             AttributeType = iota
	MobileNodeIdentifier         AttributeType = iota
	ServiceSelection             AttributeType = iota
	PMIP6HomeLMAIPv6Address      AttributeType = iota
	PMIP6VisitedLMAIPv6Address   AttributeType = iota
	PMIP6HomeLMAIPv4Address      AttributeType = iota
	PMIP6VisitedLMAIPv4Address   AttributeType = iota
	PMIP6HomeHNPrefix            AttributeType = iota
	PMIP6VisitedHNPrefix         AttributeType = iota
	PMIP6HomeInterfaceID         AttributeType = iota
	PMIP6VisitedInterfaceID      AttributeType = iota
	PMIP6HomeIPv4HoA             AttributeType = iota
	PMIP6VisitedIPv4HoA          AttributeType = iota
	PMIP6HomeDHCP4ServerAddress  AttributeType = iota
	PMIP6VisitedDHCP4ServerAddre AttributeType = iota
	PMIP6HomeDHCP6ServerAddress  AttributeType = iota
	PMIP6VisitedDHCP6ServerAddre AttributeType = iota
	UnassignedStart              AttributeType = 161
	UnassignedEnd                AttributeType = 191

	ExperimentalStart           AttributeType = 192
	ExperimentalEnd             AttributeType = 223
	ImplementationSpecificStart AttributeType = 224
	ImplementationSpecificEnd   AttributeType = 240
	ReservedStart               AttributeType = 241
	ReservedEnd                 AttributeType = 254
)
