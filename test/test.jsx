"use strict";

var method = " -X GET";
var api = ' "http://localhost:5741/system/trash"';
var header = ' -H "Content-Type: application/json"';
var param = ' -d "{\\"entry\\":\\"D:\\\\Working\\\\tmp\\"}"';
var cmd = "curl" + method + api + header + param;
var res = system.callSystem(cmd + " -s");
alert(res);
