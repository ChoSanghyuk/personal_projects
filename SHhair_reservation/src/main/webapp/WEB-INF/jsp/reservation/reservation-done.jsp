<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8" %>
<!DOCTYPE html>
<html>
<head>
	<title>done</title>
	<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-KyZXEAg3QhqLMpG8r+8fhAXLRk2vvoC2f3B09zVXn8CA5QIVfZOJ3BCsw2P0p/We" crossorigin="anonymous">
</head>
<body>
<jsp:include page="../include/navbar.jsp" flush="false"/>
<div class="container mt-5" style="height:300px">
	<div class="row">
		<div class="col-2 fs-4">
			<span>예약</span>
			<span style="color:red;">신청완료</span> 
		</div>
		<hr>
	</div>

	<div class="row">
		<div class="col-12 fs-1 text-center">
			예약이 완료되었습니다		
		</div>
	</div>
	<div class = row>
		<div class="offset-2 col-8">
			<table class="table table-bordered table-striped fs-5">
				<tr>
					<td>예약자 성함</td>
					<td>${customerName}</td>
				</tr>
				<tr>
					<td>예약 날짜</td>
					<td>${theReservation.getDate()}</td>
				</tr>
				<tr>
					<td>예약 시간</td>
					<td>${theReservation.getStartTime()} ~ ${theReservation.getEndTime()}</td>
				</tr>
				<tr>
					<td>예약 종류</td>
					<td>${theReservation.getType()}</td>
				</tr>
			
			</table>
		</div>
	</div>
</div>
</body>