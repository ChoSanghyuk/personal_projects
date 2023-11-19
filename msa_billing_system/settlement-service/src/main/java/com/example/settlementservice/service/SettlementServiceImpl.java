package com.example.settlementservice.service;

import com.example.settlementservice.entity.SettlementEntity;
import com.example.settlementservice.repository.SettlementRepository;
import com.example.settlementservice.vo.RequestId;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class SettlementServiceImpl implements SettlementService{

    SettlementRepository settlementRepository;

    @Autowired
    public SettlementServiceImpl(SettlementRepository settlementRepository){
        this.settlementRepository = settlementRepository;
    }
    @Override
    public SettlementEntity getSettlement(String userId) {
        return settlementRepository.findByUserId(userId);
    }
}
