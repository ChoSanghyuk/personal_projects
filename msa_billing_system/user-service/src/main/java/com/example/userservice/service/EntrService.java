package com.example.userservice.service;

import com.example.userservice.dto.EntrDto;
import com.example.userservice.entity.ProductEntity;
import com.example.userservice.entity.UserEntity;

public interface EntrService {

    public void createEntr(EntrDto entrDto);

    public UserEntity getUserById(String userId);

    public ProductEntity getProductByName(String productName);
}
