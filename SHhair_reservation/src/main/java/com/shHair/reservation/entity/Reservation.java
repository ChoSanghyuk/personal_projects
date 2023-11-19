package com.shHair.reservation.entity;

import javax.persistence.CascadeType;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "reservation")
public class Reservation {

	@Id
	@GeneratedValue(strategy = GenerationType.IDENTITY)
	@Column(name="id")
	private int id;
	
//	@ManyToOne(cascade= {CascadeType.DETACH , CascadeType.MERGE, CascadeType.REFRESH , CascadeType.PERSIST})
	@Column(name="customer_id")
	private int customerId;
	
	@Column(name="hair_type")
	private String type;
	
	@Column(name="reservation_date")
	private String date; // 210719
	
	@Column(name="start_time")
	private String startTime; // 123000
	
	@Column(name="end_time")
	private String endTime; // 123000
		
	public Reservation() {
		
	}

	public Reservation(int customerId ,String type, String date, String startTime, String endTime) {
		this.customerId = customerId;
		this.type = type;
		this.date = date;
		this.startTime = startTime;
		this.endTime = endTime;
		
	}

	public int getId() {
		return id;
	}

	public void setId(int id) {
		this.id = id;
	}

	public int getCustomerId() {
		return this.customerId;
	}

	public void setCustomerId(int customerId) {
		this.customerId = customerId;
	}

	public String getType() {
		return type;
	}

	public void setType(String type) {
		this.type = type;
	}

	public String getDate() {
		return date;
	}

	public void setDate(String date) {
		this.date = date;
	}

	public String getStartTime() {
		return startTime;
	}

	public void setStartTime(String startTime) {
		this.startTime = startTime;
	}

	public String getEndTime() {
		return endTime;
	}

	public void setEndTime(String endTime) {
		this.endTime = endTime;
	}

	@Override
	public String toString() {
		return  date + "\t" +  startTime + "\t ~ \t" + endTime +"\t" + type; 
	}

	
}
