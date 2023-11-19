package com.example.settlementservice.entity;

import lombok.Data;

import javax.persistence.*;

@Entity
@Table(name="settlement")
@Data
public class SettlementEntity {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private long id;

    @Column(nullable = false)
    private String userId;

    @Column(nullable = false)
    private Long rate;

    @Column(nullable = false)
    private Long totalCost;
}
