package com.shHair.reservation.controller;

import java.util.List;

import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.shHair.reservation.entity.Customer;
import com.shHair.reservation.entity.LoginInfo;
import com.shHair.reservation.entity.Reservation;
import com.shHair.reservation.service.SHhairService;

@RestController
@RequestMapping("/api")
public class SHhairRestController {

	private SHhairService shhairService;
	
	@Autowired
	public SHhairRestController(SHhairService theSHhairService) {
		this.shhairService = theSHhairService;
	}
	
	// 아직 LoginInfo Rest Mapping은 만들지 않음
	
	@GetMapping("/loginInfo/{username}")
	public LoginInfo findLoginInfo(@PathVariable String username) {
		
		LoginInfo theLoginInfo = shhairService.getLoginInfo(username) ;
		
		if(theLoginInfo == null) {
			//exception handling
		}
		return theLoginInfo ;
	}
	
	
	@GetMapping("/customers")
	public List<Customer> findAllCustomers(){
		return shhairService.getCustomers();
	}
	
	@GetMapping("/reservations")
	public List<Reservation> findAllReservations(){
		return shhairService.getReservations();
	}
	
	@GetMapping("/customer/{customerId}")
	public Customer findCustomer(@PathVariable int customerId) {
		
		Customer theCustomer = shhairService.getCustomerById(customerId) ;
		
		if(theCustomer == null) {
			//exception handling
		}
		return theCustomer ;
	}
	
	@GetMapping("/reservation/{reservationId}")
	public Reservation findReservation(@PathVariable int ReservationId) {
		
		Reservation theReservation = shhairService.getReservationById(ReservationId) ;
		
		if(theReservation == null) {
			//exception handling
		}
		return theReservation ;
	}
	
	@PostMapping("/customer")
	public Customer addCustomer(@RequestBody Customer theCustomer) {
		theCustomer.setId(0);
		shhairService.saveCustomer(theCustomer);
		
		return theCustomer;
	}
	
	@PostMapping("/reservation")
	public Reservation addReservation(@RequestBody Reservation theReservation) {
		theReservation.setId(0);
		shhairService.saveReservation(theReservation);
		
		return theReservation;
	}
	
	@PutMapping("/customer")
	public Customer updateCustomer(@RequestBody Customer theCustomer) {
		shhairService.saveCustomer(theCustomer);
		return theCustomer;
	}
	
	@PutMapping("/reservation")
	public Reservation updateReservation(@RequestBody Reservation theReservation) {
		shhairService.saveReservation(theReservation);
		return theReservation;
	}
	
	@DeleteMapping("/customer/{customerId}")
	public String deleteCustomer(@PathVariable int customerId) {
		Customer theCustomer = shhairService.getCustomerById(customerId);
		
		if(theCustomer == null) {
			//exception handling
		}
		
		shhairService.deleteCustomer(customerId);
		return "Deleted Customer : " + theCustomer.getName();
	}
	
	@DeleteMapping("/reservation/{reservationId}")
	public String deleteReservation(@PathVariable int reservationId) {
		Reservation theReservation = shhairService.getReservationById(reservationId);
		
		if(theReservation == null) {
			//exception handling
		}
		
		shhairService.deleteReservation(reservationId);
		return "Deleted Reservation : " + theReservation.getDate() + " " + theReservation.getStartTime() + " ~ " + theReservation.getEndTime() ;
	}
	
//	@GetMapping("/getreservation")
//	public String getReservation(HttpServletRequest request, Model model) {
//		
//		String date = request.getParameter("reservation-date");
//		model.addAttribute("date", date);
//		System.out.println(date);
//		
//		return date;
//	}
	
}
