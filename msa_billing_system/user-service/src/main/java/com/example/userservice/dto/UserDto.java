package com.example.userservice.dto;

import com.example.userservice.entity.EntrEntity;
import lombok.Data;

import java.util.ArrayList;
import java.util.List;

@Data
public class UserDto {

    private String userId;
    private List<EntrEntity> entrs = new ArrayList<>();
}
