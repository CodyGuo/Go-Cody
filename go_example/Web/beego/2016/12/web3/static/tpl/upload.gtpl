<html>
<head>
	<title>上传文件</title>
</head>
<body>
	<form action="http://127.0.0.1:8080/upload" enctype="multipart/form-data" method="post">
		<input type="file" name="uploadfile" />
		<input type="hidden" name="token" value="{{.}}" />
		<input type="submit" value="upload">
	</form>
</body>
</html>