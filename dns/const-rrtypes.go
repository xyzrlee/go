package dns

type RRType = uint16

const (
	RRTypeA          RRType = 1
	RRTypeNS         RRType = 2
	RRTypeMD         RRType = 3
	RRTypeMF         RRType = 4
	RRTypeCNAME      RRType = 5
	RRTypeSOA        RRType = 6
	RRTypeMB         RRType = 7
	RRTypeMG         RRType = 8
	RRTypeMR         RRType = 9
	RRTypeNULL       RRType = 10
	RRTypeWKS        RRType = 11
	RRTypePTR        RRType = 12
	RRTypeHINFO      RRType = 13
	RRTypeMINFO      RRType = 14
	RRTypeMX         RRType = 15
	RRTypeTXT        RRType = 16
	RRTypeRP         RRType = 17
	RRTypeAFSDB      RRType = 18
	RRTypeX25        RRType = 19
	RRTypeISDN       RRType = 20
	RRTypeRT         RRType = 21
	RRTypeNSAP       RRType = 22
	RRTypeNSAPPTR    RRType = 23
	RRTypeSIG        RRType = 24
	RRTypeKEY        RRType = 25
	RRTypePX         RRType = 26
	RRTypeGPOS       RRType = 27
	RRTypeAAAA       RRType = 28
	RRTypeLOC        RRType = 29
	RRTypeNXT        RRType = 30
	RRTypeEID        RRType = 31
	RRTypeNIMLOC     RRType = 32
	RRTypeSRV        RRType = 33
	RRTypeATMA       RRType = 34
	RRTypeNAPTR      RRType = 35
	RRTypeKX         RRType = 36
	RRTypeCERT       RRType = 37
	RRTypeA6         RRType = 38
	RRTypeDNAME      RRType = 39
	RRTypeSINK       RRType = 40
	RRTypeOPT        RRType = 41
	RRTypeAPL        RRType = 42
	RRTypeDS         RRType = 43
	RRTypeSSHFP      RRType = 44
	RRTypeIPSECKEY   RRType = 45
	RRTypeRRSIG      RRType = 46
	RRTypeNSEC       RRType = 47
	RRTypeDNSKEY     RRType = 48
	RRTypeDHCID      RRType = 49
	RRTypeNSEC3      RRType = 50
	RRTypeNSEC3PARAM RRType = 51
	RRTypeTLSA       RRType = 52
	RRTypeSMIMEA     RRType = 53
	RRTypeHIP        RRType = 55
	RRTypeNINFO      RRType = 56
	RRTypeRKEY       RRType = 57
	RRTypeTALINK     RRType = 58
	RRTypeCDS        RRType = 59
	RRTypeCDNSKEY    RRType = 60
	RRTypeOPENPGPKEY RRType = 61
	RRTypeCSYNC      RRType = 62
	RRTypeZONEMD     RRType = 63
	RRTypeSPF        RRType = 99
	RRTypeUINFO      RRType = 100
	RRTypeUID        RRType = 101
	RRTypeGID        RRType = 102
	RRTypeUNSPEC     RRType = 103
	RRTypeNID        RRType = 104
	RRTypeL32        RRType = 105
	RRTypeL64        RRType = 106
	RRTypeLP         RRType = 107
	RRTypeEUI48      RRType = 108
	RRTypeEUI64      RRType = 109
	RRTypeTKEY       RRType = 249
	RRTypeTSIG       RRType = 250
	RRTypeIXFR       RRType = 251
	RRTypeAXFR       RRType = 252
	RRTypeMAILB      RRType = 253
	RRTypeMAILA      RRType = 254
	RRTypeANY        RRType = 255
	RRTypeURI        RRType = 256
	RRTypeCAA        RRType = 257
	RRTypeAVC        RRType = 258
	RRTypeDOA        RRType = 259
	RRTypeTA         RRType = 32768
	RRTypeDLV        RRType = 32769
)

