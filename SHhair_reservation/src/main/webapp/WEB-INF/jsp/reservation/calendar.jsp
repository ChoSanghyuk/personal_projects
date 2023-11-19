<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8" %>
<%@ page import="java.util.Calendar"%>
<%@ page import="java.util.ArrayList"%>
<%@ page import="com.shHair.reservation.entity.Time"%>

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
	<jsp:include page="../include/navbar.jsp" flush="false"/>
	<div class="container mt-5">
	<div class="row">
		<div class = "col-6 text-center fs-2">
				<%=y%>년 <%=m+1%>월
		</div>
		<% if(request.getParameter("date") != null) {%>
		<div class = "col-6 fs-5 " style="margin-top : auto">
				날짜 : <%=request.getParameter("date") %>	 &nbsp; 종류: <%=request.getParameter("type")%>
		</div>
		<%} %>
	</div>
	<div class = "row">
		<div class = "col-6">
		<form action="" method="GET" id="date_type_select"></form>
		<table class = "table table-striped text-center">
			
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
				s += "-";
				if(m+1 < 10) s += 0;
				s += m+1;
				s += "-";
				if(d < 10) s += 0;
				s += d;
				if(s.equals(request.getParameter("date"))){
			%>
				<td>
					<input type="radio" class="btn-check" form="date_type_select" name="date" id="<%=s%>>" value="<%=s%>" autocomplete="off" checked>
					<label class="btn btn-outline-success" for="<%=s%>>"><%=d%></label>
				</td>
			<%					
				}else {
			%>
				<td>
					<input type="radio" class="btn-check" form="date_type_select" name="date" id="<%=s%>>" value="<%=s%>" autocomplete="off">
					<label class="btn btn-outline-success" for="<%=s%>>"><%=d%></label>
				</td>
				
			<%
				}
				if( count%7 == 0){
					out.print("</tr><tr>");
				}
			}
			%>
			</tr>
		</table>
		<div>
			<select form="date_type_select" name = "type" class="form-select form-select-lg mb-3" aria-label=".form-select-lg example">
				<% if("커트".equals(request.getParameter("type"))) { %> <option value="커트" selected> 커트 </option> <%} else %> <option value="커트" > 커트 </option>
				<% if("파마".equals(request.getParameter("type"))) { %> <option value="파마" selected> 파마 </option> <%} else %> <option value="파마" > 파마 </option>
				<% if("염색".equals(request.getParameter("type"))) { %> <option value="염색" selected> 염색 </option> <%} else %> <option value="염색" > 염색 </option>
			</select>
		</div>
		<div>
			<button form="date_type_select" type="submit" class="btn btn-primary">시간 선택하기</button>
		</div>
		</div>
		<div class = "col-6">
			<form action="" method="POST" id="complete_reservation">
				<input name="customerId" type="hidden" value="<%=request.getAttribute("customerId")%>">
				<input name="type" type="hidden" value="<%=request.getParameter("type")%>">
				<input name="date" type="hidden" value="<%=request.getParameter("date")%>">
			</form>
			<div class="list-group list-box">
			<div class="bg-success p-2 text-dark bg-opacity-10 text-center">시간 선택</div>
			<%
	      		ArrayList<Time> availableTimes = (ArrayList<Time>) request.getAttribute("available");
	      	
		      	
				for(Time time : availableTimes) {
					
			%>
				  <label class="form-check-label list-group-item" >
				  <input class="form-check-input" form="complete_reservation" type="radio" value="<%=time %>" name="time">
				    <%=time %>
				  </label>
				
	      		
	        <% 	}   %>
			</div>
			<button form="complete_reservation">예약하기</button>
		</div>
	</div>
	</div>
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-U1DAWAznBHeqEIlVSCgzq+c9gqGAJn5c/t99JyeKa9xxaYpSvHU5awsuZVVFIhvj" crossorigin="anonymous"></script>
</body>

</html>