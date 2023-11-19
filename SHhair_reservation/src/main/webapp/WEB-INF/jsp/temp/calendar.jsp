<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8" %>
<%@ page import="java.util.Calendar"%>
<%@ page import="java.util.ArrayList"%>

<%@taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>

<%
Calendar cal = Calendar.getInstance();
int y = cal.get(Calendar.YEAR);
int m = cal.get(Calendar.MONTH);

cal.set(y,m,1);
int dayOfWeek = cal.get(Calendar.DAY_OF_WEEK); // 일:0 ~ 토:7
int lastDay = cal.getActualMaximum(Calendar.DATE);

%>

<!DOCTYPE html>
<html>

<head>
	<title>SH hair Reservation</title>

<link href="<c:url value="resources/css/calendar.css" />" rel="stylesheet">
<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-KyZXEAg3QhqLMpG8r+8fhAXLRk2vvoC2f3B09zVXn8CA5QIVfZOJ3BCsw2P0p/We" crossorigin="anonymous">
</head>
<body >
	
	<div class="container mt-5">
	<div class="row">
		<div class = "col-12 col-lg-6">
		<form action="/reservation" method="GET" id="date_select_form"></form>
		<table class = "table table-striped text-center">
			<caption class="caption-top caption-font"><%=y %>년 <%=m+1 %>월</caption>
			<tr>
				<th>일</th>
				<th>월</th>
				<th>화</th>
				<th>수</th>
				<th>목</th>
				<th>금</th>
				<th>토</th>
			</tr>
			<tr>
	
			<%
			int count = 0;
			
			for(int s=1 ; s < dayOfWeek ; s++){
				out.print("<td></td>");
				count++;
			}
			
			for(int d = 1 ; d<=lastDay ; d++) {
				count++;
				String s = "" ;
				s += y;
				if(m+1 < 10) s += 0;
				s += m+1;
				if(d < 10) s += 0;
				s += d;
			%>
				<td>
					<button form="date_select_form" class="button-shape" name="reservation-date" value="<%=s %>"><%=d%></button>
					
				</td>
				
			<%
				
				if( count%7 == 0){
					out.print("</tr><tr>");
				}
			}
			%>
			</tr>
		</table>
			<div class="row">
				<button class="col-12" data-bs-toggle="modal" data-bs-target="#staticBackdrop">예약하기</button>
			</div>
		</div>
		<div class = "col-12 col-lg-6">
		<table class = "table table-striped text-center">
			<caption class="caption-top caption-font"><%=y %>년 <%=m+2 %>월</caption>
			<tr>
				<th>일</th>
				<th>월</th>
				<th>화</th>
				<th>수</th>
				<th>목</th>
				<th>금</th>
				<th>토</th>
			</tr>
			<tr>
	
			<%
			cal.set(y,m+1,1);
			dayOfWeek = cal.get(Calendar.DAY_OF_WEEK); // 일:0 ~ 토:7
			lastDay = cal.getActualMaximum(Calendar.DATE);
			count = 0;
			
			for(int s=1 ; s < dayOfWeek ; s++){
				out.print("<td></td>");
				count++;
			}
			
			for(int d = 1 ; d<=lastDay ; d++) {
				count++;
			%>
				<td><%=d %></td>
			<%
				if( count%7 == 0){
					out.print("</tr><tr>");
				}
			}
			%>
			</tr>
		
		</table>
		<button data-bs-toggle="modal" data-bs-target="#staticBackdrop">예약하기</button>
		</div>
		
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
			      	ArrayList<String> unavailableTimes = (ArrayList<String>) request.getAttribute("unavailable");
			      
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
					<input type="hidden" name="date" value=<%=request.getParameter("reservation-date") %> >
					<input type="hidden" name="type" value="cut" >
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
	</div>
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-U1DAWAznBHeqEIlVSCgzq+c9gqGAJn5c/t99JyeKa9xxaYpSvHU5awsuZVVFIhvj" crossorigin="anonymous"></script>
</body>

</html>