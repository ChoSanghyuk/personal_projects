<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8" %>
<%@ page import="java.util.ArrayList"%>
<!DOCTYPE html>
<html>
<head>
	<title>done</title>
	<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-KyZXEAg3QhqLMpG8r+8fhAXLRk2vvoC2f3B09zVXn8CA5QIVfZOJ3BCsw2P0p/We" crossorigin="anonymous">
</head>
<body>
<jsp:include page="../include/navbar.jsp" flush="false"/>
<div class="d-flex flex-column justify-content-center align-items-center" style="height:300px">
	<h1>예약 정보</h1>
	<form action="/reservation/update/${theReservation.getId()}" >
		<div>
			<label for="date">예약 날짜:</label>
			<input id="date" name="date" type="date" value="${theReservation.getDate()}">
		</div>
		<div>
			<label for="type">예약 종류:</label>
			<input id="type" name="type" type="text" value="${theReservation.getType()}">
		</div>
		<button data-bs-toggle="modal" data-bs-target="#staticBackdrop">예약하기</button>
	</form>
</div>
<div class="modal fade" id="staticBackdrop" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="staticBackdropLabel" aria-hidden="true">
	  <div class="modal-dialog">
	    <div class="modal-content">
	      <div class="modal-header">
	        <h5 class="modal-title" id="staticBackdropLabel">시간 선택</h5>
	        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
	      </div>
	      <div class="modal-body">
	      <form action="reservation/reserve" method="post" id="send_reservation_info_form">
	      	<%
	      		ArrayList<String> unavailableTimes = (ArrayList<String>) request.getAttribute("available");
	      	
		      	int hour = 10;
				int min = 0;
				for(int i = 0 ; i < 18 ; i++) {
					if(i%2==0) {
						hour += 1;
						min = 0;
					} else {
						min += 3;
					}
					String time = ""+hour+":"+min+"0";
					if( unavailableTimes.contains(time)) {
			%>
				<input type="radio" class="btn-check" id="<%=time %>" value="<%=time %>" name="time" autocomplete="off" disabled>
				<label class="btn btn-outline-success" for="<%=time %>"><%=time %></label>
			<% 
					} else {
	      	%>
	      		<input type="radio" class="btn-check" id="<%=time%>" value="<%=time %>" name="time" autocomplete="off">
				<label class="btn btn-outline-success" for="<%=time%>"><%=time%></label>
	        <% 		}
				}	%>
				<input id="id" name="id" type="hidden" value="${theReservation.getId()}">
				<input id="id" name="id" type="hidden" value="${theReservation.getCustomerId()}">
				<input type="hidden" name="date" value=<%=request.getParameter("date") %> >
				<input type="hidden" name="type" value=<%=request.getParameter("type")%> >
				<input type="hidden" name="customer_id" value=1 >
				
	      </form>
	      </div>
	      <div class="modal-footer">
	      
	        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
	        <button form="send_reservation_info_form" type="submit" class="btn btn-primary">예약하기</button>
	      </div>
	    </div>
	  </div>
	</div>	
</body>