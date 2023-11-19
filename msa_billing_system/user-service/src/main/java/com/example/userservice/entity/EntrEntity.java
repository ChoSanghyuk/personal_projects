package com.example.userservice.entity;

import com.fasterxml.jackson.annotation.JsonBackReference;
import lombok.Data;

import javax.persistence.*;

@Data
@Entity
@Table(name = "entr")
public class EntrEntity {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @JsonBackReference
    @ManyToOne
    @JoinColumn(name="users_id")
    private UserEntity user;

    @JsonBackReference
    @ManyToOne
    @JoinColumn(name="products_id")
    private ProductEntity product;

}
