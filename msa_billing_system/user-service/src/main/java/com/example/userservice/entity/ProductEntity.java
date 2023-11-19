package com.example.userservice.entity;

import com.fasterxml.jackson.annotation.JsonManagedReference;
import lombok.Data;

import javax.persistence.*;
import java.util.ArrayList;
import java.util.List;

@Data
@Entity
@Table(name = "products")
public class ProductEntity {

//    @Id
//    @GeneratedValue(strategy = GenerationType.IDENTITY)
//    private Long id;

    @Id
    @Column(nullable = false)
    private String name;

    @Column(nullable = false)
    private Long rate;

    @JsonManagedReference
    @OneToMany(mappedBy = "product")
    private List<EntrEntity> entrs = new ArrayList<>();


}
