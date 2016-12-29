// websocket
var wsUri = "ws://" + location.hostname + ":" + location.port + "/ws";
var ws = null;

function NewWebScoket(){
  if(ws){
    return false
  }
  if ('WebSocket' in window) {
      ws = new WebSocket(wsUri);
  } else if ('MozWebSocket' in window) {
      ws = new MozWebSocket(wsUri);
  }

  // 发送api
  ws.onsend = function(data) {
      var dataStr = JSON.stringify(data);
      ws.send(dataStr);
      printDbg("send: " + dataStr);
  }
};

var $pingBtn = $("#ping"),
    $logPre = $("#log");

var writeLog = function(message) {
    message = message + "<br>"
    $logPre.append(message)
    $logPre.scrollTop = $logPre.scrollHeight;
};

$pingBtn.click(function(evt){
  NewWebScoket();
  ws.onopen  = function(evt){
    printDbg("WebSocket connected.")
    ip = document.forms["ping-Addr"].ipAddr.value
    printDbg("开始ping..."+ip);
    $pingBtn.text("Stop")
      .attr("data-type", "stop")
      .addClass("btn-danger")
      .removeClass("btn-primary");
    $logPre.empty()

    ws.onsend({
      'ip': ip
    });
    setTimeout(function () {
      ws.onsend({
        'stop':true
      });
      printDbg("发送停止...")
    }, 2000);
  };

  ws.onclose = function(evt){
    printDbg("WebSocket connection closed.");
    $pingBtn.text("Ping")
      .attr("data-type", "run")
      .addClass("btn btn-primary")
      .removeClass("btn-danger");
    ws = null;
  };

ws.onmessage = function(evt) {
    printDbg("接收到消息...\n" + evt.data)
    writeLog(evt.data)
    $(window).scrollTop($(window).scrollTop() + 1);
  };

  ws.onerror = function(evt) {
      printDbg("WebSocket error: " + evt.data)
  };

  return false;
});

// 浏览器关闭
window.onbeforeunload = function() {
    ws.close();
    printDbg("WebSocket connection closed.")
    return
}
