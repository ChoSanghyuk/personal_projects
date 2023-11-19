package com.shHair.reservation.entity;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;

import javax.persistence.CascadeType;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.OneToMany;
import javax.persistence.OneToOne;
import javax.persistence.Table;

@Entity
@Table(name="customer")
public class Customer {

	@Id
	@GeneratedValue(strategy=GenerationType.IDENTITY)
	@Column(name="id")
	private int id;
	
	@Column(name="name")
	private String name;
	
	@Column(name="phone_number")
	private String phoneNumber;
	
	@Column(name="sex")
	private int sex;
	// 1 : 남자 , 2: 여자
	
	@Column(name="cut_time")
	private int cutTime;
	
	@Column(name="perm_time")
	private int permTime;
	
	@Column(name="dye_time")
	private int dyeTime;
	
//	@OneToOne(cascade = {CascadeType.ALL})
////	@JoinColumn(name="cu_id", referencedColumnName = "id")
//	private LoginInfo loginInfo;
	
	
	@OneToMany(fetch = FetchType.LAZY , cascade = {CascadeType.PERSIST, CascadeType.DETACH , CascadeType.MERGE , CascadeType.REFRESH})
	@JoinColumn(name="customer_id")
	private List<Reservation> reservations;

	
	
	public Customer() {
		
	}
	
	public Customer(String name, String phoneNumber, int sex) {
		this.name = name;
		this.phoneNumber = phoneNumber;
		this.sex = sex;
		this.cutTime = 0;
		this.permTime = 0;
		this.dyeTime=0;
	}

	
	public int getId() {
		return id;
	}

	public void setId(int id) {
		this.id = id;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public String getPhoneNumber() {
		return phoneNumber;
	}

	public void setPhoneNumber(String phoneNumber) {
		this.phoneNumber = phoneNumber;
	}

	public int getSex() {
		return sex;
	}

	public void setSex(int sex) {
		this.sex = sex;
	}

	public int getCutTime() {
		return cutTime;
	}

	public void setCutTime(int cutTime) {
		this.cutTime = cutTime;
	}

	public int getPermTime() {
		return permTime;
	}

	public void setPermTime(int permTime) {
		this.permTime = permTime;
	}
	
	
	
	
	public int getDyeTime() {
		return dyeTime;
	}

	public void setDyeTime(int dyeTime) {
		this.dyeTime = dyeTime;
	}

//	public LoginInfo getLoginInfo() {
//		return loginInfo;
//	}
//
//	public void setLoginInfo(LoginInfo loginInfo) {
//		this.loginInfo = loginInfo;
//	}

	public List<Reservation> getReservations() {
		Collections.sort(reservations, new Comparator<Reservation>() {
		    @Override
		    public int compare(Reservation a1, Reservation a2) {

		        return -a1.getDate().compareTo(a2.getDate()) ; 
		    }
		});
		return reservations;
	}

	public void setReservations(List<Reservation> reservations) {
		this.reservations = reservations;
	}
	
	public void add(Reservation tempReservation) {
		if( reservations == null) {
			reservations = new ArrayList<>();
		}
		
		reservations.add(tempReservation);
//		tempReservation.setCustomer(this);
	}

	@Override
	public String toString() {
		return "Customer [id=" + id + ", name=" + name + ", phoneNumber=" + phoneNumber + ", sex=" + sex + ", cutTime="
				+ cutTime + ", permTime=" + permTime + ", dyeTime=" + dyeTime
				 + "]";
	}



	
	
	
}
