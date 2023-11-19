package com.example.userservice.vo;

import com.example.userservice.entity.ProductEntity;
import com.example.userservice.entity.UserEntity;
import lombok.Data;

@Data
public class ResponseEntr {

    private UserEntity user;

    private ProductEntity product;
}
