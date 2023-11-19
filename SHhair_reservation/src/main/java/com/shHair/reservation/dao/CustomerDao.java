package com.shHair.reservation.dao;

import java.util.List;

import com.shHair.reservation.entity.Customer;

public interface CustomerDao {

	public List<Customer> findAll();
	
	public List<Customer> findByName(String name);
	
	public Customer findById(int theId);
		
	public void save(Customer theCustomer);
	
	public void deleteById(int theId);
}
