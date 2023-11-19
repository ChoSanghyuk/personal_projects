package com.shHair.reservation.dao;

import java.util.List;

import javax.persistence.EntityManager;

import org.hibernate.Session;
import org.hibernate.query.Query;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import com.shHair.reservation.entity.Reservation;

@Repository
public class ReservationDaoImpl implements ReservationDao {
	
	private EntityManager entityManager;
	
	@Autowired
	public ReservationDaoImpl(EntityManager theEntityManager) {
		this.entityManager = theEntityManager;
	}

	@Override
	public List<Reservation> findAll() {
		Session currentSession = entityManager.unwrap(Session.class);
		
		Query<Reservation> theQuery = currentSession.createQuery("from Reservation" ,  Reservation.class);
		
		List<Reservation> reservations = theQuery.getResultList();
		
		return reservations;
	}

	@Override
	public Reservation findById(int theId) {
		Session currentSession = entityManager.unwrap(Session.class);
		Reservation theReservation = currentSession.get(Reservation.class, theId);
		
		return theReservation;
	}
	
	@Override
	public List<Reservation> findByDate(String date) {
		Session currentSession = entityManager.unwrap(Session.class);
		
		Query<Reservation> theQuery = currentSession.createQuery("from Reservation where date=:reservationDate" ,  Reservation.class);
		theQuery.setParameter("reservationDate", date);
		List<Reservation> reservations = theQuery.getResultList();
		return reservations;
	}

	@Override
	public List<Reservation> getReservationsAfterDate(int customerId, String date) {
		Session currentSession = entityManager.unwrap(Session.class);
		
		Query<Reservation> theQuery = currentSession.createQuery("from Reservation where date>=:reservationDate and customer_id=:id" 
										,  Reservation.class);
		theQuery.setParameter("reservationDate", date);
		theQuery.setParameter("id", customerId);
		List<Reservation> reservations = theQuery.getResultList();
		return reservations;
	}

	@Override
	public void save(Reservation theReservation) {
		Session currentSession = entityManager.unwrap(Session.class);
		
		currentSession.saveOrUpdate(theReservation);
	}

	@Override
	public void deleteById(int theId) {
		Session currentSession = entityManager.unwrap(Session.class);
		
		Query theQuery = currentSession.createQuery("delete from Reservation where id=:reservationId");
		theQuery.setParameter("reservationId", theId);
		
		theQuery.executeUpdate();
		
	}



	


}
