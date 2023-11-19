package com.example.userservice.repository;

import com.example.userservice.entity.ProductEntity;
import com.example.userservice.entity.UserEntity;
import org.springframework.data.repository.CrudRepository;

public interface ProductRepository extends CrudRepository<ProductEntity, Long> {

    public ProductEntity findProductByName(String productName);
}
