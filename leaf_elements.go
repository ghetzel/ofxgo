package ofxgo

// A list of all the leaf elements in OFX 1.0.3 (the last SGML version of the
// spec). These are all the elements that are possibly left unclosed, and which
// can have no children of their own. Fortunately these two sets of elements
// are the same. We use this list when parsing to remove ambiguities about
// element nesting.
//
// Generated using the following command with the 1.0.3 SPEC .dtd file:
//   # sed -rn 's/^<!ELEMENT\s+([A-Z0-9]+)\s+-\s+[oO]\s+%.*TYPE\s*>.*$/\t"\1",/p' *.dtd | sort
var ofxLeafElements = []string{
	"ACCESSKEY",
	"ACCRDINT",
	"ACCTID",
	"ACCTKEY",
	"ACCTREQUIRED",
	"ACCTTYPE",
	"ADDR1",
	"ADDR2",
	"ADDR3",
	"ADJAMT",
	"ADJDATE",
	"ADJDESC",
	"ADJNO",
	"APPID",
	"APPVER",
	"ASSETCLASS",
	"AUCTION",
	"AUTHTOKEN",
	"AUTHTOKENFIRST",
	"AUTHTOKENINFOURL",
	"AUTHTOKENLABEL",
	"AVAILACCTS",
	"AVAILCASH",
	"AVGCOSTBASIS",
	"BALAMT",
	"BALCLOSE",
	"BALDNLD",
	"BALMIN",
	"BALOPEN",
	"BALTYPE",
	"BANKID",
	"BILLREFINFO",
	"BRANCHID",
	"BROKERID",
	"BUYPOWER",
	"BUYTYPE",
	"CALLPRICE",
	"CALLTYPE",
	"CANADDPAYEE",
	"CANBILLPAY",
	"CANCELWND",
	"CANEMAIL",
	"CANMODMDLS",
	"CANMODPMTS",
	"CANMODXFERS",
	"CANNOTIFY",
	"CANPENDING",
	"CANRECUR",
	"CANSCHED",
	"CANUSEDESC",
	"CANUSERANGE",
	"CASESEN",
	"CHARTYPE",
	"CHECKING",
	"CHECKNUM",
	"CHGPINFIRST",
	"CHGUSERINFO",
	"CHKANDDEB",
	"CHKERROR",
	"CHKNUMEND",
	"CHKNUMSTART",
	"CHKSTATUS",
	"CITY",
	"CLIENTACTREQ",
	"CLIENTROUTING",
	"CLIENTUID",
	"CLIENTUIDREQ",
	"CLOSINGAVAIL",
	"CLTCOOKIE",
	"CODE",
	"COMMISSION",
	"CONFMSG",
	"CORRECTACTION",
	"CORRECTFITID",
	"COUNTRY",
	"COUPONFREQ",
	"COUPONRT",
	"CREDITLIMIT",
	"CSPHONE",
	"CURDEF",
	"CURRATE",
	"CURSYM",
	"DATEBIRTH",
	"DAYPHONE",
	"DAYSTOPAY",
	"DAYSWITH",
	"DEBADJ",
	"DEBTCLASS",
	"DEBTTYPE",
	"DENOMINATOR",
	"DEPANDCREDIT",
	"DESC",
	"DFLTDAYSTOPAY",
	"DIFFFIRSTPMT",
	"DIFFLASTPMT",
	"DOMXFERFEE",
	"DSCAMT",
	"DSCDATE",
	"DSCDESC",
	"DSCRATE",
	"DTACCTUP",
	"DTASOF",
	"DTAUCTION",
	"DTAVAIL",
	"DTCALL",
	"DTCHANGED",
	"DTCLIENT",
	"DTCLOSE",
	"DTCOUPON",
	"DTCREATED",
	"DTDUE",
	"DTEND",
	"DTEXPIRE",
	"DTINFOCHG",
	"DTMAT",
	"DTNEXT",
	"DTOPEN",
	"DTPLACED",
	"DTPMTDUE",
	"DTPMTPRC",
	"DTPOSTED",
	"DTPOSTEND",
	"DTPOSTSTART",
	"DTPRICEASOF",
	"DTPROFUP",
	"DTPURCHASE",
	"DTSERVER",
	"DTSETTLE",
	"DTSTART",
	"DTTRADE",
	"DTUSER",
	"DTXFERPRC",
	"DTXFERPRJ",
	"DTYIELDASOF",
	"DURATION",
	"EMAIL",
	"EVEPHONE",
	"EXTDPMTCHK",
	"EXTDPMTFOR",
	"FAXPHONE",
	"FEE",
	"FEEMSG",
	"FEES",
	"FIASSETCLASS",
	"FICERTID",
	"FID",
	"FIID",
	"FINALAMT",
	"FINAME",
	"FINCHG",
	"FIRSTNAME",
	"FITID",
	"FRACCASH",
	"FREQ",
	"FROM",
	"GAIN",
	"GENUSERKEY",
	"GETMIMESUP",
	"HASEXTDPMT",
	"HELDINACCT",
	"IDSCOPE",
	"INCBAL",
	"INCIMAGES",
	"INCOMETYPE",
	"INCOO",
	"INITIALAMT",
	"INTLXFERFEE",
	"INVACCTTYPE",
	"INVALIDACCTTYPE",
	"INVDATE",
	"INVDESC",
	"INVNO",
	"INVPAIDAMT",
	"INVTOTALAMT",
	"LANGUAGE",
	"LASTNAME",
	"LIMITPRICE",
	"LITMAMT",
	"LITMDESC",
	"LOAD",
	"LOSTSYNC",
	"MAILSUP",
	"MARGINBALANCE",
	"MARKDOWN",
	"MARKUP",
	"MAX",
	"MEMO",
	"MESSAGE",
	"MFACHALLENGEFIRST",
	"MFACHALLENGESUPT",
	"MFAPHRASEA",
	"MFAPHRASEID",
	"MFAPHRASELABEL",
	"MFTYPE",
	"MIDDLENAME",
	"MIN",
	"MINPMTDUE",
	"MINUNITS",
	"MKTGINFO",
	"MKTVAL",
	"MODELWND",
	"MODPENDING",
	"NAME",
	"NEWUNITS",
	"NEWUSERPASS",
	"NINSTS",
	"NONCE",
	"NUMERATOR",
	"OFXSEC",
	"OLDUNITS",
	"OODNLD",
	"OPTACTION",
	"OPTBUYTYPE",
	"OPTIONLEVEL",
	"OPTSELLTYPE",
	"OPTTYPE",
	"ORG",
	"PARVALUE",
	"PAYACCT",
	"PAYANDCREDIT",
	"PAYEEID",
	"PAYEELSTID",
	"PAYINSTRUCT",
	"PERCENT",
	"PHONE",
	"PINCH",
	"PMTBYADDR",
	"PMTBYPAYEEID",
	"PMTBYXFER",
	"PMTPRCCODE",
	"POSDNLD",
	"POSTALCODE",
	"POSTPROCWND",
	"POSTYPE",
	"PROCDAYSOFF",
	"PROCENDTM",
	"PURANDADV",
	"RATING",
	"RECSRVRTID",
	"REFNUM",
	"REFRESH",
	"REFRESHSUPT",
	"REINVCG",
	"REINVDIV",
	"REJECTIFMISSING",
	"RELFITID",
	"RELTYPE",
	"RESPFILEER",
	"RESTRICTION",
	"SECLISTRQDNLD",
	"SECNAME",
	"SECURED",
	"SECURITYNAME",
	"SELLALL",
	"SELLREASON",
	"SELLTYPE",
	"SESSCOOKIE",
	"SEVERITY",
	"SHORTBALANCE",
	"SHPERCTRCT",
	"SIC",
	"SIGNONREALM",
	"SPACES",
	"SPECIAL",
	"SPNAME",
	"SRVRTID",
	"STATE",
	"STOCKTYPE",
	"STOPPRICE",
	"STPCHKFEE",
	"STRIKEPRICE",
	"STSVIAMODS",
	"SUBACCT",
	"SUBACCTFROM",
	"SUBACCTSEC",
	"SUBACCTTO",
	"SUBJECT",
	"SUPTXDL",
	"SVC",
	"SVCSTATUS",
	"SWITCHALL",
	"SYNCMODE",
	"TAN",
	"TAXES",
	"TAXEXEMPT",
	"TAXID",
	"TEMPPASS",
	"TFERACTION",
	"TICKER",
	"TO",
	"TOKEN",
	"TOKENONLY",
	"TOTAL",
	"TOTALFEES",
	"TOTALINT",
	"TRANDNLD",
	"TRANSPSEC",
	"TRNAMT",
	"TRNTYPE",
	"TRNUID",
	"TSKEYEXPIRE",
	"TSPHONE",
	"TYPEDESC",
	"UNIQUEID",
	"UNIQUEIDTYPE",
	"UNITPRICE",
	"UNITS",
	"UNITSSTREET",
	"UNITSUSER",
	"UNITTYPE",
	"URL",
	"USEHTML",
	"USERCRED1",
	"USERCRED1LABEL",
	"USERCRED2",
	"USERCRED2LABEL",
	"USERID",
	"USERKEY",
	"USERPASS",
	"USPRODUCTTYPE",
	"VALUE",
	"VER",
	"WITHHOLDING",
	"XFERDAYSWITH",
	"XFERDEST",
	"XFERDFLTDAYSTOPAY",
	"XFERPRCCODE",
	"XFERSRC",
	"YIELD",
	"YIELDTOCALL",
	"YIELDTOMAT",
}
