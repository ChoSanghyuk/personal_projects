package com.example.usageservice.vo;


//import jakarta.validation.constraints.NotNull;
import lombok.Data;

@Data
public class RequestUsage {

//    @NotNull(message = "userId should not be null")
    private String userId;

//    @NotNull(message = "amount should not be null")
    private Long amount;

}
