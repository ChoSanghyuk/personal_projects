package com.shHair.reservation.dao;

import java.util.List;

import javax.persistence.EntityManager;

import org.hibernate.Session;
import org.hibernate.query.Query;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import com.shHair.reservation.entity.Customer;
import com.shHair.reservation.entity.LoginInfo;

@Repository
public class LoginInfoDaoImpl implements LoginInfoDao {

	private EntityManager entityManager;
	
	@Autowired
	public LoginInfoDaoImpl(EntityManager theEntityManager) {
		this.entityManager = theEntityManager;
	}
	
	@Override
	public List<LoginInfo> findAll() {
		Session currentSession = entityManager.unwrap(Session.class);
		Query<LoginInfo> theQuery = currentSession.createQuery("from LoginInfo" ,  LoginInfo.class);
		List<LoginInfo> loginInfos = theQuery.getResultList();
		return loginInfos;
	}

	@Override
	public LoginInfo findById(String username) {
		Session currentSession = entityManager.unwrap(Session.class);
		LoginInfo theLoginInfo = currentSession.get(LoginInfo.class, username);
		
		return theLoginInfo;
	}

	@Override
	public void save(LoginInfo theLoginInfo) {
		Session currentSession = entityManager.unwrap(Session.class);
		
		currentSession.saveOrUpdate(theLoginInfo);
	}

	@Override
	public void deleteById(String username) {
		Session currentSession = entityManager.unwrap(Session.class);
		
		Query theQuery = currentSession.createQuery("delete from LoginInfo where id=:user");
		theQuery.setParameter("user", username);
		
		theQuery.executeUpdate();

	}

}
