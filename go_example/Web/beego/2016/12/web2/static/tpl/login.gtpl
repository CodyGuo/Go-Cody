<html>
<head>
	<title>test2</title>
</head>

<body>
<form action="http://127.0.0.1:8080/login" method="post">
	<input type="checkbox" name="interest" value="football">足球
	<input type="checkbox" name="interest" value="basketball">篮球
	<input type="checkbox" name="interest" value="tennis">网球
	用户名: <input type="text" name="username">
	密码: <input type="password" name="password">
	<input type="hidden" name="token" value="{{.}}">
	<input type="submit" value="登录">
</form>
</body>
</html>