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
		<form action="/admin/customer/update/${theCustomer.getId()}" method="POST" class = "row" id="customer-update">
			<div class = "col-6">
				<div class = "row fs-5 my-1">
					<div class = "col-3">회원 번호</div>
					<div class = "col-9">${theCustomer.getId()}</div>
				</div>
				<div class = "row fs-5 my-1">
					<label class="col-3" for="phoneNumber">전화 번호</label>
					<div class="col-9">					
		    			<input class="form-control" type="text" name="phoneNumber" id="phoneNumber" value="${theCustomer.getPhoneNumber()}" required>
					</div>
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
					<label class = "col-3" for="cutTime">커트시간</label>
					<div class="col-9">					
						<input class = "form-control" type="number" id="cutTime" name="cutTime" min="-30" max="180" step="10" value = "${theCustomer.getCutTime()}" >					
					</div>
				</div>
				<div class = "row fs-5 my-1">
					<label class = "col-3" for="permTime">파마시간</label>
					<div class="col-9">					
						<input class = "form-control" type="number" id="permTime" name="permTime" min="-60" max="180" step="10" value = "${theCustomer.getPermTime()}" >					
					</div>
				</div>
				<div class = "row fs-5 my-1">
					<label class = "col-3" for="dyeTime">염색시간</label>
					<div class="col-9">					
						<input class = "form-control" type="number" id="dyeTime" name="dyeTime" min="-60" max="180" step="10" value = "${theCustomer.getDyeTime()}" >					
					</div>
				</div>
			</div>
			
			<div class = "col-6">
				<div class = "row fs-5 my-1">
					<label class = "col-3" for="detail" style="margin : auto">특이사항</label>
					<div class="col-9">					
						<textarea class="form-control" rows="10" id="detail" name="detail"></textarea>
					</div>
				</div>
			</div>
		</form>
		<div class = "row fs-5 my-1">
			<div class = "col-3 align-middle" style="margin : auto">예약내역</div>
			<div class = "col-9 overflow-auto" style="height : 200px;">
				<c:forEach var = "theReservation" items="${theCustomer.getReservations()}">
					<p><c:out value="${theReservation}" /> </p>
				</c:forEach>
			</div>
		</div>
	<input type="submit" class="btn btn-primary" form="customer-update" value="저장">
	<a href="/admin/customer/${theCustomer.getId()}"><button class="btn btn-primary">취소</button></a>
	</div>


<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-U1DAWAznBHeqEIlVSCgzq+c9gqGAJn5c/t99JyeKa9xxaYpSvHU5awsuZVVFIhvj" crossorigin="anonymous"></script>
</body>

</html>