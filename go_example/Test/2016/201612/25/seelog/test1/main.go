package main

import (
	"fmt"

	log "github.com/cihub/seelog"
)

func main() {
	defer log.Flush()
	testConfig := `
  <seelog type="sync">
  	<outputs formatid="main">
  		<console/>
      <file path="./log/log.log"/>
      <rollingfile type="size" filename="./log/roll.log" maxsize="100" maxrolls="5" />
      <smtp senderaddress="codyguo@aliyun.com" sendername="Automatic notification service" hostname="mail.aliyun.com" hostport="25" username="codyguo" password="aptech1!">
        <recipient address="469708924@qq.com"/>
      </smtp>
  	</outputs>
  	<formats>
  		<format id="main" format="[%Lev] [%Date-%Time (%Func %Line)] %Msg%n"/>
	</formats>
  </seelog>`

	logger, err := log.LoggerFromConfigAsBytes([]byte(testConfig))
	if err != nil {
		fmt.Println(err)
	}

	loggerErr := log.ReplaceLogger(logger)
	if loggerErr != nil {
		fmt.Println(loggerErr)
	}

	log.Info("hello from seelog!")
	printLog()
}

func printLog() {
	log.Debug("debug log.")
}
