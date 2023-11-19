<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8" %>
<%@taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<%//@ page import="java.time.LocalDate"%> 


<!DOCTYPE html>
<html>

<head>
	<title>SH hair Administration</title>
<link href="/<c:url value="resources/css/admin-reservation-list.css" />" rel="stylesheet">
<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-KyZXEAg3QhqLMpG8r+8fhAXLRk2vvoC2f3B09zVXn8CA5QIVfZOJ3BCsw2P0p/We" crossorigin="anonymous">
</head>
<body >
	<jsp:include page="../include/admin-bar.jsp" flush="false"/>
	<div class="contaitner mt-5 container-size">
		<div class="row">		
			<a class="col-12 btn btn-outline-secondary btn-lg">고객 생성</a>
		</div>
		<div class="row my-1 text-end">
			<div class="offset-9 col-3">			
			<form action="/admin/customer" method="GET">
        		<input name="name" id="userName"  type = "text" size="15">
        		<input type = "submit" value="이름 검색">
			</form>
			</div>
		</div>
		<div class = "row text-center">
			<div class="col-2"> 고객 번호</div>
			<div class="col-2"> 고객 이름</div>
			<div class="col-2"> 성별 </div>
			<div class="col-2"> 나이 </div>
			<div class="col-2"> 회원 정보 </div>
			<div class="col-2"> 삭제 </div>
		</div>
		<hr>
		<c:forEach var = "theCustomer" items="${customers}">
			<div class="row text-center"> 
				<div class = "col-2">			
					<c:out value="${theCustomer.getId()}" /> 
				</div>
				<div class = "col-2">			
					<c:out value="${theCustomer.getName()}" /> 
				</div>
				<div class = "col-2">			
					<c:out value="${theCustomer.getSex()}" /> 
				</div>
				<div class = "col-2">			
					
				</div>
				<a class="col-2" href="/admin/customer/${theCustomer.getId()}"><button class="btn btn-primary">회원 정보</button> </a>
				<div class = "col-2">
					<form action = "/admin/customer/delete/${theCustomer.getId()}" method="POST"> 
						<button class="btn btn-primary" onclick="return confirm('삭제하시겠습니까?');">회원 삭제</button>
					</form>
				</div>
			</div>
			<hr>
		</c:forEach>
	</div>


<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-U1DAWAznBHeqEIlVSCgzq+c9gqGAJn5c/t99JyeKa9xxaYpSvHU5awsuZVVFIhvj" crossorigin="anonymous"></script>
</body>

</html>