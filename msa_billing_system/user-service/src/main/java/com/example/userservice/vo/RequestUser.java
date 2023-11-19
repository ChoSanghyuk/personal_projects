package com.example.userservice.vo;

//import jakarta.validation.constraints.NotNull;
import lombok.Data;

@Data
public class RequestUser {

    //@NotNull(message = "userId should not be null")
    private String userId;
}
