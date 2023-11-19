package com.shHair.reservation.service;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collection;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.AuthorityUtils;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

import com.shHair.reservation.entity.LoginInfo;

@Service
public class MyUserDetailsService implements UserDetailsService{

	@Autowired
	private SHhairService shhairService;
	
	@Override
	public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {

		LoginInfo loginInfo = shhairService.getLoginInfo(username) ;
		
		// 비밀먼호를 암호화할 알고리즘을 명시. noop은 암호화하지 x음
		
		
		return new User(loginInfo.getUser() , "{noop}"+loginInfo.getPwd() , getAuthorities(loginInfo) );
	}
	
	private Collection<? extends GrantedAuthority> getAuthorities(LoginInfo loginInfo) {
		System.out.println("hi1");
//		String[] userRoles = new String[] {loginInfo.getRoles()};
		Collection<SimpleGrantedAuthority> authorities = new ArrayList<>();
//				AuthorityUtils.createAuthorityList(userRoles);
		authorities.add(new SimpleGrantedAuthority(loginInfo.getRoles()));
//		System.out.println(authorities);
		return authorities;
	}
	
	

}
