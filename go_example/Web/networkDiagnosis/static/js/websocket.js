// websocket
var wsUri = "ws://" + location.hostname + ":" + location.port + "/ws";
var ws = null;

function NewWebScoket() {
    if (ws) {
        return false;
    }
    if ('WebSocket' in window) {
        ws = new WebSocket(wsUri);
    } else if ('MozWebSocket' in window) {
        ws = new MozWebSocket(wsUri);
    }
}

function writeLog(message) {
    message = message + "<br>";
    $logPre.append(message);
    $logPre.scrollTop($logPre[0].scrollHeight);
}

var $pingBtn = $("#pingBtn"),
    $logPre = $("#log"),
    $menutools = $("#menutools");

$pingBtn.click(function(evt) {
    if ($pingBtn.data("type") == 'stop') {
        ws.onsend({
            'stop': true
        });

        writeLog("\n" + $mode.text() + " 已停止...");
        return false;
    }
    NewWebScoket();
    EventwebSocket();
    return true;
});

// 事件处理
function EventwebSocket() {
    // 连接事件发送诊断类型和参数地址
    ws.onopen = function(evt) {
        groupBeg($mode.text());
        printLogDbg("WebSocket connected.");

        type = $mode.data("type");
        args = $("#ipAddr")[0].value;

        printLogDbg("开始 " + $mode.text() + " " + args);
        $pingBtn.text("停止").data("type", "stop")
            .addClass("btn-danger").removeClass("btn-primary");
        $menutools.addClass("disabled");
        $logPre.empty();

        ws.onsend({
            'type': type,
            'args': args,
        });
    };

    // 发送api
    ws.onsend = function(data) {
        var dataStr = JSON.stringify(data);
        ws.send(dataStr);
        printLogDbg("send -----> " + dataStr);
    };

    // 接收api
    ws.onmessage = function(evt) {
        printLogDbg("received -----> " + evt.data);
        writeLog(evt.data);
    };

    // 关闭api
    ws.onclose = function(evt) {
        printLogDbg("WebSocket connection closed.");
        $pingBtn.text("执行").data("type", "run")
            .addClass("btn btn-primary").removeClass("btn-danger");
        $menutools.removeClass("disabled");
        ws = null;
        groupEnd($mode.text());
    };

    ws.onerror = function(evt) {
        printLogDbg("WebSocket error: " + evt.data);
    };
}


// 浏览器关闭
window.onbeforeunload = function() {
    ws.close();
    printLogDbg("WebSocket connection closed.");
    return;
};
