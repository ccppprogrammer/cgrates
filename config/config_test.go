/*
Real-time Online/Offline Charging System (OCS) for Telecom & ISP environments
Copyright (C) ITsysCOM GmbH

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/
package config

import (
	"reflect"
	"testing"
	"time"

	"github.com/cgrates/cgrates/utils"
)

var cfg *CGRConfig
var err error

func TestCgrCfgConfigSharing(t *testing.T) {
	cfg, _ = NewDefaultCGRConfig()
	SetCgrConfig(cfg)
	cfgReturn := CgrConfig()
	if !reflect.DeepEqual(cfgReturn, cfg) {
		t.Errorf("Retrieved %v, Expected %v", cfgReturn, cfg)
	}
}

func TestCgrCfgLoadWithDefaults(t *testing.T) {
	JSN_CFG := `
{
"freeswitch_agent": {
	"enabled": true,				// starts SessionManager service: <true|false>
	"event_socket_conns":[					// instantiate connections to multiple FreeSWITCH servers
		{"address": "1.2.3.4:8021", "password": "ClueCon", "reconnects": 3, "alias":"123"},
		{"address": "1.2.3.5:8021", "password": "ClueCon", "reconnects": 5, "alias":"124"}
	],
},

}`
	eCgrCfg, err := NewDefaultCGRConfig()
	if err != nil {
		t.Error(err)
	}
	eCgrCfg.fsAgentCfg.Enabled = true
	eCgrCfg.fsAgentCfg.EventSocketConns = []*FsConnConfig{
		&FsConnConfig{Address: "1.2.3.4:8021", Password: "ClueCon", Reconnects: 3, Alias: "123"},
		&FsConnConfig{Address: "1.2.3.5:8021", Password: "ClueCon", Reconnects: 5, Alias: "124"},
	}
	if cgrCfg, err := NewCGRConfigFromJsonStringWithDefaults(JSN_CFG); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(eCgrCfg.fsAgentCfg, cgrCfg.fsAgentCfg) {
		t.Errorf("Expected: %+v, received: %+v", eCgrCfg.fsAgentCfg, cgrCfg.fsAgentCfg)
	}
}

func TestCgrCfgDataDBPortWithoutDynamic(t *testing.T) {
	JSN_CFG := `
{
"data_db": {
	"db_type": "mongo",
	}
}`

	if cgrCfg, err := NewCGRConfigFromJsonStringWithDefaults(JSN_CFG); err != nil {
		t.Error(err)
	} else if cgrCfg.DataDbType != utils.MONGO {
		t.Errorf("Expected: %+v, received: %+v", cgrCfg.DataDbType, utils.MONGO)
	} else if cgrCfg.DataDbPort != "6379" {
		t.Errorf("Expected: %+v, received: %+v", cgrCfg.DataDbPort, "6379")
	}
	JSN_CFG = `
{
"data_db": {
	"db_type": "internal",
	}
}`

	if cgrCfg, err := NewCGRConfigFromJsonStringWithDefaults(JSN_CFG); err != nil {
		t.Error(err)
	} else if cgrCfg.DataDbType != utils.INTERNAL {
		t.Errorf("Expected: %+v, received: %+v", cgrCfg.DataDbType, utils.INTERNAL)
	} else if cgrCfg.DataDbPort != "6379" {
		t.Errorf("Expected: %+v, received: %+v", cgrCfg.DataDbPort, "6379")
	}
}

func TestCgrCfgDataDBPortWithDymanic(t *testing.T) {
	JSN_CFG := `
{
"data_db": {
	"db_type": "mongo",
	"db_port": -1,
	}
}`

	if cgrCfg, err := NewCGRConfigFromJsonString(JSN_CFG); err != nil {
		t.Error(err)
	} else if cgrCfg.DataDbType != utils.MONGO {
		t.Errorf("Expected: %+v, received: %+v", cgrCfg.DataDbType, utils.MONGO)
	} else if cgrCfg.DataDbPort != "27017" {
		t.Errorf("Expected: %+v, received: %+v", cgrCfg.DataDbPort, "27017")
	}
	JSN_CFG = `
{
"data_db": {
	"db_type": "internal",
	"db_port": -1,
	}
}`

	if cgrCfg, err := NewCGRConfigFromJsonString(JSN_CFG); err != nil {
		t.Error(err)
	} else if cgrCfg.DataDbType != utils.INTERNAL {
		t.Errorf("Expected: %+v, received: %+v", cgrCfg.DataDbType, utils.INTERNAL)
	} else if cgrCfg.DataDbPort != "internal" {
		t.Errorf("Expected: %+v, received: %+v", cgrCfg.DataDbPort, "internal")
	}
}

func TestCgrCfgStorDBPortWithoutDynamic(t *testing.T) {
	JSN_CFG := `
{
"stor_db": {
	"db_type": "mongo",
	}
}`

	if cgrCfg, err := NewCGRConfigFromJsonStringWithDefaults(JSN_CFG); err != nil {
		t.Error(err)
	} else if cgrCfg.StorDBType != utils.MONGO {
		t.Errorf("Expected: %+v, received: %+v", cgrCfg.StorDBType, utils.MONGO)
	} else if cgrCfg.StorDBPort != "3306" {
		t.Errorf("Expected: %+v, received: %+v", cgrCfg.StorDBPort, "3306")
	}
}

func TestCgrCfgStorDBPortWithDymanic(t *testing.T) {
	JSN_CFG := `
{
"stor_db": {
	"db_type": "mongo",
	"db_port": -1,
	}
}`

	if cgrCfg, err := NewCGRConfigFromJsonString(JSN_CFG); err != nil {
		t.Error(err)
	} else if cgrCfg.StorDBType != utils.MONGO {
		t.Errorf("Expected: %+v, received: %+v", cgrCfg.StorDBType, utils.MONGO)
	} else if cgrCfg.StorDBPort != "27017" {
		t.Errorf("Expected: %+v, received: %+v", cgrCfg.StorDBPort, "27017")
	}
}

func TestCgrCfgListener(t *testing.T) {
	JSN_CFG := `
{
"listen": {
	"rpc_json": ":2012",
	"rpc_gob": ":2013",
	"http": ":2080",
	}
}`

	if cgrCfg, err := NewCGRConfigFromJsonString(JSN_CFG); err != nil {
		t.Error(err)
	} else if cgrCfg.RPCGOBTLSListen != "" {
		t.Errorf("Expected: %+v, received: %+v", cgrCfg.RPCGOBTLSListen, "")
	} else if cgrCfg.RPCJSONTLSListen != "" {
		t.Errorf("Expected: %+v, received: %+v", cgrCfg.RPCJSONTLSListen, "")
	}
}

func TestCgrCfgCDRC(t *testing.T) {
	JSN_RAW_CFG := `
{
"cdrc": [
	{
		"id": "*default",
		"enabled": true,							// enable CDR client functionality
		"content_fields":[							// import template, tag will match internally CDR field, in case of .csv value will be represented by index of the field value
			{"field_id": "ToR", "type": "*composed", "value": "~7:s/^(voice|data|sms|mms|generic)$/*$1/"},
			{"field_id": "AnswerTime", "type": "*composed", "value": "~1"},
			{"field_id": "Usage", "type": "*composed", "value": "~9:s/^(\\d+)$/${1}s/"},
		],
	},
],
}`
	eCgrCfg, _ := NewDefaultCGRConfig()
	eCgrCfg.CdrcProfiles["/var/spool/cgrates/cdrc/in"] = []*CdrcConfig{
		&CdrcConfig{
			ID:                       utils.META_DEFAULT,
			Enabled:                  true,
			DryRun:                   false,
			CdrsConns:                []*HaPoolConfig{&HaPoolConfig{Address: utils.MetaInternal}},
			CdrFormat:                "csv",
			FieldSeparator:           rune(','),
			DataUsageMultiplyFactor:  1024,
			Timezone:                 "",
			RunDelay:                 0,
			MaxOpenFiles:             1024,
			CdrInDir:                 "/var/spool/cgrates/cdrc/in",
			CdrOutDir:                "/var/spool/cgrates/cdrc/out",
			FailedCallsPrefix:        "missed_calls",
			CDRPath:                  utils.HierarchyPath([]string{""}),
			CdrSourceId:              "freeswitch_csv",
			Filters:                  []string{},
			Tenant:                   NewRSRParsersMustCompile("cgrates.org", true),
			ContinueOnSuccess:        false,
			PartialRecordCache:       time.Duration(10 * time.Second),
			PartialCacheExpiryAction: "*dump_to_file",
			HeaderFields:             make([]*FCTemplate, 0),
			ContentFields: []*FCTemplate{
				&FCTemplate{FieldId: "ToR", Type: utils.META_COMPOSED,
					Value: NewRSRParsersMustCompile("~7:s/^(voice|data|sms|mms|generic)$/*$1/", true)},
				&FCTemplate{FieldId: "AnswerTime", Type: utils.META_COMPOSED,
					Value: NewRSRParsersMustCompile("~1", true)},
				&FCTemplate{FieldId: "Usage", Type: utils.META_COMPOSED,
					Value: NewRSRParsersMustCompile("~9:s/^(\\d+)$/${1}s/", true)},
			},
			TrailerFields: make([]*FCTemplate, 0),
			CacheDumpFields: []*FCTemplate{
				&FCTemplate{ID: "CGRID", Type: utils.META_COMPOSED,
					Value: NewRSRParsersMustCompile(utils.DynamicDataPrefix+utils.CGRID, true)},
				&FCTemplate{ID: "RunID", Type: utils.META_COMPOSED,
					Value: NewRSRParsersMustCompile(utils.DynamicDataPrefix+utils.RunID, true)},
				&FCTemplate{ID: "TOR", Type: utils.META_COMPOSED,
					Value: NewRSRParsersMustCompile(utils.DynamicDataPrefix+utils.ToR, true)},
				&FCTemplate{ID: "OriginID", Type: utils.META_COMPOSED,
					Value: NewRSRParsersMustCompile(utils.DynamicDataPrefix+utils.OriginID, true)},
				&FCTemplate{ID: "RequestType", Type: utils.META_COMPOSED,
					Value: NewRSRParsersMustCompile(utils.DynamicDataPrefix+utils.RequestType, true)},
				&FCTemplate{ID: "Tenant", Type: utils.META_COMPOSED,
					Value: NewRSRParsersMustCompile(utils.DynamicDataPrefix+utils.Tenant, true)},
				&FCTemplate{ID: "Category", Type: utils.META_COMPOSED,
					Value: NewRSRParsersMustCompile(utils.DynamicDataPrefix+utils.Category, true)},
				&FCTemplate{ID: "Account", Type: utils.META_COMPOSED,
					Value: NewRSRParsersMustCompile(utils.DynamicDataPrefix+utils.Account, true)},
				&FCTemplate{ID: "Subject", Type: utils.META_COMPOSED,
					Value: NewRSRParsersMustCompile(utils.DynamicDataPrefix+utils.Subject, true)},
				&FCTemplate{ID: "Destination", Type: utils.META_COMPOSED,
					Value: NewRSRParsersMustCompile(utils.DynamicDataPrefix+utils.Destination, true)},
				&FCTemplate{ID: "SetupTime", Type: utils.META_COMPOSED,
					Value:  NewRSRParsersMustCompile(utils.DynamicDataPrefix+utils.SetupTime, true),
					Layout: "2006-01-02T15:04:05Z07:00"},
				&FCTemplate{ID: "AnswerTime", Type: utils.META_COMPOSED,
					Value:  NewRSRParsersMustCompile(utils.DynamicDataPrefix+utils.AnswerTime, true),
					Layout: "2006-01-02T15:04:05Z07:00"},
				&FCTemplate{ID: "Usage", Type: utils.META_COMPOSED,
					Value: NewRSRParsersMustCompile(utils.DynamicDataPrefix+utils.Usage, true)},
				&FCTemplate{ID: "Cost", Type: utils.META_COMPOSED,
					Value: NewRSRParsersMustCompile(utils.DynamicDataPrefix+utils.COST, true)},
			},
		},
	}
	if cgrCfg, err := NewCGRConfigFromJsonStringWithDefaults(JSN_RAW_CFG); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(eCgrCfg.CdrcProfiles, cgrCfg.CdrcProfiles) {
		t.Errorf("Expected: %+v,\n received: %+v",
			utils.ToJSON(eCgrCfg.CdrcProfiles["/var/spool/cgrates/cdrc/in"][0]),
			utils.ToJSON(cgrCfg.CdrcProfiles["/var/spool/cgrates/cdrc/in"][0]))
	}
}

func TestHttpAgentCfg(t *testing.T) {
	JSN_RAW_CFG := `
{
"http_agent": [
	{
		"id": "conecto1",
		"url": "/conecto",					// relative URL for requests coming in
		"sessions_conns": [
			{"address": "*internal"}		// connection towards SessionService
		],
		"tenant": "cgrates.org",
		"timezone": "",						// timezone for timestamps where not specified, empty for general defaults <""|UTC|Local|$IANA_TZ_DB>
		"request_payload":	"*url",			// source of input data <*url>
		"reply_payload":	"*xml",			// type of output data <*xml>
		"request_processors": [],
	}
],
}
	`
	eCgrCfg, _ := NewDefaultCGRConfig()
	eCgrCfg.httpAgentCfg = []*HttpAgentCfg{
		&HttpAgentCfg{
			ID:             "conecto1",
			Url:            "/conecto",
			Tenant:         NewRSRParsersMustCompile("cgrates.org", true),
			Timezone:       "",
			RequestPayload: utils.MetaUrl,
			ReplyPayload:   utils.MetaXml,
			SessionSConns: []*HaPoolConfig{
				&HaPoolConfig{Address: utils.MetaInternal}},
			RequestProcessors: nil,
		},
	}
	if cgrCfg, err := NewCGRConfigFromJsonStringWithDefaults(JSN_RAW_CFG); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(eCgrCfg.HttpAgentCfg(), cgrCfg.HttpAgentCfg()) {
		t.Errorf("Expected: %s, received: %s",
			utils.ToJSON(eCgrCfg.httpAgentCfg), utils.ToJSON(cgrCfg.httpAgentCfg))
	}
}

func TestCgrCfgLoadJSONDefaults(t *testing.T) {
	cgrCfg, err = NewDefaultCGRConfig()
	if err != nil {
		t.Error(err)
	}
}

func TestCgrCfgJSONDefaultsGeneral(t *testing.T) {
	if cgrCfg.HttpSkipTlsVerify != false {
		t.Error(cgrCfg.HttpSkipTlsVerify)
	}
	if cgrCfg.RoundingDecimals != 5 {
		t.Error(cgrCfg.RoundingDecimals)
	}
	if cgrCfg.DBDataEncoding != "msgpack" {
		t.Error(cgrCfg.DBDataEncoding)
	}
	if cgrCfg.TpExportPath != "/var/spool/cgrates/tpe" {
		t.Error(cgrCfg.TpExportPath)
	}
	if cgrCfg.PosterAttempts != 3 {
		t.Error(cgrCfg.PosterAttempts)
	}
	if cgrCfg.FailedPostsDir != "/var/spool/cgrates/failed_posts" {
		t.Error(cgrCfg.FailedPostsDir)
	}
	if cgrCfg.DefaultReqType != "*rated" {
		t.Error(cgrCfg.DefaultReqType)
	}
	if cgrCfg.DefaultCategory != "call" {
		t.Error(cgrCfg.DefaultCategory)
	}
	if cgrCfg.DefaultTenant != "cgrates.org" {
		t.Error(cgrCfg.DefaultTenant)
	}
	if cgrCfg.DefaultTimezone != "Local" {
		t.Error(cgrCfg.DefaultTimezone)
	}
	if cgrCfg.ConnectAttempts != 3 {
		t.Error(cgrCfg.ConnectAttempts)
	}
	if cgrCfg.Reconnects != -1 {
		t.Error(cgrCfg.Reconnects)
	}
	if cgrCfg.ConnectTimeout != 1*time.Second {
		t.Error(cgrCfg.ConnectTimeout)
	}
	if cgrCfg.ReplyTimeout != 2*time.Second {
		t.Error(cgrCfg.ReplyTimeout)
	}
	if cgrCfg.ResponseCacheTTL != 0*time.Second {
		t.Error(cgrCfg.ResponseCacheTTL)
	}
	if cgrCfg.InternalTtl != 2*time.Minute {
		t.Error(cgrCfg.InternalTtl)
	}
	if cgrCfg.LockingTimeout != 0 {
		t.Error(cgrCfg.LockingTimeout)
	}
	if cgrCfg.Logger != utils.MetaSysLog {
		t.Error(cgrCfg.Logger)
	}
	if cgrCfg.LogLevel != 6 {
		t.Error(cgrCfg.LogLevel)
	}
	if cgrCfg.DigestSeparator != "," {
		t.Error(cgrCfg.DigestSeparator)
	}
	if cgrCfg.DigestEqual != ":" {
		t.Error(cgrCfg.DigestEqual)
	}
	if cgrCfg.TLSServerCerificate != "" {
		t.Error(cgrCfg.TLSServerCerificate)
	}
	if cgrCfg.TLSServerKey != "" {
		t.Error(cgrCfg.TLSServerKey)
	}
}

func TestCgrCfgJSONDefaultsListen(t *testing.T) {
	if cgrCfg.RPCJSONListen != "127.0.0.1:2012" {
		t.Error(cgrCfg.RPCJSONListen)
	}
	if cgrCfg.RPCGOBListen != "127.0.0.1:2013" {
		t.Error(cgrCfg.RPCGOBListen)
	}
	if cgrCfg.HTTPListen != "127.0.0.1:2080" {
		t.Error(cgrCfg.HTTPListen)
	}
	if cgrCfg.RPCJSONTLSListen != "127.0.0.1:2022" {
		t.Error(cgrCfg.RPCJSONListen)
	}
	if cgrCfg.RPCGOBTLSListen != "127.0.0.1:2023" {
		t.Error(cgrCfg.RPCGOBListen)
	}
	if cgrCfg.HTTPTLSListen != "127.0.0.1:2280" {
		t.Error(cgrCfg.HTTPListen)
	}
}

func TestCgrCfgJSONDefaultsjsnDataDb(t *testing.T) {
	if cgrCfg.DataDbType != "redis" {
		t.Error(cgrCfg.DataDbType)
	}
	if cgrCfg.DataDbHost != "127.0.0.1" {
		t.Error(cgrCfg.DataDbHost)
	}
	if cgrCfg.DataDbPort != "6379" {
		t.Error(cgrCfg.DataDbPort)
	}
	if cgrCfg.DataDbName != "10" {
		t.Error(cgrCfg.DataDbName)
	}
	if cgrCfg.DataDbUser != "cgrates" {
		t.Error(cgrCfg.DataDbUser)
	}
	if cgrCfg.DataDbPass != "" {
		t.Error(cgrCfg.DataDbPass)
	}
	if cgrCfg.LoadHistorySize != 10 {
		t.Error(cgrCfg.LoadHistorySize)
	}
}

func TestCgrCfgJSONDefaultsStorDB(t *testing.T) {
	if cgrCfg.StorDBType != "mysql" {
		t.Error(cgrCfg.StorDBType)
	}
	if cgrCfg.StorDBHost != "127.0.0.1" {
		t.Error(cgrCfg.StorDBHost)
	}
	if cgrCfg.StorDBPort != "3306" {
		t.Error(cgrCfg.StorDBPort)
	}
	if cgrCfg.StorDBName != "cgrates" {
		t.Error(cgrCfg.StorDBName)
	}
	if cgrCfg.StorDBUser != "cgrates" {
		t.Error(cgrCfg.StorDBUser)
	}
	if cgrCfg.StorDBPass != "" {
		t.Error(cgrCfg.StorDBPass)
	}
	if cgrCfg.StorDBMaxOpenConns != 100 {
		t.Error(cgrCfg.StorDBMaxOpenConns)
	}
	if cgrCfg.StorDBMaxIdleConns != 10 {
		t.Error(cgrCfg.StorDBMaxIdleConns)
	}
	Eslice := []string{}
	if !reflect.DeepEqual(cgrCfg.StorDBCDRSIndexes, Eslice) {
		t.Error(cgrCfg.StorDBCDRSIndexes)
	}
}

func TestCgrCfgJSONDefaultsRALs(t *testing.T) {
	eHaPoolcfg := []*HaPoolConfig{}

	if cgrCfg.RALsEnabled != false {
		t.Error(cgrCfg.RALsEnabled)
	}
	if !reflect.DeepEqual(cgrCfg.RALsThresholdSConns, eHaPoolcfg) {
		t.Error(cgrCfg.RALsThresholdSConns)
	}
	if !reflect.DeepEqual(cgrCfg.RALsCDRStatSConns, eHaPoolcfg) {
		t.Error(cgrCfg.RALsCDRStatSConns)
	}
	if !reflect.DeepEqual(cgrCfg.RALsPubSubSConns, eHaPoolcfg) {
		t.Error(cgrCfg.RALsPubSubSConns)
	}
	if !reflect.DeepEqual(cgrCfg.RALsUserSConns, eHaPoolcfg) {
		t.Error(cgrCfg.RALsUserSConns)
	}
	if !reflect.DeepEqual(cgrCfg.RALsAliasSConns, eHaPoolcfg) {
		t.Error(cgrCfg.RALsAliasSConns)
	}
	if cgrCfg.RpSubjectPrefixMatching != false {
		t.Error(cgrCfg.RpSubjectPrefixMatching)
	}
	if cgrCfg.LcrSubjectPrefixMatching != false {
		t.Error(cgrCfg.LcrSubjectPrefixMatching)
	}
	eMaxCU := map[string]time.Duration{
		utils.ANY:   time.Duration(189 * time.Hour),
		utils.VOICE: time.Duration(72 * time.Hour),
		utils.DATA:  time.Duration(107374182400),
		utils.SMS:   time.Duration(10000),
	}
	if !reflect.DeepEqual(eMaxCU, cgrCfg.RALsMaxComputedUsage) {
		t.Errorf("Expecting: %+v, received: %+v", eMaxCU, cgrCfg.RALsMaxComputedUsage)
	}
}

func TestCgrCfgJSONDefaultsScheduler(t *testing.T) {
	if cgrCfg.SchedulerEnabled != false {
		t.Error(cgrCfg.SchedulerEnabled)
	}
}

func TestCgrCfgJSONDefaultsCDRS(t *testing.T) {
	eHaPoolCfg := []*HaPoolConfig{}
	var eCdrExtr []*utils.RSRField
	if cgrCfg.CDRSEnabled != false {
		t.Error(cgrCfg.CDRSEnabled)
	}
	if !reflect.DeepEqual(eCdrExtr, cgrCfg.CDRSExtraFields) {
		t.Errorf(" expecting: %+v, received: %+v", eCdrExtr, cgrCfg.CDRSExtraFields)
	}
	if cgrCfg.CDRSStoreCdrs != true {
		t.Error(cgrCfg.CDRSStoreCdrs)
	}
	if cgrCfg.CDRScdrAccountSummary != false {
		t.Error(cgrCfg.CDRScdrAccountSummary)
	}
	if cgrCfg.CDRSSMCostRetries != 5 {
		t.Error(cgrCfg.CDRSSMCostRetries)
	}
	if !reflect.DeepEqual(cgrCfg.CDRSRaterConns, []*HaPoolConfig{&HaPoolConfig{Address: "*internal"}}) {
		t.Error(cgrCfg.CDRSRaterConns)
	}
	if !reflect.DeepEqual(cgrCfg.CDRSChargerSConns, eHaPoolCfg) {
		t.Error(cgrCfg.CDRSChargerSConns)
	}
	if !reflect.DeepEqual(cgrCfg.CDRSPubSubSConns, eHaPoolCfg) {
		t.Error(cgrCfg.CDRSPubSubSConns)
	}
	if !reflect.DeepEqual(cgrCfg.CDRSAttributeSConns, eHaPoolCfg) {
		t.Error(cgrCfg.CDRSAttributeSConns)
	}
	if !reflect.DeepEqual(cgrCfg.CDRSUserSConns, eHaPoolCfg) {
		t.Error(cgrCfg.CDRSUserSConns)
	}
	if !reflect.DeepEqual(cgrCfg.CDRSAliaseSConns, eHaPoolCfg) {
		t.Error(cgrCfg.CDRSAliaseSConns)
	}
	if !reflect.DeepEqual(cgrCfg.CDRSCDRStatSConns, eHaPoolCfg) {
		t.Error(cgrCfg.CDRSCDRStatSConns)
	}
	if !reflect.DeepEqual(cgrCfg.CDRSThresholdSConns, eHaPoolCfg) {
		t.Error(cgrCfg.CDRSThresholdSConns)
	}
	if !reflect.DeepEqual(cgrCfg.CDRSStatSConns, eHaPoolCfg) {
		t.Error(cgrCfg.CDRSStatSConns)
	}
	if cgrCfg.CDRSOnlineCDRExports != nil {
		t.Error(cgrCfg.CDRSOnlineCDRExports)
	}
}

func TestCgrCfgJSONLoadCDRS(t *testing.T) {
	JSN_RAW_CFG := `
{
"cdrs": {
	"enabled": true,
	"chargers_conns": [
		{"address": "*internal"}
	],
	"rals_conns": [
		{"address": "*internal"}			// address where to reach the Rater for cost calculation, empty to disable functionality: <""|*internal|x.y.z.y:1234>
	],
},
}
	`
	cgrCfg, err := NewCGRConfigFromJsonStringWithDefaults(JSN_RAW_CFG)
	if err != nil {
		t.Error(err)
	}
	if !cgrCfg.CDRSEnabled {
		t.Error(cgrCfg.CDRSEnabled)
	}
	if !reflect.DeepEqual(cgrCfg.CDRSChargerSConns,
		[]*HaPoolConfig{&HaPoolConfig{Address: utils.MetaInternal}}) {
		t.Error(cgrCfg.CDRSChargerSConns)
	}
	if !reflect.DeepEqual(cgrCfg.CDRSRaterConns,
		[]*HaPoolConfig{&HaPoolConfig{Address: utils.MetaInternal}}) {
		t.Error(cgrCfg.CDRSRaterConns)
	}
}

func TestCgrCfgJSONDefaultsCDRStats(t *testing.T) {
	if cgrCfg.CDRStatsEnabled != false {
		t.Error(cgrCfg.CDRStatsEnabled)
	}
	if cgrCfg.CDRStatsSaveInterval != 1*time.Minute {
		t.Error(cgrCfg.CDRStatsSaveInterval)
	}
}

func TestCgrCfgJSONDefaultsCdreProfiles(t *testing.T) {
	eFields := []*FCTemplate{}
	eContentFlds := []*FCTemplate{
		&FCTemplate{ID: "CGRID", Type: "*composed",
			Value: NewRSRParsersMustCompile("~CGRID", true)},
		&FCTemplate{ID: "RunID", Type: "*composed",
			Value: NewRSRParsersMustCompile("~RunID", true)},
		&FCTemplate{ID: "TOR", Type: "*composed",
			Value: NewRSRParsersMustCompile("~ToR", true)},
		&FCTemplate{ID: "OriginID", Type: "*composed",
			Value: NewRSRParsersMustCompile("~OriginID", true)},
		&FCTemplate{ID: "RequestType", Type: "*composed",
			Value: NewRSRParsersMustCompile("~RequestType", true)},
		&FCTemplate{ID: "Tenant", Type: "*composed",
			Value: NewRSRParsersMustCompile("~Tenant", true)},
		&FCTemplate{ID: "Category", Type: "*composed",
			Value: NewRSRParsersMustCompile("~Category", true)},
		&FCTemplate{ID: "Account", Type: "*composed",
			Value: NewRSRParsersMustCompile("~Account", true)},
		&FCTemplate{ID: "Subject", Type: "*composed",
			Value: NewRSRParsersMustCompile("~Subject", true)},
		&FCTemplate{ID: "Destination", Type: "*composed",
			Value: NewRSRParsersMustCompile("~Destination", true)},
		&FCTemplate{ID: "SetupTime", Type: "*composed",
			Value:  NewRSRParsersMustCompile("~SetupTime", true),
			Layout: "2006-01-02T15:04:05Z07:00"},
		&FCTemplate{ID: "AnswerTime", Type: "*composed",
			Value:  NewRSRParsersMustCompile("~AnswerTime", true),
			Layout: "2006-01-02T15:04:05Z07:00"},
		&FCTemplate{ID: "Usage", Type: "*composed",
			Value: NewRSRParsersMustCompile("~Usage", true)},
		&FCTemplate{ID: "Cost", Type: "*composed",
			Value:            NewRSRParsersMustCompile("~Cost", true),
			RoundingDecimals: 4},
	}
	eCdreCfg := map[string]*CdreConfig{
		"*default": {
			ExportFormat:        utils.MetaFileCSV,
			ExportPath:          "/var/spool/cgrates/cdre",
			Filters:             []string{},
			Tenant:              "cgrates.org",
			Synchronous:         false,
			Attempts:            1,
			FieldSeparator:      ',',
			UsageMultiplyFactor: map[string]float64{utils.ANY: 1.0},
			CostMultiplyFactor:  1.0,
			HeaderFields:        eFields,
			ContentFields:       eContentFlds,
			TrailerFields:       eFields,
		},
	}
	if !reflect.DeepEqual(cgrCfg.CdreProfiles, eCdreCfg) {
		t.Errorf("received: %+v, expecting: %+v", cgrCfg.CdreProfiles, eCdreCfg)
	}
}

func TestCgrCfgJSONDefaultsSMGenericCfg(t *testing.T) {
	eSessionSCfg := &SessionSCfg{
		Enabled:       false,
		ListenBijson:  "127.0.0.1:2014",
		ChargerSConns: []*HaPoolConfig{},
		RALsConns: []*HaPoolConfig{
			&HaPoolConfig{Address: "*internal"}},
		CDRsConns: []*HaPoolConfig{
			&HaPoolConfig{Address: "*internal"}},
		ResSConns:               []*HaPoolConfig{},
		ThreshSConns:            []*HaPoolConfig{},
		StatSConns:              []*HaPoolConfig{},
		SupplSConns:             []*HaPoolConfig{},
		AttrSConns:              []*HaPoolConfig{},
		SessionReplicationConns: []*HaPoolConfig{},
		DebitInterval:           0 * time.Second,
		MinCallDuration:         0 * time.Second,
		MaxCallDuration:         3 * time.Hour,
		SessionTTL:              0 * time.Second,
		SessionIndexes:          utils.StringMap{},
		ClientProtocol:          1.0,
		ChannelSyncInterval:     0,
	}
	if !reflect.DeepEqual(eSessionSCfg, cgrCfg.sessionSCfg) {
		t.Errorf("expecting: %s, received: %s",
			utils.ToJSON(eSessionSCfg), utils.ToJSON(cgrCfg.sessionSCfg))
	}

}
func TestCgrCfgJSONDefaultsCacheCFG(t *testing.T) {
	eCacheCfg := CacheConfig{
		utils.CacheDestinations: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheReverseDestinations: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheRatingPlans: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheRatingProfiles: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheLCRRules: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheCDRStatS: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheActions: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheActionPlans: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheAccountActionPlans: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheActionTriggers: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheSharedGroups: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheAliases: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheReverseAliases: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheDerivedChargers: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheTimings: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheResourceProfiles: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheResources: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheEventResources: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(1 * time.Minute), StaticTTL: false},
		utils.CacheStatQueueProfiles: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(1 * time.Minute), StaticTTL: false, Precache: false},
		utils.CacheStatQueues: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(1 * time.Minute), StaticTTL: false, Precache: false},
		utils.CacheThresholdProfiles: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheThresholds: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheFilters: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheSupplierProfiles: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheAttributeProfiles: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheChargerProfiles: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheResourceFilterIndexes: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheStatFilterIndexes: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheThresholdFilterIndexes: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheSupplierFilterIndexes: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheAttributeFilterIndexes: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
		utils.CacheChargerFilterIndexes: &CacheParamConfig{Limit: -1,
			TTL: time.Duration(0), StaticTTL: false, Precache: false},
	}

	if !reflect.DeepEqual(eCacheCfg, cgrCfg.CacheCfg()) {
		t.Errorf("received: %s, \nexpecting: %s",
			utils.ToJSON(eCacheCfg), utils.ToJSON(cgrCfg.CacheCfg()))
	}
}

func TestCgrCfgJSONDefaultsFsAgentConfig(t *testing.T) {
	eFsAgentCfg := &FsAgentConfig{
		Enabled: false,
		SessionSConns: []*HaPoolConfig{
			&HaPoolConfig{Address: "*internal"}},
		SubscribePark:       true,
		CreateCdr:           false,
		ExtraFields:         nil,
		EmptyBalanceContext: "",
		EmptyBalanceAnnFile: "",
		MaxWaitConnection:   2 * time.Second,
		EventSocketConns: []*FsConnConfig{
			&FsConnConfig{Address: "127.0.0.1:8021",
				Password: "ClueCon", Reconnects: 5, Alias: "127.0.0.1:8021"}},
	}

	if !reflect.DeepEqual(cgrCfg.fsAgentCfg, eFsAgentCfg) {
		t.Errorf("received: %+v, expecting: %+v", cgrCfg.fsAgentCfg, eFsAgentCfg)
	}
}

func TestCgrCfgJSONDefaultsKamAgentConfig(t *testing.T) {
	eKamAgentCfg := &KamAgentCfg{
		Enabled: false,
		SessionSConns: []*HaPoolConfig{
			&HaPoolConfig{Address: "*internal"}},
		CreateCdr: false,
		EvapiConns: []*KamConnConfig{
			&KamConnConfig{
				Address: "127.0.0.1:8448", Reconnects: 5}},
	}
	if !reflect.DeepEqual(cgrCfg.kamAgentCfg, eKamAgentCfg) {
		t.Errorf("received: %+v, expecting: %+v",
			utils.ToJSON(cgrCfg.kamAgentCfg), utils.ToJSON(eKamAgentCfg))
	}
}

func TestCgrCfgJSONDefaultssteriskAgentCfg(t *testing.T) {
	eAstAgentCfg := &AsteriskAgentCfg{
		Enabled: false,
		SessionSConns: []*HaPoolConfig{
			&HaPoolConfig{Address: "*internal"}},
		CreateCDR: false,
		AsteriskConns: []*AsteriskConnCfg{
			&AsteriskConnCfg{Address: "127.0.0.1:8088",
				User: "cgrates", Password: "CGRateS.org",
				ConnectAttempts: 3, Reconnects: 5}},
	}

	if !reflect.DeepEqual(cgrCfg.asteriskAgentCfg, eAstAgentCfg) {
		t.Errorf("received: %+v, expecting: %+v", cgrCfg.asteriskAgentCfg, eAstAgentCfg)
	}
}

func TestCgrCfgJSONDefaultsPubSubS(t *testing.T) {
	if cgrCfg.PubSubServerEnabled != false {
		t.Error(cgrCfg.PubSubServerEnabled)
	}
}

func TestCgrCfgJSONDefaultsAliasesS(t *testing.T) {
	if cgrCfg.AliasesServerEnabled != false {
		t.Error(cgrCfg.AliasesServerEnabled)
	}
}

func TestCgrCfgJSONDefaultsUserS(t *testing.T) {
	eStrSlc := []string{}
	if cgrCfg.UserServerEnabled != false {
		t.Error(cgrCfg.UserServerEnabled)
	}

	if !reflect.DeepEqual(cgrCfg.UserServerIndexes, eStrSlc) {
		t.Errorf("received: %+v, expecting: %+v", cgrCfg.UserServerIndexes, eStrSlc)
	}
}

func TestCgrCfgJSONDefaultFiltersCfg(t *testing.T) {
	eFiltersCfg := &FilterSCfg{
		StatSConns:     []*HaPoolConfig{},
		IndexedSelects: true,
	}
	if !reflect.DeepEqual(cgrCfg.filterSCfg, eFiltersCfg) {
		t.Errorf("received: %+v, expecting: %+v", cgrCfg.filterSCfg, eFiltersCfg)
	}
}

func TestCgrCfgJSONDefaultSAttributeSCfg(t *testing.T) {
	eAliasSCfg := &AttributeSCfg{
		Enabled:             false,
		StringIndexedFields: nil,
		PrefixIndexedFields: &[]string{},
		ProcessRuns:         1,
	}
	if !reflect.DeepEqual(eAliasSCfg, cgrCfg.attributeSCfg) {
		t.Errorf("received: %+v, expecting: %+v", eAliasSCfg, cgrCfg.attributeSCfg)
	}
}

func TestCgrCfgJSONDefaultSChargerSCfg(t *testing.T) {
	eChargerSCfg := &ChargerSCfg{
		Enabled:             false,
		AttributeSConns:     []*HaPoolConfig{},
		StringIndexedFields: nil,
		PrefixIndexedFields: &[]string{},
	}
	if !reflect.DeepEqual(eChargerSCfg, cgrCfg.chargerSCfg) {
		t.Errorf("received: %+v, expecting: %+v", eChargerSCfg, cgrCfg.chargerSCfg)
	}
}

func TestCgrCfgJSONDefaultsResLimCfg(t *testing.T) {
	eResLiCfg := &ResourceSConfig{
		Enabled:             false,
		ThresholdSConns:     []*HaPoolConfig{},
		StoreInterval:       0,
		StringIndexedFields: nil,
		PrefixIndexedFields: &[]string{},
	}
	if !reflect.DeepEqual(cgrCfg.resourceSCfg, eResLiCfg) {
		t.Errorf("expecting: %s, received: %s", utils.ToJSON(eResLiCfg), utils.ToJSON(cgrCfg.resourceSCfg))
	}

}

func TestCgrCfgJSONDefaultStatsCfg(t *testing.T) {
	eStatsCfg := &StatSCfg{
		Enabled:             false,
		StoreInterval:       0,
		ThresholdSConns:     []*HaPoolConfig{},
		StringIndexedFields: nil,
		PrefixIndexedFields: &[]string{},
	}
	if !reflect.DeepEqual(cgrCfg.statsCfg, eStatsCfg) {
		t.Errorf("received: %+v, expecting: %+v", cgrCfg.statsCfg, eStatsCfg)
	}
}

func TestCgrCfgJSONDefaultThresholdSCfg(t *testing.T) {
	eThresholdSCfg := &ThresholdSCfg{
		Enabled:             false,
		StoreInterval:       0,
		StringIndexedFields: nil,
		PrefixIndexedFields: &[]string{},
	}
	if !reflect.DeepEqual(eThresholdSCfg, cgrCfg.thresholdSCfg) {
		t.Errorf("received: %+v, expecting: %+v", eThresholdSCfg, cgrCfg.thresholdSCfg)
	}
}

func TestCgrCfgJSONDefaultSupplierSCfg(t *testing.T) {
	eSupplSCfg := &SupplierSCfg{
		Enabled:             false,
		StringIndexedFields: nil,
		PrefixIndexedFields: &[]string{},
		AttributeSConns:     []*HaPoolConfig{},
		RALsConns: []*HaPoolConfig{
			&HaPoolConfig{Address: "*internal"},
		},
		ResourceSConns: []*HaPoolConfig{},
		StatSConns:     []*HaPoolConfig{},
	}
	if !reflect.DeepEqual(eSupplSCfg, cgrCfg.supplierSCfg) {
		t.Errorf("received: %+v, expecting: %+v", eSupplSCfg, cgrCfg.supplierSCfg)
	}
}

func TestCgrCfgJSONDefaultsDiameterAgentCfg(t *testing.T) {
	testDA := &DiameterAgentCfg{
		Enabled:         false,
		Listen:          "127.0.0.1:3868",
		DictionariesDir: "/usr/share/cgrates/diameter/dict/",
		SessionSConns: []*HaPoolConfig{
			&HaPoolConfig{Address: "*internal"}},
		PubSubConns:       []*HaPoolConfig{},
		CreateCDR:         true,
		DebitInterval:     5 * time.Minute,
		Timezone:          "",
		OriginHost:        "CGR-DA",
		OriginRealm:       "cgrates.org",
		VendorId:          0,
		ProductName:       "CGRateS",
		RequestProcessors: nil,
	}

	if !reflect.DeepEqual(cgrCfg.diameterAgentCfg.Enabled, testDA.Enabled) {
		t.Errorf("expecting: %+v, received: %+v", cgrCfg.diameterAgentCfg.Enabled, testDA.Enabled)
	}
	if !reflect.DeepEqual(cgrCfg.diameterAgentCfg.Listen, testDA.Listen) {
		t.Errorf("expecting: %+v, received: %+v", cgrCfg.diameterAgentCfg.Listen, testDA.Listen)
	}
	if !reflect.DeepEqual(cgrCfg.diameterAgentCfg.DictionariesDir, testDA.DictionariesDir) {
		t.Errorf("expecting: %+v, received: %+v", cgrCfg.diameterAgentCfg.DictionariesDir, testDA.DictionariesDir)
	}
	if !reflect.DeepEqual(cgrCfg.diameterAgentCfg.SessionSConns, testDA.SessionSConns) {
		t.Errorf("expecting: %+v, received: %+v", cgrCfg.diameterAgentCfg.SessionSConns, testDA.SessionSConns)
	}
	if !reflect.DeepEqual(cgrCfg.diameterAgentCfg.PubSubConns, testDA.PubSubConns) {
		t.Errorf("expecting: %+v, received: %+v", cgrCfg.diameterAgentCfg.PubSubConns, testDA.PubSubConns)
	}
	if !reflect.DeepEqual(cgrCfg.diameterAgentCfg.CreateCDR, testDA.CreateCDR) {
		t.Errorf("expecting: %+v, received: %+v", cgrCfg.diameterAgentCfg.CreateCDR, testDA.CreateCDR)
	}
	if !reflect.DeepEqual(cgrCfg.diameterAgentCfg.DebitInterval, testDA.DebitInterval) {
		t.Errorf("expecting: %+v, received: %+v", cgrCfg.diameterAgentCfg.DebitInterval, testDA.DebitInterval)
	}
	if !reflect.DeepEqual(cgrCfg.diameterAgentCfg.Timezone, testDA.Timezone) {
		t.Errorf("received: %+v, expecting: %+v", cgrCfg.diameterAgentCfg.Timezone, testDA.Timezone)
	}
	if !reflect.DeepEqual(cgrCfg.diameterAgentCfg.OriginHost, testDA.OriginHost) {
		t.Errorf("received: %+v, expecting: %+v", cgrCfg.diameterAgentCfg.OriginHost, testDA.OriginHost)
	}
	if !reflect.DeepEqual(cgrCfg.diameterAgentCfg.OriginRealm, testDA.OriginRealm) {
		t.Errorf("received: %+v, expecting: %+v", cgrCfg.diameterAgentCfg.OriginRealm, testDA.OriginRealm)
	}
	if !reflect.DeepEqual(cgrCfg.diameterAgentCfg.VendorId, testDA.VendorId) {
		t.Errorf("received: %+v, expecting: %+v", cgrCfg.diameterAgentCfg.VendorId, testDA.VendorId)
	}
	if !reflect.DeepEqual(cgrCfg.diameterAgentCfg.ProductName, testDA.ProductName) {
		t.Errorf("received: %+v, expecting: %+v", cgrCfg.diameterAgentCfg.ProductName, testDA.ProductName)
	}
	if !reflect.DeepEqual(cgrCfg.diameterAgentCfg.RequestProcessors, testDA.RequestProcessors) {
		t.Errorf("expecting: %+v, received: %+v", testDA.RequestProcessors, cgrCfg.diameterAgentCfg.RequestProcessors)
	}
}

func TestCgrCfgJSONDefaultsMailer(t *testing.T) {
	if cgrCfg.MailerServer != "localhost" {
		t.Error(cgrCfg.MailerServer)
	}
	if cgrCfg.MailerAuthUser != "cgrates" {
		t.Error(cgrCfg.MailerAuthUser)
	}
	if cgrCfg.MailerAuthPass != "CGRateS.org" {
		t.Error(cgrCfg.MailerAuthPass)
	}
	if cgrCfg.MailerFromAddr != "cgr-mailer@localhost.localdomain" {
		t.Error(cgrCfg.MailerFromAddr)
	}
}

func TestCgrCfgJSONDefaultsSureTax(t *testing.T) {
	localt, err := time.LoadLocation("Local")
	if err != nil {
		t.Error("time parsing error", err)
	}
	eSureTaxCfg := &SureTaxCfg{
		Url:                  "",
		ClientNumber:         "",
		ValidationKey:        "",
		BusinessUnit:         "",
		Timezone:             localt,
		IncludeLocalCost:     false,
		ReturnFileCode:       "0",
		ResponseGroup:        "03",
		ResponseType:         "D4",
		RegulatoryCode:       "03",
		ClientTracking:       utils.ParseRSRFieldsMustCompile("CGRID", utils.INFIELD_SEP),
		CustomerNumber:       utils.ParseRSRFieldsMustCompile("Subject", utils.INFIELD_SEP),
		OrigNumber:           utils.ParseRSRFieldsMustCompile("Subject", utils.INFIELD_SEP),
		TermNumber:           utils.ParseRSRFieldsMustCompile("Destination", utils.INFIELD_SEP),
		BillToNumber:         utils.ParseRSRFieldsMustCompile("", utils.INFIELD_SEP),
		Zipcode:              utils.ParseRSRFieldsMustCompile("", utils.INFIELD_SEP),
		P2PZipcode:           utils.ParseRSRFieldsMustCompile("", utils.INFIELD_SEP),
		P2PPlus4:             utils.ParseRSRFieldsMustCompile("", utils.INFIELD_SEP),
		Units:                utils.ParseRSRFieldsMustCompile("^1", utils.INFIELD_SEP),
		UnitType:             utils.ParseRSRFieldsMustCompile("^00", utils.INFIELD_SEP),
		TaxIncluded:          utils.ParseRSRFieldsMustCompile("^0", utils.INFIELD_SEP),
		TaxSitusRule:         utils.ParseRSRFieldsMustCompile("^04", utils.INFIELD_SEP),
		TransTypeCode:        utils.ParseRSRFieldsMustCompile("^010101", utils.INFIELD_SEP),
		SalesTypeCode:        utils.ParseRSRFieldsMustCompile("^R", utils.INFIELD_SEP),
		TaxExemptionCodeList: utils.ParseRSRFieldsMustCompile("", utils.INFIELD_SEP),
	}

	if !reflect.DeepEqual(cgrCfg.sureTaxCfg, eSureTaxCfg) {
		t.Errorf("received: %+v, expecting: %+v", cgrCfg.sureTaxCfg, eSureTaxCfg)
	}
}

func TestCgrCfgJSONDefaultsHTTP(t *testing.T) {
	if cgrCfg.HTTPJsonRPCURL != "/jsonrpc" {
		t.Error(cgrCfg.HTTPJsonRPCURL)
	}
	if cgrCfg.HTTPWSURL != "/ws" {
		t.Error(cgrCfg.HTTPWSURL)
	}
	if cgrCfg.HTTPFreeswitchCDRsURL != "/freeswitch_json" {
		t.Error(cgrCfg.HTTPFreeswitchCDRsURL)
	}
	if cgrCfg.HTTPCDRsURL != "/cdr_http" {
		t.Error(cgrCfg.HTTPCDRsURL)
	}
	if cgrCfg.HTTPUseBasicAuth != false {
		t.Error(cgrCfg.HTTPUseBasicAuth)
	}
	if !reflect.DeepEqual(cgrCfg.HTTPAuthUsers, map[string]string{}) {
		t.Error(cgrCfg.HTTPAuthUsers)
	}
}

func TestRadiusAgentCfg(t *testing.T) {
	testRA := &RadiusAgentCfg{
		Enabled:            false,
		ListenNet:          "udp",
		ListenAuth:         "127.0.0.1:1812",
		ListenAcct:         "127.0.0.1:1813",
		ClientSecrets:      map[string]string{utils.META_DEFAULT: "CGRateS.org"},
		ClientDictionaries: map[string]string{utils.META_DEFAULT: "/usr/share/cgrates/radius/dict/"},
		SessionSConns:      []*HaPoolConfig{&HaPoolConfig{Address: utils.MetaInternal}},
		CDRRequiresSession: false,
		Timezone:           "",
		RequestProcessors:  nil,
	}
	if !reflect.DeepEqual(cgrCfg.radiusAgentCfg, testRA) {
		t.Errorf("expecting: %+v, received: %+v", cgrCfg.radiusAgentCfg, testRA)
	}
}

func TestDbDefaults(t *testing.T) {
	dbdf := NewDbDefaults()
	flagInput := utils.MetaDynamic
	dbs := []string{utils.MONGO, utils.REDIS, utils.MYSQL, utils.INTERNAL}
	for _, dbtype := range dbs {
		host := dbdf.DBHost(dbtype, flagInput)
		if host != utils.LOCALHOST {
			t.Errorf("received: %+v, expecting: %+v", host, utils.LOCALHOST)
		}
		user := dbdf.DBUser(dbtype, flagInput)
		if user != utils.CGRATES {
			t.Errorf("received: %+v, expecting: %+v", user, utils.CGRATES)
		}
		port := dbdf.DBPort(dbtype, flagInput)
		if port != dbdf[dbtype]["DbPort"] {
			t.Errorf("received: %+v, expecting: %+v", port, dbdf[dbtype]["DbPort"])
		}
		name := dbdf.DBName(dbtype, flagInput)
		if name != dbdf[dbtype]["DbName"] {
			t.Errorf("received: %+v, expecting: %+v", name, dbdf[dbtype]["DbName"])
		}
		pass := dbdf.DBPass(dbtype, flagInput)
		if pass != dbdf[dbtype]["DbPass"] {
			t.Errorf("received: %+v, expecting: %+v", pass, dbdf[dbtype]["DbPass"])
		}
	}
}

func TestCgrLoaderCfgITDefaults(t *testing.T) {
	eCfg := []*LoaderSConfig{
		&LoaderSConfig{
			Id:           utils.META_DEFAULT,
			Enabled:      false,
			DryRun:       false,
			RunDelay:     0,
			LockFileName: ".cgr.lck",
			CacheSConns: []*HaPoolConfig{
				&HaPoolConfig{
					Address: utils.MetaInternal,
				},
			},
			FieldSeparator: ",",
			TpInDir:        "/var/spool/cgrates/loader/in",
			TpOutDir:       "/var/spool/cgrates/loader/out",
			Data: []*LoaderDataType{
				&LoaderDataType{
					Type:     utils.MetaAttributes,
					Filename: utils.AttributesCsv,
					Fields: []*CfgCdrField{
						&CfgCdrField{Tag: "TenantID",
							FieldId:   "Tenant",
							Type:      utils.META_COMPOSED,
							Value:     utils.ParseRSRFieldsMustCompile("0", utils.INFIELD_SEP),
							Mandatory: true},
						&CfgCdrField{Tag: "ProfileID",
							FieldId:   "ID",
							Type:      utils.META_COMPOSED,
							Value:     utils.ParseRSRFieldsMustCompile("1", utils.INFIELD_SEP),
							Mandatory: true},
						&CfgCdrField{Tag: "Contexts",
							FieldId: "Contexts",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("2", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "FilterIDs",
							FieldId: "FilterIDs",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("3", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "ActivationInterval",
							FieldId: "ActivationInterval",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("4", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "FieldName",
							FieldId: "FieldName",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("5", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Initial",
							FieldId: "Initial",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("6", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Substitute",
							FieldId: "Substitute",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("7", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Append",
							FieldId: "Append",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("8", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Weight",
							FieldId: "Weight",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("9", utils.INFIELD_SEP)},
					},
				},
				&LoaderDataType{
					Type:     utils.MetaFilters,
					Filename: utils.FiltersCsv,
					Fields: []*CfgCdrField{
						&CfgCdrField{Tag: "Tenant",
							FieldId:   "Tenant",
							Type:      utils.META_COMPOSED,
							Value:     utils.ParseRSRFieldsMustCompile("0", utils.INFIELD_SEP),
							Mandatory: true},
						&CfgCdrField{Tag: "ID",
							FieldId:   "ID",
							Type:      utils.META_COMPOSED,
							Value:     utils.ParseRSRFieldsMustCompile("1", utils.INFIELD_SEP),
							Mandatory: true},
						&CfgCdrField{Tag: "FilterType",
							FieldId: "FilterType",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("2", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "FilterFieldName",
							FieldId: "FilterFieldName",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("3", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "FilterFieldValues",
							FieldId: "FilterFieldValues",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("4", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "ActivationInterval",
							FieldId: "ActivationInterval",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("5", utils.INFIELD_SEP)},
					},
				},
				&LoaderDataType{
					Type:     utils.MetaResources,
					Filename: utils.ResourcesCsv,
					Fields: []*CfgCdrField{
						&CfgCdrField{Tag: "Tenant",
							FieldId:   "Tenant",
							Type:      utils.META_COMPOSED,
							Value:     utils.ParseRSRFieldsMustCompile("0", utils.INFIELD_SEP),
							Mandatory: true},
						&CfgCdrField{Tag: "ID",
							FieldId:   "ID",
							Type:      utils.META_COMPOSED,
							Value:     utils.ParseRSRFieldsMustCompile("1", utils.INFIELD_SEP),
							Mandatory: true},
						&CfgCdrField{Tag: "FilterIDs",
							FieldId: "FilterIDs",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("2", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "ActivationInterval",
							FieldId: "ActivationInterval",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("3", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "TTL",
							FieldId: "UsageTTL",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("4", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Limit",
							FieldId: "Limit",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("5", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "AllocationMessage",
							FieldId: "AllocationMessage",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("6", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Blocker",
							FieldId: "Blocker",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("7", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Stored",
							FieldId: "Stored",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("8", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Weight",
							FieldId: "Weight",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("9", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "ThresholdIDs",
							FieldId: "ThresholdIDs",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("10", utils.INFIELD_SEP)},
					},
				},
				&LoaderDataType{
					Type:     utils.MetaStats,
					Filename: utils.StatsCsv,
					Fields: []*CfgCdrField{
						&CfgCdrField{Tag: "Tenant",
							FieldId:   "Tenant",
							Type:      utils.META_COMPOSED,
							Value:     utils.ParseRSRFieldsMustCompile("0", utils.INFIELD_SEP),
							Mandatory: true},
						&CfgCdrField{Tag: "ID",
							FieldId:   "ID",
							Type:      utils.META_COMPOSED,
							Value:     utils.ParseRSRFieldsMustCompile("1", utils.INFIELD_SEP),
							Mandatory: true},
						&CfgCdrField{Tag: "FilterIDs",
							FieldId: "FilterIDs",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("2", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "ActivationInterval",
							FieldId: "ActivationInterval",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("3", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "QueueLength",
							FieldId: "QueueLength",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("4", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "TTL",
							FieldId: "TTL",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("5", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Metrics",
							FieldId: "Metrics",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("6", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "MetricParams",
							FieldId: "Parameters",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("7", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Blocker",
							FieldId: "Blocker",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("8", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Stored",
							FieldId: "Stored",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("9", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Weight",
							FieldId: "Weight",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("10", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "MinItems",
							FieldId: "MinItems",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("11", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "ThresholdIDs",
							FieldId: "ThresholdIDs",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("12", utils.INFIELD_SEP)},
					},
				},
				&LoaderDataType{
					Type:     utils.MetaThresholds,
					Filename: utils.ThresholdsCsv,
					Fields: []*CfgCdrField{
						&CfgCdrField{Tag: "Tenant",
							FieldId:   "Tenant",
							Type:      utils.META_COMPOSED,
							Value:     utils.ParseRSRFieldsMustCompile("0", utils.INFIELD_SEP),
							Mandatory: true},
						&CfgCdrField{Tag: "ID",
							FieldId:   "ID",
							Type:      utils.META_COMPOSED,
							Value:     utils.ParseRSRFieldsMustCompile("1", utils.INFIELD_SEP),
							Mandatory: true},
						&CfgCdrField{Tag: "FilterIDs",
							FieldId: "FilterIDs",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("2", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "ActivationInterval",
							FieldId: "ActivationInterval",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("3", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "MaxHits",
							FieldId: "MaxHits",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("4", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "MinHits",
							FieldId: "MinHits",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("5", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "MinSleep",
							FieldId: "MinSleep",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("6", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Blocker",
							FieldId: "Blocker",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("7", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Weight",
							FieldId: "Weight",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("8", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "ActionIDs",
							FieldId: "ActionIDs",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("9", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Async",
							FieldId: "Async",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("10", utils.INFIELD_SEP)},
					},
				},
				&LoaderDataType{
					Type:     utils.MetaSuppliers,
					Filename: utils.SuppliersCsv,
					Fields: []*CfgCdrField{
						&CfgCdrField{Tag: "Tenant",
							FieldId:   "Tenant",
							Type:      utils.META_COMPOSED,
							Value:     utils.ParseRSRFieldsMustCompile("0", utils.INFIELD_SEP),
							Mandatory: true},
						&CfgCdrField{Tag: "ID",
							FieldId:   "ID",
							Type:      utils.META_COMPOSED,
							Value:     utils.ParseRSRFieldsMustCompile("1", utils.INFIELD_SEP),
							Mandatory: true},
						&CfgCdrField{Tag: "FilterIDs",
							FieldId: "FilterIDs",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("2", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "ActivationInterval",
							FieldId: "ActivationInterval",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("3", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Sorting",
							FieldId: "Sorting",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("4", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "SortingParamameters",
							FieldId: "SortingParamameters",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("5", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "SupplierID",
							FieldId: "SupplierID",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("6", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "SupplierFilterIDs",
							FieldId: "SupplierFilterIDs",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("7", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "SupplierAccountIDs",
							FieldId: "SupplierAccountIDs",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("8", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "SupplierRatingPlanIDs",
							FieldId: "SupplierRatingPlanIDs",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("9", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "SupplierResourceIDs",
							FieldId: "SupplierResourceIDs",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("10", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "SupplierStatIDs",
							FieldId: "SupplierStatIDs",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("11", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "SupplierWeight",
							FieldId: "SupplierWeight",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("12", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "SupplierBlocker",
							FieldId: "SupplierBlocker",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("13", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "SupplierParameters",
							FieldId: "SupplierParameters",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("14", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Weight",
							FieldId: "Weight",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("15", utils.INFIELD_SEP)},
					},
				},
				&LoaderDataType{
					Type:     utils.MetaChargers,
					Filename: utils.ChargersCsv,
					Fields: []*CfgCdrField{
						&CfgCdrField{Tag: "Tenant",
							FieldId:   "Tenant",
							Type:      utils.META_COMPOSED,
							Value:     utils.ParseRSRFieldsMustCompile("0", utils.INFIELD_SEP),
							Mandatory: true},
						&CfgCdrField{Tag: "ID",
							FieldId:   "ID",
							Type:      utils.META_COMPOSED,
							Value:     utils.ParseRSRFieldsMustCompile("1", utils.INFIELD_SEP),
							Mandatory: true},
						&CfgCdrField{Tag: "FilterIDs",
							FieldId: "FilterIDs",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("2", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "ActivationInterval",
							FieldId: "ActivationInterval",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("3", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "RunID",
							FieldId: "RunID",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("4", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "AttributeIDs",
							FieldId: "AttributeIDs",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("5", utils.INFIELD_SEP)},
						&CfgCdrField{Tag: "Weight",
							FieldId: "Weight",
							Type:    utils.META_COMPOSED,
							Value:   utils.ParseRSRFieldsMustCompile("6", utils.INFIELD_SEP)},
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(eCfg, cgrCfg.loaderCfg) {
		t.Errorf("received: %+v, expecting: %+v",
			utils.ToJSON(eCfg), utils.ToJSON(cgrCfg.loaderCfg))
	}
}

func TestCgrCfgJSONDefaultDispatcherSCfg(t *testing.T) {
	eDspSCfg := &DispatcherSCfg{
		Enabled:             false,
		RALsConns:           []*HaPoolConfig{},
		ResSConns:           []*HaPoolConfig{},
		ThreshSConns:        []*HaPoolConfig{},
		StatSConns:          []*HaPoolConfig{},
		SupplSConns:         []*HaPoolConfig{},
		AttrSConns:          []*HaPoolConfig{},
		SessionSConns:       []*HaPoolConfig{},
		ChargerSConns:       []*HaPoolConfig{},
		DispatchingStrategy: utils.MetaFirst,
	}
	if !reflect.DeepEqual(cgrCfg.dispatcherSCfg, eDspSCfg) {
		t.Errorf("received: %+v, expecting: %+v", cgrCfg.dispatcherSCfg, eDspSCfg)
	}
}

func TestCgrLoaderCfgDefault(t *testing.T) {
	eLdrCfg := &LoaderCgrCfg{
		TpID:           "",
		DataPath:       "",
		DisableReverse: false,
		CachesConns: []*HaPoolConfig{
			&HaPoolConfig{
				Address:   "127.0.0.1:2012",
				Transport: utils.MetaJSONrpc,
			},
		},
		SchedulerConns: []*HaPoolConfig{
			&HaPoolConfig{
				Address: "127.0.0.1:2012",
			},
		},
	}
	if !reflect.DeepEqual(cgrCfg.LoaderCgrConfig, eLdrCfg) {
		t.Errorf("received: %+v, expecting: %+v", utils.ToJSON(cgrCfg.LoaderCgrConfig), utils.ToJSON(eLdrCfg))
	}
}

func TestCgrMigratorCfgDefault(t *testing.T) {
	eMgrCfg := &MigratorCgrCfg{
		OutDataDBType:     "redis",
		OutDataDBHost:     "127.0.0.1",
		OutDataDBPort:     "6379",
		OutDataDBName:     "10",
		OutDataDBUser:     "cgrates",
		OutDataDBPassword: "",
		OutDataDBEncoding: "msgpack",
		OutStorDBType:     "mysql",
		OutStorDBHost:     "127.0.0.1",
		OutStorDBPort:     "3306",
		OutStorDBName:     "cgrates",
		OutStorDBUser:     "cgrates",
		OutStorDBPassword: "",
	}
	if !reflect.DeepEqual(cgrCfg.MigratorCgrConfig, eMgrCfg) {
		t.Errorf("received: %+v, expecting: %+v", utils.ToJSON(cgrCfg.MigratorCgrConfig), utils.ToJSON(eMgrCfg))
	}
}

func TestCgrMigratorCfg2(t *testing.T) {
	JSN_CFG := `
{
"migrator": {
	"out_datadb_type": "redis",
	"out_datadb_host": "0.0.0.0",
	"out_datadb_port": "9999",
	"out_datadb_name": "9999",
	"out_datadb_user": "cgrates",
	"out_datadb_password": "",
	"out_datadb_encoding" : "msgpack",
	"out_stordb_type": "mysql",
	"out_stordb_host": "0.0.0.0",
	"out_stordb_port": "9999",
	"out_stordb_name": "cgrates",
	"out_stordb_user": "cgrates",
	"out_stordb_password": "",
},
}`

	if cgrCfg, err := NewCGRConfigFromJsonString(JSN_CFG); err != nil {
		t.Error(err)
	} else if cgrCfg.MigratorCgrConfig.OutDataDBHost != "0.0.0.0" {
		t.Errorf("Expected: 0.0.0.0 , received: %+v", cgrCfg.MigratorCgrConfig.OutDataDBHost)
	} else if cgrCfg.MigratorCgrConfig.OutDataDBPort != "9999" {
		t.Errorf("Expected: 9999, received: %+v", cgrCfg.MigratorCgrConfig.OutDataDBPassword)
	}
}
