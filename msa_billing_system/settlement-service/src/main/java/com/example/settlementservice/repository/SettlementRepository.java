package com.example.settlementservice.repository;

import com.example.settlementservice.entity.SettlementEntity;
import org.springframework.data.repository.CrudRepository;

public interface SettlementRepository extends CrudRepository<SettlementEntity, Long> {

    SettlementEntity findByUserId(String userId);
}
