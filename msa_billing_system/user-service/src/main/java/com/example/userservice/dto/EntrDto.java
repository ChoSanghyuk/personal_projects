package com.example.userservice.dto;

import com.example.userservice.entity.ProductEntity;
import com.example.userservice.entity.UserEntity;
import lombok.Data;


@Data
public class EntrDto {

    private UserEntity user;
    private ProductEntity product;

}
