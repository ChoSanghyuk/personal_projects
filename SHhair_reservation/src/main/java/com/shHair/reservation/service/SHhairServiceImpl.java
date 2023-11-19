package com.shHair.reservation.service;

import java.util.ArrayList;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import com.shHair.reservation.dao.CustomerDao;
import com.shHair.reservation.dao.LoginInfoDao;
import com.shHair.reservation.dao.ReservationDao;
import com.shHair.reservation.entity.Customer;
import com.shHair.reservation.entity.LoginInfo;
import com.shHair.reservation.entity.Reservation;
import com.shHair.reservation.entity.Time;

import java.util.Collections;
import java.util.Comparator;

@Service
public class SHhairServiceImpl implements SHhairService {
	
	private LoginInfoDao loginInfoDao;
	private CustomerDao customerDao;
	private ReservationDao reservatonDao;
	
	@Autowired
	public SHhairServiceImpl(CustomerDao theCustomerDao , ReservationDao theReservationDao, LoginInfoDao theLoginInfoDao ) {
		this.loginInfoDao = theLoginInfoDao;
		this.customerDao = theCustomerDao;
		this.reservatonDao =  theReservationDao;
	}

	@Override
	@Transactional
	public List<LoginInfo> getLoginInfos() {
		return loginInfoDao.findAll();
	}

	@Override
	@Transactional
	public LoginInfo getLoginInfo(String username) {
		return loginInfoDao.findById(username);
	}

	@Override
	@Transactional
	public void save(LoginInfo theLoginInfo) {
		loginInfoDao.save(theLoginInfo);
	}

	@Override
	@Transactional
	public void deleteById(String username) {
		loginInfoDao.deleteById(username);
	}
	
	@Override
	@Transactional
	public List<Customer> getCustomers() {
		return customerDao.findAll();
	}
	
	@Override
	@Transactional
	public List<Customer> getCustomersByName(String name){
		return customerDao.findByName(name);
	}

	@Override
	@Transactional
	public Customer getCustomerById(int theId) {
		return customerDao.findById(theId);
	}

	@Override
	@Transactional
	public void saveCustomer(Customer theCustomer) {
		customerDao.save(theCustomer);
	}

	@Override
	@Transactional
	public void deleteCustomer(int theId) {
		customerDao.deleteById(theId);
	}

	@Override
	@Transactional
	public List<Reservation> getReservations() {
		return reservatonDao.findAll();
	}

	@Override
	@Transactional
	public Reservation getReservationById(int theId) {
		return reservatonDao.findById(theId);
	}
	
	@Override
	public List<Reservation> getReservationsByDate(String date) {
		return reservatonDao.findByDate(date);
	}

	@Override
	@Transactional
	public List<Reservation> getReservationsAtferDate(int customerId, String date) {
		return reservatonDao.getReservationsAfterDate(customerId, date);
	}

	@Override
	@Transactional
	public void saveReservation(Reservation theReservation) {
		reservatonDao.save(theReservation);
	}

	@Override
	@Transactional
	public void deleteReservation(int theId) {
		reservatonDao.deleteById(theId);
	}

	@Override
	public List<Time> getAvailabeTime(String date, String type, int customerId, int reservationId) {
		
		if(date == null) {
			return new ArrayList<Time>();
		}
		
		List<Reservation> reservations = getReservationsByDate(date);
		Customer theCustomer = getCustomerById(customerId);
		Collections.sort(reservations, new Comparator<Reservation>() {

			@Override
			public int compare(Reservation o1, Reservation o2) {
				return o1.getStartTime().compareTo(o2.getStartTime());
			}
			
		});
		List<Time> availabletimes = new ArrayList<>();
		int reservationsIdx = 0;

		String startTime;
		String endTime;
		int howLong;
		boolean isAvailable;
		for(int i = 0 ; i < 18 ; i++) {
			startTime = "" + (10 + i/2) + ":" + ( i%2 == 0 ? "00":"30") ;
			
			if(type.equals("커트")) {
				howLong = 30 + theCustomer.getCutTime();
			} else if(type.equals("파마")) {
				howLong = 120 + theCustomer.getPermTime();
			} else {
				howLong = 60 + theCustomer.getDyeTime();
			}
			howLong += i*30;
			
			endTime = "" + (10 + howLong/60) + ":" + (howLong%60==0? "00":"30") ;
			isAvailable = true;
			for(Reservation theReservation : reservations) {
				if(theReservation.getId() == reservationId ) {
					continue;
				}
				if( (endTime.compareTo(theReservation.getStartTime().substring(0,5))>0 && startTime.compareTo(theReservation.getEndTime().substring(0,5)) <0  )) {
					isAvailable = false;
					break;
				} else if(endTime.compareTo(theReservation.getStartTime()) <=0) {
					break;
				}
			}
			if(isAvailable) {
				availabletimes.add(new Time(startTime, endTime));
			}
	
		}

		return availabletimes;
	}


}
