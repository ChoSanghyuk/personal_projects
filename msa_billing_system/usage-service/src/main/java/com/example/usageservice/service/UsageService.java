package com.example.usageservice.service;

import com.example.usageservice.entity.UsageEntity;
import com.example.usageservice.vo.RequestUsage;

public interface UsageService {

    UsageEntity createUser(RequestUsage usage);

}
