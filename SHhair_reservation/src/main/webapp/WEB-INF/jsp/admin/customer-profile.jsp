<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8" %>
<%@taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<%@ page import="com.shHair.reservation.entity.Customer"%>

<!DOCTYPE html>
<html>

<head>
	<title>SH hair Administration</title>
<link href="/<c:url value="resources/css/admin-reservation-list.css" />" rel="stylesheet">
<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-KyZXEAg3QhqLMpG8r+8fhAXLRk2vvoC2f3B09zVXn8CA5QIVfZOJ3BCsw2P0p/We" crossorigin="anonymous">
</head>
<body >
	<jsp:include page="../include/admin-bar.jsp" flush="false"/>
	<% Customer theCustomer = (Customer) request.getAttribute("theCustomer"); %>
	<div class="contaitner mt-5 container-size">
		<div class = "row text-center fs-1 bg-success bg-opacity-10">
			<div class = "col-12">${theCustomer.getName()} 회원정보</div>
		</div>
		<div class = "row">
			<div class = "col-6">
				<div class = "row fs-5 my-1">
					<div class = "col-3">회원 번호</div>
					<div class = "col-9">${theCustomer.getId()}</div>
				</div>
				<div class = "row fs-5 my-1">
					<div class = "col-3">전화 번호</div>
					<div class = "col-9">${theCustomer.getPhoneNumber()}</div>
				</div>
				<div class = "row fs-5 my-1">
					<div class = "col-3">성별</div>
					<% if(theCustomer.getSex() == 1){ %>
					<div class = "col-9">남자</div>
					<%}else { %>
					<div class = "col-9">여자</div>
					<%} %>
				</div>
				<div class = "row fs-5 my-1">
					<div class = "col-3">커트시간</div>
					<div class = "col-9">+${theCustomer.getCutTime()}</div>
				</div>
				<div class = "row fs-5 my-1">
					<div class = "col-3">파마시간</div>
					<div class = "col-9">+${theCustomer.getPermTime()}</div>
				</div>
				<div class = "row fs-5 my-1">
					<div class = "col-3">염색시간</div>
					<div class = "col-9">+${theCustomer.getDyeTime()}</div>
				</div>
			</div>
			
			<div class = "col-6">
				<div class = "row fs-5 my-1">
					<div class = "col-3" style="margin : auto">특이사항</div>
					<textarea class = "col-9" rows="8" cols=""></textarea>
				</div>
			</div>
		</div>
		<div class = "row fs-5 my-1">
			<div class = "col-3 align-middle" style="margin : auto">예약내역</div>
			<div class = "col-9 overflow-auto" style="height : 200px;">
				<c:forEach var = "theReservation" items="${theCustomer.getReservations()}" >
					<p><c:out value="${theReservation}" /> </p>
				</c:forEach>
			</div>
		</div>
	<a href="/admin/customer/update/${theCustomer.getId()}" style="float:left;"><button class="btn btn-primary mx-1">수정하기</button></a>
	<form action = "/admin/customer/delete/${theCustomer.getId()}" method="POST"> <button class="btn btn-primary" style="float:left;">회원삭제</button></form>
	</div>


<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-U1DAWAznBHeqEIlVSCgzq+c9gqGAJn5c/t99JyeKa9xxaYpSvHU5awsuZVVFIhvj" crossorigin="anonymous"></script>
</body>

</html>