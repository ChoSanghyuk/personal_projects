package com.example.userservice.vo;

import com.example.userservice.entity.ProductEntity;
import com.example.userservice.entity.UserEntity;
//import javax.validation.constraints.NotNull;
import lombok.Data;

@Data
public class RequestEntr {

    //@NotNull(message = "userId should not be null")
    private String userId;

   //@NotNull(message = "product should not be null")
    private String productName;
}
