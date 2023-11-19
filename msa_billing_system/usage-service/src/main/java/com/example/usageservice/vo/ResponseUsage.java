package com.example.usageservice.vo;

import lombok.Data;
import java.util.Date;

@Data
public class ResponseUsage {

    private String userId;

    private Long amount;

    private Date createdAt;

}
