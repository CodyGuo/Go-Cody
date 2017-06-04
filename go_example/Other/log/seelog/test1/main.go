package main

import (
	log "github.com/cihub/seelog"
)

var testConfig string

func main() {
	initConfig()
	defer log.Flush()
	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)
	log.Info("hello")
}

func initConfig() {
	testConfig = `
<seelog minlevel="trace" maxlevel="critical">
    <outputs>
        <filter levels="trace">
            <console formatid="colored_trace"/>
        </filter>
        <filter levels="debug">
            <console formatid="colored_debug"/>
        </filter>
        <filter levels="info">
            <console formatid="colored_info"/>
        </filter>
        <filter levels="warn">
            <console formatid="colored_warn"/>
        </filter>
        <filter levels="error">
            <console formatid="colored_error"/>
        </filter>
        <filter levels="critical">
            <console formatid="colored_critical"/>
        </filter>
    </outputs>
    <formats>
        <format id="colored_trace" format="%EscM(0)[%Date(2/Jan/2006 15:04:05)] %EscM(36;1)[%l]%EscM(0)%EscM(36) %Msg%n%EscM(0)"/>
        <format id="colored_debug" format="%EscM(0)[%Date(2/Jan/2006 15:04:05)] %EscM(34)[%l]%EscM(0)%EscM(34;1) %Msg%n%EscM(0)"/>
        <format id="colored_info" format="%EscM(0)[%Date(2/Jan/2006 15:04:05)] %EscM(32;1)[%l]%EscM(0)%EscM(32) %Msg%n%EscM(0)"/>
        <format id="colored_warn" format="%EscM(0)[%Date(2/Jan/2006 15:04:05)] %EscM(33;1)[%l]%EscM(0)%EscM(33) %Msg%n%EscM(0)"/>
        <format id="colored_error" format="%EscM(0)[%Date(2/Jan/2006 15:04:05)] %EscM(35;1)[%l]%EscM(0)%EscM(35) %Msg%n%EscM(0)"/>
        <format id="colored_critical" format="%EscM(0)[%Date(2/Jan/2006 15:04:05)] %EscM(31;1)[%l]%EscM(0)%EscM(31) %Msg%n%EscM(0)"/>
    </formats>
</seelog>
`
}
