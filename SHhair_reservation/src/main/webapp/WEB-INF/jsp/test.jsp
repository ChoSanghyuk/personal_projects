<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8" %>
<%@ page import="java.util.Calendar"%>
<%@ page import="java.util.ArrayList"%>

<%@taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>


<!DOCTYPE html>
<html>

<head>
	<title>SH hair test</title>
</head>
<body >
	
<h1>Hello World</h1>
<p >${msg}</p>
<form method="POST" action="/test/b">
	<input type="text" name="id1" th:value="${id2}"/>
	<input type="submit" value="click" />
</form>
</body>
</html>