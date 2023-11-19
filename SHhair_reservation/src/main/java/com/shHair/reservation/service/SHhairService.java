package com.shHair.reservation.service;

import java.util.List;

import com.shHair.reservation.entity.Customer;
import com.shHair.reservation.entity.LoginInfo;
import com.shHair.reservation.entity.Reservation;
import com.shHair.reservation.entity.Time;

public interface SHhairService {
	
	// LogInfo 관련
	public List<LoginInfo> getLoginInfos();
	
	public LoginInfo getLoginInfo(String username);
		
	public void save(LoginInfo theLoginInfo);
	
	public void deleteById(String username);

	// Customer 관련
	public List<Customer> getCustomers();
	
	public List<Customer> getCustomersByName(String name);
	
	public Customer getCustomerById(int theId);
	
	public void saveCustomer(Customer theCustomer);
	
	public void deleteCustomer(int theId);
	
	// Reservation 관련
	public List<Reservation> getReservations();
	
	public Reservation getReservationById(int theId);
	
	public List<Reservation> getReservationsByDate(String date);
	
	public List<Reservation> getReservationsAtferDate(int customerId, String date);
	
	public void saveReservation(Reservation theReservation);
	
	public void deleteReservation(int theId);
	
	// Request 처리 관련
	public List<Time> getAvailabeTime(String date, String type, int customerId, int reservationId);
}
