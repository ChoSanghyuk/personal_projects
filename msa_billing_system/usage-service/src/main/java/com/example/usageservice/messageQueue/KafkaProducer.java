package com.example.usageservice.messageQueue;

import com.example.usageservice.dto.UsageDto;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.netflix.discovery.converters.Auto;
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

    public UsageDto send(String topic, UsageDto usageDto){

        ObjectMapper mapper = new ObjectMapper();
        String jsonInString = "";

        try{
            jsonInString = mapper.writeValueAsString(usageDto);
        } catch(JsonProcessingException ex){
            ex.printStackTrace();
        }

        kafkaTemplate.send(topic, jsonInString);
        log.info("kafka Producer sent data from the Usage service" + usageDto);

        return usageDto;

    }
}
