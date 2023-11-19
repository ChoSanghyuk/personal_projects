package com.example.usageservice.dto;

import lombok.Data;
import java.util.Date;

@Data
public class UsageDto {

    private String userId;

    private Long amount;
}
