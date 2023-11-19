package com.example.userservice.messageQueue;

import com.example.userservice.dto.SettlementInfoDto;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

@Service
@Slf4j
public class KafkaProducer {
    private KafkaTemplate<String, String> kafkaTemplate;

    @Autowired
    public KafkaProducer(KafkaTemplate<String, String> kafkaTemplate){
        this.kafkaTemplate = kafkaTemplate;
    }

    public SettlementInfoDto send(String topic, SettlementInfoDto settlementInfoDto){

        ObjectMapper mapper = new ObjectMapper();
        String jsonInString = "";

        try{
            jsonInString = mapper.writeValueAsString(settlementInfoDto);
        } catch(JsonProcessingException ex){
            ex.printStackTrace();
        }

        kafkaTemplate.send(topic, jsonInString);
        log.info("kafka Producer sent data from the Usage service" + settlementInfoDto);

        return settlementInfoDto;

    }
}
