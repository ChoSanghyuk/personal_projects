package com.example.userservice.controller;

import com.example.userservice.dto.EntrDto;
import com.example.userservice.dto.SettlementInfoDto;
import com.example.userservice.dto.UserDto;
import com.example.userservice.entity.ProductEntity;
import com.example.userservice.entity.UserEntity;
import com.example.userservice.messageQueue.KafkaProducer;
import com.example.userservice.service.EntrService;
import com.example.userservice.service.UserService;
import com.example.userservice.service.UserServiceImpl;
import com.example.userservice.vo.RequestEntr;
import com.example.userservice.vo.RequestUser;
import com.example.userservice.vo.ResponseEntr;
import com.example.userservice.vo.ResponseUser;
import lombok.extern.slf4j.Slf4j;
import org.apache.http.HttpStatus;
import org.modelmapper.ModelMapper;
import org.modelmapper.convention.MatchingStrategies;

import org.modelmapper.spi.MatchingStrategy;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/")
@Slf4j
public class UserController {

    UserService userService;
    EntrService entrService;
    KafkaProducer kafkaProducer;

    @Autowired
    public UserController(UserService userService, EntrService entrService, KafkaProducer kafkaProducer){
        this.userService = userService;
        this.entrService = entrService;
        this.kafkaProducer = kafkaProducer;
    }

    @PostMapping("/signup")
    public ResponseEntity<ResponseUser> createUser(@RequestBody RequestUser user){
        ModelMapper mapper = new ModelMapper();
        mapper.getConfiguration().setMatchingStrategy(MatchingStrategies.STRICT);

        UserDto userDto = mapper.map(user, UserDto.class);
        userService.createUser(userDto);

        ResponseUser responseUser = mapper.map(userDto, ResponseUser.class);

        //return ReponseEntity.status(HttpStatus.SC_CREATED).body(responseUser);
        return new ResponseEntity<>(responseUser, null, HttpStatus.SC_CREATED);
    }

    @PostMapping("/entr")
    public ResponseEntity<ResponseEntr> createEntr(@RequestBody RequestEntr entr){

        ModelMapper mapper = new ModelMapper();
        mapper.getConfiguration().setMatchingStrategy(MatchingStrategies.STRICT);

//        EntrDto entrDto = mapper.map(entr, EntrDto.class);
        EntrDto entrDto = new EntrDto();
        entrDto.setUser(entrService.getUserById(entr.getUserId()));
        entrDto.setProduct(entrService.getProductByName(entr.getProductName()));
        entrService.createEntr(entrDto);

        SettlementInfoDto settlementInfoDto = new SettlementInfoDto();
        settlementInfoDto.setUserId(entrDto.getUser().getUserId());
        settlementInfoDto.setRate(entrDto.getProduct().getRate());


        kafkaProducer.send("create-settlement" , settlementInfoDto);
        log.info(settlementInfoDto.toString());

        ResponseEntr responseEntr = mapper.map(entrDto, ResponseEntr.class);
        return new ResponseEntity<>(responseEntr, null, HttpStatus.SC_CREATED);

    }

}
