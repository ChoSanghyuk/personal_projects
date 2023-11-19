package com.shHair.reservation.entity;

import javax.persistence.CascadeType;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.OneToOne;
import javax.persistence.Table;

@Entity
@Table(name="logininfo")
public class LoginInfo {

	
	@Id
	@Column(name="user")
	private String user;
	
	@Column(name="pwd")
	private String pwd;
	
	@Column(name="roles")
	private String roles;
	
//	@OneToOne(cascade= CascadeType.ALL)
	@Column(name="cu_id")
	private int customerId;
//	
//	@OneToOne
//	@JoinColumn(name="cu_id", referencedColumnName="id")
//	private Customer customer;
	
	public LoginInfo() {
		
	}
	
	
	public LoginInfo(String user, String pwd, String roles, int customerId) {
		this.user = user;
		this.pwd = pwd;
		this.roles = roles;
		this.customerId = customerId;
	}

	public String getUser() {
		return user;
	}

	public void setUser(String user) {
		this.user = user;
	}

	public String getPwd() {
		return pwd;
	}

	public void setPwd(String pwd) {
		this.pwd = pwd;
	}



	public String getRoles() {
		return roles;
	}



	public void setRoles(String roles) {
		this.roles = roles;
	}



	public int getCustomerId() {
		return customerId;
	}



	public void setCustomerId(int customerId) {
		this.customerId = customerId;
	}



	@Override
	public String toString() {
		return "LoginInfo [user=" + user + ", pwd=" + pwd + ", roles=" + roles + ", customerId=" + customerId + "]";
	}


	
	

	
}
