var DEBUG = "[Debug]";
var $mode = $("#js_mode");

function selectMode(m) {
    switch (m) {
        case lping:
            $mode.text("ping ").data("type", lping);
            printLogDbg("select -----> ping");
            break;
        case ltcpdump:
            $mode.text("抓包").data("type", ltcpdump);
            printLogDbg("select -----> tcpdump");
            break;
        case ltraceroute:
            $mode.text("traceroute").data("type", ltraceroute);
            printLogDbg("select -----> traceroute");
            break;
        case lnslookup:
            $mode.text("nslookup").data("type", lnslookup);
            printLogDbg("select -----> nslookup");
            break;
        default:
            printLogDbg("select -----> error");
    }
}
