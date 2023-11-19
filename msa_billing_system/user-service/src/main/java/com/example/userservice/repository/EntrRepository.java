package com.example.userservice.repository;

import com.example.userservice.entity.EntrEntity;
import org.springframework.data.repository.CrudRepository;

public interface EntrRepository  extends CrudRepository<EntrEntity, Long> {
}
