package com.example.settlementservice.service;

import com.example.settlementservice.entity.SettlementEntity;
import com.example.settlementservice.vo.RequestId;

public interface SettlementService {

    SettlementEntity getSettlement(String userId);
}
