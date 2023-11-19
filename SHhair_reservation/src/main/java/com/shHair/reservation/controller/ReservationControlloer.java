package com.shHair.reservation.controller;


import java.util.ArrayList;
import java.util.HashMap;

import java.util.List;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpSession;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.User;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;

import com.shHair.reservation.entity.Customer;
import com.shHair.reservation.entity.LoginInfo;
import com.shHair.reservation.entity.Reservation;
import com.shHair.reservation.service.SHhairService;


@Controller
@RequestMapping("/reservation")
public class ReservationControlloer {
	
	@Autowired
	private SHhairService mySHhairService;

	@RequestMapping("/asdf")
	public String showCalendar() {
		return "reservation/reservation-select-time";
	}
	
	@GetMapping("")
	public String getReservation(HttpServletRequest request, Model model) {
		String date = request.getParameter("date");
		String type = request.getParameter("type");
		
		int theCustomerId = getLoggedInCustomerId();
		
		model.addAttribute("customerId", theCustomerId);
		model.addAttribute("available", mySHhairService.getAvailabeTime(date, type,theCustomerId,0));
		return "reservation/calendar";
	}
	
	@PostMapping("")
	public String reserve(HttpServletRequest request, Model model) {

		int customerId = Integer.parseInt(request.getParameter("customerId"));
		String theType = request.getParameter("type");
		String theDate = request.getParameter("date");
		String theStartTime = request.getParameter("time").substring(0,5)+":00";
		String theEndTime = request.getParameter("time").substring(8,13)+":00";
		
		Reservation theReservation = new Reservation(customerId, theType, theDate, theStartTime, theEndTime);

		theReservation.setId(0);
		mySHhairService.saveReservation(theReservation);

		return "redirect:/reservation/reservation-done/" + theReservation.getId();
	}
	
	@GetMapping("/reservation-done/{id}")
	public String checkReservation(@PathVariable int id, Model model) {
		
		Reservation theReservation = mySHhairService.getReservationById(id);
		Customer theCustomer = mySHhairService.getCustomerById(theReservation.getCustomerId());
		
		model.addAttribute("customerName",theCustomer.getName());
		model.addAttribute("theReservation",theReservation);
		return "reservation/reservation-done";
	}
	
	@GetMapping("/confirm")
	public String confirmReservation(HttpServletRequest request, Model model) {
		int theCustomerId = getLoggedInCustomerId();
		Customer theCustomer = mySHhairService.getCustomerById(theCustomerId);
		model.addAttribute("customerName", theCustomer.getName());
		model.addAttribute("theReservation", theCustomer.getReservations().get(0));
		
		model.addAttribute("theReservatoins", mySHhairService.getReservationsAtferDate(theCustomer.getId(), "2021-09-28" ));
				
		return "reservation/reservation-check";
	}
	

	@GetMapping("/update/{reservationId}")
	public String getReservation(@PathVariable int reservationId, HttpServletRequest request, Model model) {
		// 권한 설정
		Reservation theReservation = mySHhairService.getReservationById(reservationId);
		model.addAttribute("theReservation", theReservation);
		
		model.addAttribute("customerId", theReservation.getCustomerId());
		model.addAttribute("available", mySHhairService.getAvailabeTime(request.getParameter("date"),
				request.getParameter("type"),theReservation.getCustomerId(),reservationId));
		
		return "reservation/reservation-update";
	}
	
	@PostMapping("/update/{reservationId}")
	public String updateReservation(@PathVariable int reservationId, HttpServletRequest request) {
		
		Reservation theReservation = mySHhairService.getReservationById(reservationId);
		String date = request.getParameter("date");
		theReservation.setDate(date);
		theReservation.setType(request.getParameter("type"));
		theReservation.setStartTime(request.getParameter("time").substring(0,5)+":00");
		theReservation.setEndTime(request.getParameter("time").substring(8,13)+":00");

		if(theReservation == null) {
			// exception handling
		}
		
		mySHhairService.saveReservation(theReservation);
		
		return "redirect:/reservation/reservation-done/" + theReservation.getId();
	}
	
	@PostMapping("/delete/{reservationId}")
	public String deleteReservation(@PathVariable int reservationId) {
		// 권한 설정
		
		Reservation theReservation = mySHhairService.getReservationById(reservationId);
		if(theReservation == null) {
			//exception handling
		}
		mySHhairService.deleteReservation(reservationId);
		return "redirect:/admin";
	}
	
	public int getLoggedInCustomerId() {
		Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
		User theUser = (User) authentication.getPrincipal();
		LoginInfo theLoginInfo = mySHhairService.getLoginInfo(theUser.getUsername());
		return theLoginInfo.getCustomerId();
	}
	
}
