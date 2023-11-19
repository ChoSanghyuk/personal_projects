package com.example.userservice.entity;

import com.fasterxml.jackson.annotation.JsonManagedReference;
import lombok.Data;
import javax.persistence.*;
import java.util.ArrayList;
import java.util.List;

@Data
@Entity
@Table(name = "users")
public class UserEntity {

//    @Id
//    @GeneratedValue(strategy = GenerationType.IDENTITY)
//    private Long id;

    @Id
    @Column(nullable = false)
    private String userId;

    @JsonManagedReference
    @OneToMany(mappedBy = "user")
    private List<EntrEntity> entrs = new ArrayList<>();

}
