package com.shHair.reservation.controller;

import java.util.Collections;
import java.util.Comparator;
import java.util.HashMap;
import java.util.List;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpSession;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import com.shHair.reservation.entity.Customer;
import com.shHair.reservation.entity.Reservation;
import com.shHair.reservation.service.SHhairService;

@Controller
@RequestMapping("/admin")
public class AdminController {
	
	@Autowired
	private SHhairService mySHhairService;

//	예약 관련
	@GetMapping("/reservation")
	public String listReservation(HttpServletRequest request, Model model) {
		
		String date = request.getParameter("date");
		List<Reservation> reservations = mySHhairService.getReservationsByDate(date);
		Collections.sort(reservations, new Comparator<Reservation>() {

			@Override
			public int compare(Reservation o1, Reservation o2) {
				return o1.getStartTime().compareTo(o2.getStartTime());
			}
			
		});
		HashMap<Integer, String> customerNames = new HashMap<>();
		for(Reservation reservation : reservations) {
			Customer theCustomer = mySHhairService.getCustomerById(reservation.getCustomerId());
			customerNames.put(reservation.getCustomerId(),theCustomer.getName() );
		}
		model.addAttribute("reservations", reservations);
		model.addAttribute("customerNames", customerNames);

		return "admin/reservation-list";
	}
	
	public static String calculatateTime(String time, int fix) {
		int hour = Integer.parseInt(time.substring(0,2));
		int min = Integer.parseInt(time.substring(3,5));
		min += fix;
		hour += Math.floorDiv(min, 60);
		min = Math.floorMod(min, 60);
		return Integer.toString(hour)+":"+Integer.toString(min);
	}
	
	@GetMapping("/reservation/update/{reservationId}")
	public String updateReservation(@PathVariable int reservationId, HttpServletRequest request, Model model) {
		Reservation theReservation = mySHhairService.getReservationById(reservationId);
		String date = request.getParameter("date");
		String type = request.getParameter("type");
		
		model.addAttribute("theReservation", theReservation);
		model.addAttribute("available", mySHhairService.getAvailabeTime(date,type, theReservation.getCustomerId(), theReservation.getId()));
		return "reservation/reservation-update";
	}
	
	@PostMapping("/reservation/update/{reservationId}")
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
		
		return "redirect:/admin/reservation?date="+date;
	}
	
	@PostMapping("/reservation/delete/{reservationId}")
	public String deleteReservation(@PathVariable int reservationId) {
		
		Reservation theReservation = mySHhairService.getReservationById(reservationId);
		if(theReservation == null) {
			//exception handling
		}
		mySHhairService.deleteReservation(reservationId);
		
		return "redirect:/admin/reservation?date="+theReservation.getDate();
	}
	
	
//	고객 관리 
	@GetMapping("/customer")
	public String listUser(HttpServletRequest request, Model model) {
		
		String name = request.getParameter("name");
		List<Customer> customers = mySHhairService.getCustomersByName(name);
		model.addAttribute("customers", customers);
		
		return "admin/customer-list";
	}
	
	@GetMapping("/customer/{customerId}")
	public String getProfile(@PathVariable int customerId, Model model) {
		
		Customer theCustomer = mySHhairService.getCustomerById(customerId);
		model.addAttribute("theCustomer", theCustomer);
		
		return "admin/customer-profile";
	}
	
	@GetMapping("/customer/update/{customerId}")
	public String getUpdateInfo(@PathVariable int customerId, Model model) {
		
		Customer theCustomer = mySHhairService.getCustomerById(customerId);
		model.addAttribute("theCustomer", theCustomer);
		
		return "admin/customer-update";
	}
	
	@PostMapping("/customer/update/{customerId}")
	public String updateCustomer(@PathVariable int customerId, HttpServletRequest request) {
		
		Customer theCustomer = mySHhairService.getCustomerById(customerId);
		theCustomer.setPhoneNumber(request.getParameter("phoneNumber"));
		theCustomer.setCutTime(Integer.parseInt(request.getParameter("cutTime")));
		theCustomer.setPermTime(Integer.parseInt(request.getParameter("permTime")));
		theCustomer.setDyeTime(Integer.parseInt(request.getParameter("dyeTime")));
		
		mySHhairService.saveCustomer(theCustomer);
		
		return "redirect:/admin/customer/" + theCustomer.getId();
	}
	
	@PostMapping("/customer/delete/{customerId}")
	public String deleteCustomer(@PathVariable int customerId) {
		
		Customer theCustomer = mySHhairService.getCustomerById(customerId);
		if(theCustomer == null) {
			//exception handling
		}
		mySHhairService.deleteCustomer(customerId);
		
		return "redirect:/admin/customer";
	}
	
}
