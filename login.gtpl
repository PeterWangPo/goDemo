<html>
	<head>
	</head>
	<body>
		<form action="/login" method="post">
			username:<input type="text" name="username" /><br />
			password:<input type="password" name="password" /><br />
			<input type="hidden" name="token" value="{{.}}" />
			<input type="submit" value="submit" />
		</form>
	</body>
</html>