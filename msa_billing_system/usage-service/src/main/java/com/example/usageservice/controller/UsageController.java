package com.example.usageservice.controller;

import com.example.usageservice.dto.UsageDto;
import com.example.usageservice.entity.UsageEntity;
import com.example.usageservice.messageQueue.KafkaProducer;
import com.example.usageservice.service.UsageService;
import com.example.usageservice.vo.RequestUsage;
import com.example.usageservice.vo.ResponseUsage;
import org.apache.http.HttpStatus;
import org.jboss.jandex.TypeTarget;
import org.modelmapper.ModelMapper;
import org.modelmapper.convention.MatchingStrategies;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/")
public class UsageController {

    UsageService usageService;
    KafkaProducer kafkaProducer;

    @Autowired
    public UsageController(UsageService usageService, KafkaProducer kafkaProducer){

        this.usageService = usageService;
        this.kafkaProducer = kafkaProducer;
    }


    @PostMapping("/use")
    public ResponseEntity<ResponseUsage> use(@RequestBody RequestUsage usage){

        ModelMapper mapper = new ModelMapper();
        mapper.getConfiguration().setMatchingStrategy(MatchingStrategies.STRICT);

        UsageEntity usageEntity = usageService.createUser(usage);

        UsageDto usageDto = mapper.map(usageEntity, UsageDto.class);
        ResponseUsage responseUser = mapper.map(usageEntity, ResponseUsage.class);

        /* send message */
        kafkaProducer.send("update-settlement" , usageDto);

        return new ResponseEntity<>(responseUser, null, HttpStatus.SC_CREATED);

    }
}
