<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8" %>
<%@taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<%@ page import="java.time.LocalDate"%>


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
			<a class="col-12 btn btn-outline-secondary btn-lg">예약 생성</a>
		</div>
		<div class="row my-1">
			<div class="offset-8 col-4 text-end">			
			<form action="/admin/reservation" method="GET">
				<% String today = request.getParameter("date") == null ? LocalDate.now().toString() : request.getParameter("date"); %>
        		<input name="date" id="date" placeholder="yyyy-mm-dd" type = "date" size="15" value="<%=today%>">
        		<input type = "submit" value="날짜 검색" >
			</form>
			</div>
		</div>
		<div class = "row text-center">
			<div class="col-2"> 고객 이름</div>
			<div class="col-2"> 예약 타입</div>
			<div class="col-2"> 시작 시간</div>
			<div class="col-2"> 예정 종료 시간</div>
			<div class="col-2"> 예약 수정</div>
			<div class="col-2"> 예약 삭제</div>
		</div>
		<hr>
		<c:forEach var = "theReservation" items="${reservations}">
			<div class="row text-center"> 
				<div class = "col-2">			
					<c:out value="${customerNames.get(theReservation.getCustomerId())}" /> 
				</div>
				<div class = "col-2">			
					<c:out value="${theReservation.getType()}" /> 
				</div>
				<div class = "col-2">			
					<c:out value="${theReservation.getStartTime()}" /> 
				</div>
				<div class = "col-2">			
					<c:out value="${theReservation.getEndTime()}" /> 
				</div>
				<div class = "col-2">
					<form action = "/admin/reservation/update/${theReservation.getId()}" method="GET"> <button>수정</button></form>
				</div>
				<div class = "col-2">
					<form action = "/admin/reservation/delete/${theReservation.getId()}" method="POST"> 
						<button onclick="return confirm('삭제하시겠습니까?');">삭제</button>
					</form>
				</div>
			</div>
			<hr>
		</c:forEach>
	
	</div>


<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-U1DAWAznBHeqEIlVSCgzq+c9gqGAJn5c/t99JyeKa9xxaYpSvHU5awsuZVVFIhvj" crossorigin="anonymous"></script>
</body>

</html>