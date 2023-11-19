package com.shHair.reservation.dao;

import java.util.List;

import com.shHair.reservation.entity.Customer;
import com.shHair.reservation.entity.LoginInfo;

public interface LoginInfoDao {
	
	public List<LoginInfo> findAll();
	
	public LoginInfo findById(String username);
		
	public void save(LoginInfo theLoginInfo);
	
	public void deleteById(String username);

}
