package com.example.usageservice.repository;

import com.example.usageservice.entity.UsageEntity;
import org.springframework.data.repository.CrudRepository;

public interface UsageRepository extends CrudRepository<UsageEntity, Long> {


}
