package com.example.userservice.service;

import com.example.userservice.dto.EntrDto;
import com.example.userservice.entity.EntrEntity;
import com.example.userservice.entity.ProductEntity;
import com.example.userservice.entity.UserEntity;
import com.example.userservice.repository.EntrRepository;
import com.example.userservice.repository.ProductRepository;
import com.example.userservice.repository.UserRepository;
import org.modelmapper.ModelMapper;
import org.modelmapper.convention.MatchingStrategies;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class EntrServiceImpl implements EntrService{

    EntrRepository entrRepository;
    UserRepository userRepository;
    ProductRepository productRepository;

    @Autowired
    public EntrServiceImpl(EntrRepository entrRepository, UserRepository userRepository, ProductRepository productRepository){
        this.entrRepository = entrRepository;
        this.userRepository = userRepository;
        this.productRepository = productRepository;
    }

    @Override
    public void createEntr(EntrDto entrDto) {

        ModelMapper mapper = new ModelMapper();
        mapper.getConfiguration().setMatchingStrategy(MatchingStrategies.STRICT);

        EntrEntity entrEntity = mapper.map(entrDto, EntrEntity.class);
        entrRepository.save(entrEntity);

    }

    @Override
    public UserEntity getUserById(String userId) {
        return userRepository.findUserByUserId(userId);
    }

    @Override
    public ProductEntity getProductByName(String productName) {
        return productRepository.findProductByName(productName);
    }
}
