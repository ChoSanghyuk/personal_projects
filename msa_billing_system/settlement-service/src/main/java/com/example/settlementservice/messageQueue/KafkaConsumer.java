package com.example.settlementservice.messageQueue;

import com.example.settlementservice.entity.SettlementEntity;
import com.example.settlementservice.repository.SettlementRepository;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.Map;

@Service
@Slf4j
public class KafkaConsumer {

    SettlementRepository settlementRepository;

    @Autowired
    public KafkaConsumer(SettlementRepository settlementRepository){
        this.settlementRepository = settlementRepository;
    }

    @KafkaListener(topics = "create-settlement")
    public void createSettlement(String kafkaMessage){
        Map<Object,Object> map = new HashMap<>();
        ObjectMapper mapper = new ObjectMapper();

        try {
            map = mapper.readValue(kafkaMessage, new TypeReference<Map<Object, Object>>() {
            });
        } catch (JsonProcessingException ex){
            ex.printStackTrace();
        }

        SettlementEntity settlementEntity = new SettlementEntity();
        settlementEntity.setUserId( (String) map.get("userId") );
        settlementEntity.setRate( Long.parseLong(map.get("rate").toString()));
        settlementEntity.setTotalCost(Long.valueOf(0));

        log.info(settlementEntity.toString());
        settlementRepository.save(settlementEntity);


    }

    @KafkaListener(topics = "update-settlement")
    public void updateSettlement(String kafkaMessage){
        log.info("kafka Message: ->" + kafkaMessage);

        Map<Object, Object> map = new HashMap<>();
        ObjectMapper mapper = new ObjectMapper();

        try{
            map = mapper.readValue(kafkaMessage, new TypeReference<Map<Object, Object>>(){
            });
        } catch (JsonProcessingException ex){
            ex.printStackTrace();
        }

        SettlementEntity settlementEntity = settlementRepository.findByUserId((String)map.get("userId"));
        if(settlementEntity != null){
            settlementEntity.setTotalCost( settlementEntity.getTotalCost() + (settlementEntity.getRate() * (Integer) map.get("amount") ) ); // 사용 금액 로직 수정
            settlementRepository.save(settlementEntity);
        }
    }
}