var RRTypeToString = map[RRType]string{
	RRTypeA:          "A",
	RRTypeNS:         "NS",
	RRTypeMD:         "MD",
	RRTypeMF:         "MF",
	RRTypeCNAME:      "CNAME",
	RRTypeSOA:        "SOA",
	RRTypeMB:         "MB",
	RRTypeMG:         "MG",
	RRTypeMR:         "MR",
	RRTypeNULL:       "NULL",
	RRTypeWKS:        "WKS",
	RRTypePTR:        "PTR",
	RRTypeHINFO:      "HINFO",
	RRTypeMINFO:      "MINFO",
	RRTypeMX:         "MX",
	RRTypeTXT:        "TXT",
	RRTypeRP:         "RP",
	RRTypeAFSDB:      "AFSDB",
	RRTypeX25:        "X25",
	RRTypeISDN:       "ISDN",
	RRTypeRT:         "RT",
	RRTypeNSAP:       "NSAP",
	RRTypeNSAPPTR:    "NSAP-PTR",
	RRTypeSIG:        "SIG",
	RRTypeKEY:        "KEY",
	RRTypePX:         "PX",
	RRTypeGPOS:       "GPOS",
	RRTypeAAAA:       "AAAA",
	RRTypeLOC:        "LOC",
	RRTypeNXT:        "NXT",
	RRTypeEID:        "EID",
	RRTypeNIMLOC:     "NIMLOC",
	RRTypeSRV:        "SRV",
	RRTypeATMA:       "ATMA",
	RRTypeNAPTR:      "NAPTR",
	RRTypeKX:         "KX",
	RRTypeCERT:       "CERT",
	RRTypeA6:         "A6",
	RRTypeDNAME:      "DNAME",
	RRTypeSINK:       "SINK",
	RRTypeOPT:        "OPT",
	RRTypeAPL:        "APL",
	RRTypeDS:         "DS",
	RRTypeSSHFP:      "SSHFP",
	RRTypeIPSECKEY:   "IPSECKEY",
	RRTypeRRSIG:      "RRSIG",
	RRTypeNSEC:       "NSEC",
	RRTypeDNSKEY:     "DNSKEY",
	RRTypeDHCID:      "DHCID",
	RRTypeNSEC3:      "NSEC3",
	RRTypeNSEC3PARAM: "NSEC3PARAM",
	RRTypeTLSA:       "TLSA",
	RRTypeSMIMEA:     "SMIMEA",
	RRTypeHIP:        "HIP",
	RRTypeNINFO:      "NINFO",
	RRTypeRKEY:       "RKEY",
	RRTypeTALINK:     "TALINK",
	RRTypeCDS:        "CDS",
	RRTypeCDNSKEY:    "CDNSKEY",
	RRTypeOPENPGPKEY: "OPENPGPKEY",
	RRTypeCSYNC:      "CSYNC",
	RRTypeZONEMD:     "ZONEMD",
	RRTypeSPF:        "SPF",
	RRTypeUINFO:      "UINFO",
	RRTypeUID:        "UID",
	RRTypeGID:        "GID",
	RRTypeUNSPEC:     "UNSPEC",
	RRTypeNID:        "NID",
	RRTypeL32:        "L32",
	RRTypeL64:        "L64",
	RRTypeLP:         "LP",
	RRTypeEUI48:      "EUI48",
	RRTypeEUI64:      "EUI64",
	RRTypeTKEY:       "TKEY",
	RRTypeTSIG:       "TSIG",
	RRTypeIXFR:       "IXFR",
	RRTypeAXFR:       "AXFR",
	RRTypeMAILB:      "MAILB",
	RRTypeMAILA:      "MAILA",
	RRTypeANY:        "ANY",
	RRTypeURI:        "URI",
	RRTypeCAA:        "CAA",
	RRTypeAVC:        "AVC",
	RRTypeDOA:        "DOA",
	RRTypeTA:         "TA",
	RRTypeDLV:        "DLV",
}
