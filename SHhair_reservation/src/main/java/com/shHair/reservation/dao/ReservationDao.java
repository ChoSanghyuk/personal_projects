package com.shHair.reservation.dao;

import java.util.List;

import com.shHair.reservation.entity.Reservation;

public interface ReservationDao {

	public List<Reservation> findAll();
	
	public Reservation findById(int theId);
	
	public List<Reservation> getReservationsAfterDate(int customerId, String date);
	
	public List<Reservation> findByDate(String date);
	
	public void save(Reservation theReservation);
	
	public void deleteById(int theId);
}
