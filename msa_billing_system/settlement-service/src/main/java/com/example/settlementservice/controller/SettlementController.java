package com.example.settlementservice.controller;

import com.example.settlementservice.entity.SettlementEntity;
import com.example.settlementservice.service.SettlementService;
import com.example.settlementservice.vo.RequestId;
import com.example.settlementservice.vo.ResponseSettlement;
import org.apache.http.HttpStatus;
import org.modelmapper.ModelMapper;
import org.modelmapper.convention.MatchingStrategies;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/")
public class SettlementController {

    SettlementService settlementService;

    @Autowired
    public SettlementController(SettlementService settlementService){
        this.settlementService = settlementService;
    }

    @GetMapping("/settlement")
    public ResponseEntity<ResponseSettlement> getSettlement(@RequestBody RequestId requestId){

        SettlementEntity settlementEntity = settlementService.getSettlement(requestId.getUserId());

        ModelMapper mapper = new ModelMapper();
        mapper.getConfiguration().setMatchingStrategy(MatchingStrategies.STRICT);

        ResponseSettlement responseSettlement = mapper.map(settlementEntity, ResponseSettlement.class);

        return new ResponseEntity<>(responseSettlement, null, HttpStatus.SC_OK);

    }
}
