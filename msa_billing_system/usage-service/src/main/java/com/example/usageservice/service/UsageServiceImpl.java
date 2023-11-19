package com.example.usageservice.service;

import com.example.usageservice.entity.UsageEntity;
import com.example.usageservice.repository.UsageRepository;
import com.example.usageservice.vo.RequestUsage;
import org.modelmapper.ModelMapper;
import org.modelmapper.convention.MatchingStrategies;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class UsageServiceImpl implements UsageService {

    UsageRepository usageRepository;

    @Autowired
    public UsageServiceImpl(UsageRepository usageRepository){
        this.usageRepository = usageRepository;
    }

    @Override
    public UsageEntity createUser(RequestUsage usage) {

        ModelMapper mapper = new ModelMapper();
        mapper.getConfiguration().setMatchingStrategy(MatchingStrategies.STRICT);

        UsageEntity usageEntity = mapper.map(usage, UsageEntity.class);

        usageRepository.save(usageEntity);

        return usageEntity;

    }
}
